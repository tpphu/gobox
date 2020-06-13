package app

import (
	"flag"
	"github.com/urfave/cli/v2"
	"os"
)

func (a *App) Up(ctx *cli.Context) error {
	a.Log.Info("Application is up")
	return nil
}

func (a *App) Down(ctx *cli.Context) error {
	a.Log.Info("Application is down")
	return nil
}

func (a *App) Seed(ctx *cli.Context) error {
	a.Log.Info("Application is seeding data")
	return nil
}

func (a *App) Migrate(ctx *cli.Context) error {
	a.Log.Info("Application is migrating schema")
	return nil
}

func (a *App) ExportEnv(name string) error {
	a.Flag = newFlagSet(name, flag.CommandLine)
	a.Flag.Parse(name)
	os.Exit(1)
	return nil
}
