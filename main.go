package main

import (
	"log"
	"os"

	"github.com/sijiramakun/seapick/recieve"
	"github.com/sijiramakun/seapick/send"
	"github.com/urfave/cli/v2"
)


var Commands = []*cli.Command{
	{
		Name:  "recieve",
		Usage: "wait and listen for a connection",
		Action: func(cCtx *cli.Context) error {
			Reciever.Receive()
			return nil
		},
	},
	{
		Name:  "send",
		Usage: "send a file to a peer",
		Flags: Sender.Flags,
		Action: func(cCtx *cli.Context) error {
			Sender.Send(cCtx)
			return nil
		},
	},
}

func main() {

	app := &cli.App{Commands: Commands, EnableBashCompletion: true}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
