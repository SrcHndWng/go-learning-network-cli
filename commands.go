package main

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

type Commands struct {
	Flags []cli.Flag
}

func (c *Commands) Ns() cli.Command {
	return cli.Command{
		Name:  "ns",
		Usage: "Looks Up the NameServers for a Particular Host",
		Flags: c.Flags,
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
	}
}

func (c *Commands) Ip() cli.Command {
	return cli.Command{
		Name:  "ip",
		Usage: "Looks up the IP addresses for a particular host",
		Flags: c.Flags,
		Action: func(c *cli.Context) error {
			ip, err := net.LookupIP(c.String("host"))
			if err != nil {
				fmt.Println(err)
			}
			for i := 0; i < len(ip); i++ {
				fmt.Println(ip[i])
			}
			return nil
		},
	}
}
