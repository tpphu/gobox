package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tpphu/gobox/container"
	"github.com/tpphu/gobox/helper"
	"github.com/tpphu/gobox/logger"
	"github.com/tpphu/gobox/service"
	"github.com/tpphu/gobox/service/http"
	cli "github.com/urfave/cli/v2"
)

// App is main application of gobox
type App struct {
	*cli.App
	httpService *http.Http
	Services    []service.Runable
	Container   *container.Container
	Log         *logger.Logger
	Flag 		*AppFlagSet
	stopChan 	chan os.Signal
}

func (a *App) Init() {
	a.initHttpService()
	db, _ := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/gomay20?charset=utf8&parseTime=True&loc=Local")
  	a.Container.Set("db", db)
}

func (a *App) Run() {
	a.App.Run(os.Args)
	go func() {
		if err := a.runHttpService(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(a.stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	<-a.stopChan
	log.Println("Shutting down server...")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.httpService.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func (a *App) Shutdown() {

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
		stopChan: make(chan os.Signal),
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}
