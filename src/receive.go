package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func receiveMode(){
	app := &cli.App{
		Commands: []*cli.Command{
		  {
			Name:        "receive",
			Aliases:     []string{"r"},
			Usage:       "receive a files mode",
			Subcommands: []*cli.Command{
			  {
				Name:  "channel",
				Aliases: []string{"ch"},
				Usage: "channel number",
				Action: func(c *cli.Context) error {
				  fmt.Println("the channel number is: ", c.Args().First())
				  return nil
				},
			  },
			},
		  },
		},
	  }
	
	err := app.Run(os.Args)
	if err != nil {
	log.Fatal(err)
	}
}