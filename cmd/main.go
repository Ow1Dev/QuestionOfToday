package cmd

import "github.com/urfave/cli/v2"

type AppVersion struct {
	Version string
	Extra   string
}

func NewMainApp(appVer AppVersion) *cli.App {
	app := cli.NewApp()
	app.Name = "QuestionOfToday"
	app.HelpName = "QuestionOfToday"
	app.Usage = "A painless self-hosted Git service"
	app.Description = `QuestionOfToday program contains "web" and other subcommands. If no subcommand is given, it starts the web server by default. Use "web" subcommand for more web server arguments, use other subcommands for other purposes.`
	app.Version = appVer.Version + appVer.Extra
	app.EnableBashCompletion = true

	subCmdWithConfig := []*cli.Command{
		CmdWeb,
	}

  app.DefaultCommand = CmdWeb.Name
  app.Commands = append(app.Commands, subCmdWithConfig...)

	return app
}

func RunMainApp(app *cli.App, args ...string) error {
	err := app.Run(args)

	if err == nil {
		return nil
	}

	return err
}
