package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers!"

	// Define flags
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.co.jp",
		},
	}

	// Create commands
	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks Up the NameServers for a Particular Host",
			Flags: myFlags,
			// the action, or code that will be executed when
			// we execute our `ns` command
			Action: func(c *cli.Context) error {
				// a simple lookup function
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				// we log the results to our console
				// using a trusty fmt.Println statement
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
