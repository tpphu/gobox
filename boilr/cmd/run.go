package cmd

import cli "github.com/spf13/cobra"

// Root contains the root cli-command containing all the other commands.
var Cmd = &cli.Command{
	Use: "boilerplate",
}

// Run executes the cli-command root.
func Run() {
	Init.PersistentFlags().StringP("path", "p", "", "Destination path.")
	Init.PersistentFlags().StringP("template", "t", "gin", "Template to use (gin, grpc)")
	Init.PersistentFlags().StringP("service-name", "s", "example", "Name of service")
	Init.PersistentFlags().StringP("resource", "r", "logfile-redis-mgo", "Resource to use: logfile, redis, mgo")
	Init.PersistentFlags().BoolP("force", "f", false, "Recreate directories if they exist")

	Cmd.AddCommand(Init)
	Cmd.Execute()
}
