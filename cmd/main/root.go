package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Root() {
	app := &cli.App{
		Name:  "glSync",
		Usage: "Sync Gitlab projects",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Path to a configuration file, defaults to $HOME/.config/glSync/config.json",
			},
		},
		Before: func(cCtx *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
