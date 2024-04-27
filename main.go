package main

import (
	"fmt"
	"github.com/sijiramakun/seapick/recieve"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var Commands = []*cli.Command{
	{
		Name:  "recieve",
		Usage: "send a file to a peer",
		Action: func(cCtx *cli.Context) error {
			Reciever.Receive()
			return nil
		},
	},
	{
		Name:  "send",
		Usage: "wait and listen for a connection",
		Subcommands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a new template",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("new task template: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "remove an existing template",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("removed task template: ", cCtx.Args().First())
					return nil
				},
			},
		},
	},
}

func main() {

	app := &cli.App{Commands: Commands, EnableBashCompletion: true}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
