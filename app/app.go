package app

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tpphu/gobox/container"
	"github.com/tpphu/gobox/helper"
	"github.com/tpphu/gobox/logger"
	"github.com/tpphu/gobox/service"
	"github.com/tpphu/gobox/service/http"
	cli "github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// App is main application of gobox
type App struct {
	*cli.App
	httpService *http.Http
	Services    []service.Runable
	Container   *container.Container
	Log         *logger.Logger
	chSignal chan os.Signal
	Flag 		*AppFlagSet
	mtx *sync.Mutex
}
/*
func (a *App) Init() {
	//a.initHttpService()
	db, _ := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/gomay20?charset=utf8&parseTime=True&loc=Local")
  	a.Container.Set("db", db)
}*/

func (a *App) startRun() {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	runService := func(s service.Runable) {
		if err := s.Run(); err != nil {
			a.Shutdown()
		}
	}
	log := a.Log
	for _, s := range a.Services {
		runS, ok := s.(service.Runable)
		if !ok {
			continue
		}
		go func(runS service.Runable) {
			log.Info("start...")
			runService(runS)
		}(runS)
	}

}

func (a *App) Run() {
	a.App.Run(os.Args)
	signal.Notify(a.chSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	log := a.Log
	a.startRun()

	for sig := range a.chSignal {
		switch sig {
		case syscall.SIGHUP:
			break
		default:
			goto SHUTDOWN
		}
	}

SHUTDOWN:
	// starting shutting down progress...
	log.Infof("Server shutting down...")
	a.Shutdown()
}

func (a *App) Shutdown() {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	for i := len(a.Services) - 1; i >= 0; i-- {
		s := a.Services[i]
		if runS, ok := s.(service.Runable); ok {
			// call for shutdown
			if err := runS.Shutdown(); err != nil {
				a.Log.Errorf("Server Shutdown failed: %v", err)
			}
		}
	}
}

func (a *App) AddService(s service.Runable) {
	a.Services = append(a.Services, s)
}

func (a *App) Provide(entry interface{}) {
	s, _ := helper.GetFieldTag(entry, "DB", "inject")
	if s != "" {
		helper.SetField(entry, "DB", a.Container.Get(s))
	}	
	a.Log.Info(s)
}

func NewApp(opts ...Option) *App {
	log := logger.New()
	log.Out = os.Stdout
	var app *App
	app = &App{
		App: &cli.App{
			Name:                 "gobox",
			Usage:                "a simple gobox application",
			EnableBashCompletion: true,
			Commands: []*cli.Command{
				{
					Name:  "up",
					Usage: "Up application",
					Action: func(c *cli.Context) error {
						return app.Up(c)
					},
				},
				{
					Name:  "down",
					Usage: "Down application",
					Action: func(c *cli.Context) error {
						return app.Down(c)
					},
				},
				{
					Name:  "seed",
					Usage: "Seed data for application",
					Action: func(c *cli.Context) error {
						return app.Seed(c)
					},
				},
				{
					Name:  "migrate",
					Usage: "Migrate data for application",
					Action: func(c *cli.Context) error {
						return app.Migrate(c)
					},
				},
				{
					Name:  "outenv",
					Usage: "Export env of application",
					Action: func(c *cli.Context) error {
						return app.ExportEnv("Example")
					},
				},
			},
		},
		Container: container.NewContainer(),
		Log:       log,
		chSignal: make(chan os.Signal, 1),
		mtx:       &sync.Mutex{},
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}
