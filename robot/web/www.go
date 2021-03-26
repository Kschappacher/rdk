package web

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/Masterminds/sprig"

	"github.com/edaniels/golog"
	"github.com/edaniels/gostream"
	"github.com/edaniels/gostream/codec/x264"

	"go.viam.com/robotcore/api"
	"go.viam.com/robotcore/board"
	"go.viam.com/robotcore/lidar"
	"go.viam.com/robotcore/robot"
	"go.viam.com/robotcore/robot/actions"
)

type robotWebApp struct {
	template *template.Template
	views    []gostream.View
	theRobot *robot.Robot
	logger   golog.Logger
}

func (app *robotWebApp) Init() error {
	_, thisFilePath, _, _ := runtime.Caller(0)
	thisDirPath, err := filepath.Abs(filepath.Dir(thisFilePath))
	if err != nil {
		return err
	}
	t, err := template.New("foo").Funcs(template.FuncMap{
		"jsSafe": func(js string) template.JS {
			return template.JS(js)
		},
		"htmlSafe": func(html string) template.HTML {
			return template.HTML(html)
		},
	}).Funcs(sprig.FuncMap()).ParseGlob(fmt.Sprintf("%s/*.html", thisDirPath))
	if err != nil {
		return err
	}
	app.template = t.Lookup("webappindex.html")
	return nil
}

func (app *robotWebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if true {
		err := app.Init()
		if err != nil {
			app.logger.Debugf("couldn't reload template: %s", err)
			return
		}
	}

	type View struct {
		JavaScript string
		Body       string
	}

	type Temp struct {
		Actions []string
		Views   []View
	}

	temp := Temp{
		Actions: actions.AllActionNames(),
	}

	for _, view := range app.views {
		htmlData := view.HTML()
		temp.Views = append(temp.Views, View{
			htmlData.JavaScript,
			htmlData.Body,
		})
	}

	err := app.template.Execute(w, temp)
	if err != nil {
		app.logger.Debugf("couldn't execute web page: %s", err)
	}
}

// ---------------

func InstallWebBase(mux *http.ServeMux, theBase api.Base, logger golog.Logger) {

	mux.Handle("/api/base", &apiCall{func(r *http.Request) (map[string]interface{}, error) {
		millisPerSec := 500.0 // TODO(erh): this is proably the wrong default
		if r.FormValue("speed") != "" {
			speed2, err := strconv.ParseFloat(r.FormValue("speed"), 64)
			if err != nil {
				return nil, err
			}
			millisPerSec = speed2
		}

		s := r.FormValue("stop")
		d := r.FormValue("distanceMillis")
		a := r.FormValue("angle")

		var err error

		if s == "t" || s == "true" {
			err = theBase.Stop(r.Context())
		} else if d != "" {
			d2, err2 := strconv.ParseInt(d, 10, 64)
			if err2 != nil {
				return nil, err2
			}

			err = theBase.MoveStraight(r.Context(), int(d2), millisPerSec, false)
		} else if a != "" {
			a2, err2 := strconv.ParseInt(a, 10, 64)
			if err2 != nil {
				return nil, err2
			}

			// TODO(erh): fix speed
			err = theBase.Spin(r.Context(), float64(a2), 64, false)
		} else {
			return nil, fmt.Errorf("no stop, distanceMillis, angle given")
		}

		if err != nil {
			return nil, fmt.Errorf("erorr moving %s", err)
		}

		return nil, nil

	}, logger})
}

type apiMethod func(r *http.Request) (map[string]interface{}, error)

type apiCall struct {
	theMethod apiMethod
	logger    golog.Logger
}

func (ac *apiCall) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := ac.theMethod(r)
	if err != nil {
		ac.logger.Warnf("error in api call: %s", err)
		res = map[string]interface{}{"err": err.Error()}
	}

	if res == nil {
		res = map[string]interface{}{"ok": true}
	}

	js, err := json.Marshal(res)
	if err != nil {
		ac.logger.Warnf("cannot marshal json: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js) //nolint
}

func InstallActions(mux *http.ServeMux, theRobot api.Robot, logger golog.Logger) {
	mux.Handle("/api/action", &apiCall{func(r *http.Request) (map[string]interface{}, error) {
		name := r.FormValue("name")
		action := actions.LookupAction(name)
		if action == nil {
			return nil, fmt.Errorf("unknown action name [%s]", name)
		}
		go action(theRobot)
		return map[string]interface{}{"started": true}, nil
	}, logger})
}

func InstallWebArms(mux *http.ServeMux, theRobot *robot.Robot, logger golog.Logger) {
	mux.Handle("/api/arm/MoveToPosition", &apiCall{func(r *http.Request) (map[string]interface{}, error) {
		req := &MoveToPositionRequest{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			return nil, err
		}

		arm := theRobot.ArmByName(req.Name)
		if arm == nil {
			return nil, fmt.Errorf("no arm with name (%s)", req.Name)
		}

		return nil, arm.MoveToPosition(req.To)
	}, logger})

	mux.Handle("/api/arm/MoveToJointPositions", &apiCall{func(r *http.Request) (map[string]interface{}, error) {
		req := &MoveToJointPositionsRequest{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			return nil, err
		}

		arm := theRobot.ArmByName(req.Name)
		if arm == nil {
			return nil, fmt.Errorf("no arm with name (%s)", req.Name)
		}

		return nil, arm.MoveToJointPositions(req.To)
	}, logger})

}

func InstallWebGrippers(mux *http.ServeMux, theRobot *robot.Robot, logger golog.Logger) {
	mux.Handle("/api/gripper", &apiCall{func(r *http.Request) (map[string]interface{}, error) {
		name := r.FormValue("name")
		gripper := theRobot.GripperByName(name)
		if gripper == nil {
			return nil, fmt.Errorf("no gripper with that name %s", name)
		}

		var err error
		res := map[string]interface{}{}

		action := r.FormValue("action")
		switch action {
		case "open":
			err = gripper.Open()
		case "grab":
			var grabbed bool
			grabbed, err = gripper.Grab()
			res["grabbed"] = grabbed
		default:
			err = fmt.Errorf("bad action: (%s)", action)
		}

		if err != nil {
			return nil, fmt.Errorf("gripper error: %s", err)
		}
		return res, nil
	}, logger})
}

func InstallSimpleCamera(mux *http.ServeMux, theRobot *robot.Robot, logger golog.Logger) {
	theFunc := func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		camera := theRobot.CameraByName(name)
		if camera == nil {
			http.Error(w, "bad camera name", http.StatusBadRequest)
			return
		}

		img, release, err := camera.Next(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer release()

		//TODO(erh): choice of encoding

		w.Header().Set("Content-Type", "image/jpeg")
		err = jpeg.Encode(w, img, nil)
		if err != nil {
			logger.Debugf("error encoding jpeg: %s", err)
		}
	}
	mux.HandleFunc("/api/camera", theFunc)
	mux.HandleFunc("/api/camera/", theFunc)

}

func InstallStatus(mux *http.ServeMux, theRobot *robot.Robot, logger golog.Logger) {
	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		status, err := theRobot.Status()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		enc := json.NewEncoder(w)
		err = enc.Encode(status)
		if err != nil {
			logger.Infof("failed to encode status %s", err)
		}
	})
}

func installBoard(mux *http.ServeMux, b board.Board, logger golog.Logger) {
	cfg := b.GetConfig()

	mux.Handle("/api/board/"+cfg.Name+"/motor", &apiCall{func(r *http.Request) (map[string]interface{}, error) {
		name := r.FormValue("name")
		theMotor := b.Motor(name)
		if theMotor == nil {
			return nil, fmt.Errorf("unknown motor: %s", r.FormValue("name"))
		}

		speed, err := strconv.ParseFloat(r.FormValue("s"), 64)
		if err != nil {
			return nil, err
		}

		rVal := 0.0
		if r.FormValue("r") != "" {
			rVal, err = strconv.ParseFloat(r.FormValue("r"), 64)
			if err != nil {
				return nil, err
			}
		}

		if rVal == 0 {
			return map[string]interface{}{}, theMotor.Go(board.DirectionFromString(r.FormValue("d")), byte(speed))
		}

		return map[string]interface{}{}, theMotor.GoFor(board.DirectionFromString(r.FormValue("d")), speed, rVal)
	}, logger})

	mux.Handle("/api/board/"+cfg.Name+"/servo", &apiCall{func(r *http.Request) (map[string]interface{}, error) {
		name := r.FormValue("name")
		theServo := b.Servo(name)
		if theServo == nil {
			return nil, fmt.Errorf("unknown servo: %s", r.FormValue("name"))
		}

		angle, err := strconv.ParseInt(r.FormValue("angle"), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("need to specify an angle to move to")
		}

		return nil, theServo.Move(uint8(angle))

	}, logger})
}

func InstallBoards(mux *http.ServeMux, theRobot *robot.Robot, logger golog.Logger) {
	for _, name := range theRobot.BoardNames() {
		installBoard(mux, theRobot.BoardByName(name), logger)
	}
}

// ---------------

func allSourcesToDisplay(theRobot *robot.Robot) ([]gostream.ImageSource, []string) {
	sources := []gostream.ImageSource{}
	names := []string{}

	for _, name := range theRobot.CameraNames() {
		cam := theRobot.CameraByName(name)
		cmp := theRobot.GetConfig().FindComponent(name)
		if cmp.Attributes.GetBool("hide", false) {
			continue
		}

		sources = append(sources, cam)
		names = append(names, name)
	}

	for _, name := range theRobot.LidarDeviceNames() {
		device := theRobot.LidarDeviceByName(name)
		cmp := theRobot.GetConfig().FindComponent(name)
		if cmp.Attributes.GetBool("hide", false) {
			continue
		}

		source := lidar.NewImageSource(image.Point{600, 600}, device)

		sources = append(sources, source)
		names = append(names, name)
	}

	return sources, names
}

// returns a closer func to be called when done
func InstallWeb(ctx context.Context, mux *http.ServeMux, theRobot *robot.Robot, options Options, logger golog.Logger) (func(), error) {
	displaySources, displayNames := allSourcesToDisplay(theRobot)
	views := []gostream.View{}
	var autoCameraTiler *gostream.AutoTiler

	if len(displaySources) != 0 {
		if options.AutoTile {
			config := x264.DefaultViewConfig
			view, err := gostream.NewView(config)
			if err != nil {
				return nil, err
			}
			view.SetOnClickHandler(func(x, y int) {
				logger.Debugw("click", "x", x, "y", y)
			})
			views = append(views, view)

			tilerHeight := 480 * len(displaySources)
			autoCameraTiler = gostream.NewAutoTiler(640, tilerHeight)
			autoCameraTiler.SetLogger(logger)

		} else {
			for idx := range displaySources {
				config := x264.DefaultViewConfig
				config.StreamNumber = idx
				config.StreamName = displayNames[idx]
				view, err := gostream.NewView(config)
				if err != nil {
					return nil, err
				}
				view.SetOnClickHandler(func(x, y int) {
					logger.Debugw("click", "x", x, "y", y)
				})
				views = append(views, view)
			}
		}
	}

	app := &robotWebApp{views: views, theRobot: theRobot, logger: logger}
	err := app.Init()
	if err != nil {
		return nil, err
	}

	// install routes
	for _, name := range theRobot.BaseNames() {
		InstallWebBase(mux, theRobot.BaseByName(name), logger)
	}

	InstallWebArms(mux, theRobot, logger)

	InstallWebGrippers(mux, theRobot, logger)

	InstallActions(mux, theRobot, logger)

	InstallSimpleCamera(mux, theRobot, logger)

	InstallBoards(mux, theRobot, logger)

	InstallStatus(mux, theRobot, logger)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("robot/web/static/"))))
	mux.Handle("/", app)

	for _, view := range views {
		handler := view.Handler()
		mux.Handle("/"+handler.Name, handler.Func)
	}

	// start background threads

	if autoCameraTiler != nil {
		for _, src := range displaySources {
			autoCameraTiler.AddSource(src)
		}
		waitCh := make(chan struct{})
		go func() {
			close(waitCh)
			gostream.StreamNamedSource(ctx, autoCameraTiler, "Cameras", views[0])
		}()
		<-waitCh
	} else {
		for idx, view := range views {
			waitCh := make(chan struct{})
			go func() {
				close(waitCh)
				gostream.StreamNamedSource(ctx, displaySources[idx], displayNames[idx], view)
			}()
			<-waitCh
		}
	}

	return func() {
		for _, view := range views {
			view.Stop()
		}
	}, nil

}

// ---

/*
helper if you don't need to customize at all
*/
func RunWeb(theRobot *robot.Robot, options Options, logger golog.Logger) error {
	defer theRobot.Close(context.Background())
	mux := http.NewServeMux()

	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	webCloser, err := InstallWeb(cancelCtx, mux, theRobot, options, logger)
	if err != nil {
		return err
	}

	if options.Pprof {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	const port = 8080
	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        mux,
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancelFunc()
		webCloser()
		if err := httpServer.Shutdown(context.Background()); err != nil {
			theRobot.Logger().Errorw("error shutting down", "error", err)
		}
	}()

	theRobot.Logger().Debugw("going to listen", "addr", fmt.Sprintf("http://localhost:%d", port), "port", port)
	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		theRobot.Logger().Fatal(err)
	}

	return nil
}
