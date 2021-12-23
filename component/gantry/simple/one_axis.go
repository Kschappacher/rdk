package simple

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/edaniels/golog"

	"go.viam.com/utils"

	"go.viam.com/rdk/component/board"
	"go.viam.com/rdk/component/gantry"
	"go.viam.com/rdk/component/motor"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/robot"
)

func init() {
	registry.RegisterComponent(gantry.Subtype, "simpleoneaxis", registry.Component{
		Constructor: func(ctx context.Context, r robot.Robot, config config.Component, logger golog.Logger) (interface{}, error) {
			return NewOneAxis(ctx, r, config, logger)
		}})
}

// NewOneAxis creates a new one axis gantry
func NewOneAxis(ctx context.Context, r robot.Robot, config config.Component, logger golog.Logger) (gantry.Gantry, error) {
	g := &oneAxis{
		name:   config.Name,
		logger: logger,
	}

	var ok bool

	g.motor, ok = r.MotorByName(config.Attributes.String("motor"))
	if !ok {
		return nil, errors.New("cannot find motor for gantry")
	}

	g.limitSwitchPins = config.Attributes.StringSlice("limitPins")
	if len(g.limitSwitchPins) != 2 {
		return nil, errors.New("need 2 limitPins")
	}
	g.limitBoard, ok = r.BoardByName(config.Attributes.String("limitBoard"))
	if !ok {
		return nil, errors.New("cannot find board for gantry")
	}
	g.limitHigh = config.Attributes.Bool("limitHigh", true)

	g.lengthMeters = config.Attributes.Float64("lengthMeters", 0.0)
	if g.lengthMeters <= 0 {
		return nil, errors.New("gantry length has to be >= 0")
	}

	g.rpm = config.Attributes.Float64("rpm", 10.0)

	err := g.init(ctx)
	if err != nil {
		return nil, err
	}

	return g, nil
}

type oneAxis struct {
	name  string
	motor motor.Motor

	limitSwitchPins []string
	limitBoard      board.Board
	limitHigh       bool

	lengthMeters float64
	rpm          float64

	positionLimits []float64

	logger golog.Logger
}

func (g *oneAxis) init(ctx context.Context) error {
	ok, err := g.motor.PositionSupported(ctx)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("gantry motor needs to support position")
	}

	positionA, err := g.testLimit(ctx, true)
	if err != nil {
		return err
	}

	positionB, err := g.testLimit(ctx, false)
	if err != nil {
		return err
	}

	g.logger.Debugf("positionA: %0.2f positionB: %0.2f", positionA, positionB)

	g.positionLimits = []float64{positionA, positionB}

	return nil
}

func (g *oneAxis) testLimit(ctx context.Context, zero bool) (float64, error) {
	defer utils.UncheckedErrorFunc(func() error {
		return g.motor.Stop(ctx)
	})

	d := -1
	if !zero {
		d = 1
	}

	err := g.motor.GoFor(ctx, g.rpm, float64(d*10000))

	if err != nil {
		return 0, err
	}

	start := time.Now()
	for {
		hit, err := g.limitHit(ctx, zero)
		if err != nil {
			return 0, err
		}
		if hit {
			err = g.motor.Stop(ctx)
			if err != nil {
				return 0, err
			}
			break
		}

		elapsed := start.Sub(start)
		if elapsed > (time.Second * 15) {
			return 0, errors.New("gantry timed out testing limit")
		}

		if !utils.SelectContextOrWait(ctx, time.Millisecond*10) {
			return 0, ctx.Err()
		}
	}

	return g.motor.Position(ctx)
}

func (g *oneAxis) limitHit(ctx context.Context, zero bool) (bool, error) {
	offset := 0
	if !zero {
		offset = 1
	}
	pin := g.limitSwitchPins[offset]
	high, err := g.limitBoard.GPIOGet(ctx, pin)

	return high == g.limitHigh, err
}

// Position returns the position in meters
func (g *oneAxis) CurrentPosition(ctx context.Context) ([]float64, error) {
	pos, err := g.motor.Position(ctx)
	if err != nil {
		return nil, err
	}

	theRange := g.positionLimits[1] - g.positionLimits[0]
	x := g.lengthMeters * ((pos - g.positionLimits[0]) / theRange)

	g.logger.Debugf("oneAxis CurrentPosition %v -> %v", pos, x)

	return []float64{x}, nil
}

func (g *oneAxis) Lengths(ctx context.Context) ([]float64, error) {
	return []float64{g.lengthMeters}, nil
}

// position is in meters
func (g *oneAxis) MoveToPosition(ctx context.Context, positions []float64) error {
	if len(positions) != 1 {
		return fmt.Errorf("oneAxis gantry MoveToPosition needs 1 position, got: %v", positions)
	}

	if positions[0] < 0 || positions[0] > g.lengthMeters {
		return fmt.Errorf("oneAxis gantry position out of range, got %v max is %v", positions[0], g.lengthMeters)
	}

	theRange := g.positionLimits[1] - g.positionLimits[0]

	x := positions[0] / g.lengthMeters
	x = g.positionLimits[0] + (x * theRange)

	g.logger.Debugf("oneAxis SetPosition %v -> %v", positions[0], x)

	return g.motor.GoTo(ctx, g.rpm, x)
}

func (g *oneAxis) ModelFrame() *referenceframe.Model {
	m := referenceframe.NewModel()
	f, err := referenceframe.NewTranslationalFrame(
		g.name,
		[]bool{true},
		[]referenceframe.Limit{{0, g.lengthMeters}},
	)
	if err != nil {
		panic(fmt.Errorf("error creating frame, should be impossible %w", err))
	}
	m.OrdTransforms = append(m.OrdTransforms, f)

	return m
}

func (g *oneAxis) CurrentInputs(ctx context.Context) ([]referenceframe.Input, error) {
	res, err := g.CurrentPosition(ctx)
	if err != nil {
		return nil, err
	}
	return referenceframe.FloatsToInputs(res), nil
}
func (g *oneAxis) GoToInputs(ctx context.Context, goal []referenceframe.Input) error {
	return g.MoveToPosition(ctx, referenceframe.InputsToFloats(goal))
}
