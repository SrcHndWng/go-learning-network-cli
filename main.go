package main

import (
	"log"
	"os"

	"github.com/SrcHndWng/go-learning-network-cli/commands"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers!"

	// Define flags
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.co.jp",
		},
	}

	// Create commands
	app.Commands = commands.Create(flags)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
