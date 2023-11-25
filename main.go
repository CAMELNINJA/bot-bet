package main

import (
	"os"

	"github.com/CAMELNINJA/bot-bet.git/cmd/bot"
	"github.com/urfave/cli"

	"log/slog"
)

var (
	cliApp     *cli.App
	configPath string
)

func init() {
	// Initialise a CLI app
	cliApp = cli.NewApp()
	cliApp.Name = "bet bot"
	cliApp.Usage = "bet bot"
	cliApp.Author = "Aidashev Kamil"
	cliApp.Email = "rfvbkm0220@gmail.com"
	cliApp.Version = "0.0.0"
	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config",
			Value:       "./config/local.env",
			Destination: &configPath,
		},
	}
}
func main() {
	// Set the CLI app commands
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start bot",
			Action: func(c *cli.Context) error {
				return bot.StartBot(configPath)
			},
		},
	}
	// Run the CLI app
	if err := cliApp.Run(os.Args); err != nil {
		slog.Error("error running cli app", err)
	}
}
