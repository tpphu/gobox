package gobox

import (
	"fmt"
	"os"

	"github.com/tpphu/gobox/logger"
	"github.com/tpphu/gobox/provider"
	"github.com/urfave/cli/v2"
)

// App is interface of gobox
type App interface {
	// Inject Env
	Load()
	Help()
	Migrate()
	Seed()
	Up()
	Down()
}

type application struct {
	cli.App
	Provider provider.Provider
	Log      *logger.Logger
}

func (a *application) Load() {

}

func (a *application) Up() {
	a.Log.Info("Application is up")
	a.Run(os.Args)
}

func (a *application) Down() {

}

func (a *application) Seed() {

}

func (a *application) Migrate() {

}

func (a *application) Help() {

}

func Default() App {
	fmt.Println("test")
	app := application{
		App: cli.App{
			Name:  "greet",
			Usage: "say a greeting",
			Action: func(c *cli.Context) error {
				fmt.Println("Greetings")
				return nil
			},
		},
		Provider: provider.NewProvider(),
		Log:      logger.New(),
	}
	app.Log.Out = os.Stdout
	
	return &app
}
