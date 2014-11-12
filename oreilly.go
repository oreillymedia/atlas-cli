package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

// Declare package level vairables

func main() {

	// Start the "real" program

	app := cli.NewApp()
	app.Name = "oreilly"
	app.Usage = "OºReilly command line API"
	app.Version = "0.0.5-alpha"
	app.Action = func(c *cli.Context) {
		fmt.Println("Nothing to do.  Try `help` or `-h` to see what's possible.")
	}

	atlas_user := &Credentials{}
	atlas_user.Login()

	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Set your login/API credentials",
			Action: func(c *cli.Context) {
				atlas_user.Query()
				atlas_user.Save()
			},
		},
		{
			Name:  "whoami",
			Usage: "Display your login/API credentials",
			Action: func(c *cli.Context) {
				fmt.Printf("You are %s.\n", atlas_user.User)
			},
		},
		{
			Name: "atlas",
			Usage: "Work with Atlas projects",
			Subcommands: []cli.Command{
				{
					Name:  "build",
					Usage: "Build a project",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "pdf",
							Usage: " build a pdf",
						},
						cli.BoolFlag{
							Name:  "html",
							Usage: " build html format",
						},
						cli.BoolFlag{
							Name:  "epub",
							Usage: " build epub format",
						},
						cli.BoolFlag{
							Name:  "mobi",
							Usage: " build mobi format",
						},
						cli.StringFlag{
							Name:  "branch, b",
							Value: "master",
							Usage: "branch to build",
						},
					},
					Action: func(c *cli.Context) {
						args := &BuildArgs{}
						args.Parse(c)
						build := &Builds{}
						build.Build(*atlas_user, *args)
					},
				},
			},
		},
		{
			Name: "product",
			Usage: "Product metadata and ownership server (requires admin access)",
			Subcommands: []cli.Command{
				{
					Name: "grant",		
					Usage: "Grant a user a permission on a project ID (admin only)",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "oracle_id, id",
							Usage: "Product ID (oracle id)",
						},
						cli.StringFlag{
							Name: "email, e",
							Usage: "Email address of grantee",
						},
					},
					Action: func(c *cli.Context) {
						p := &Product{}
						p.Grant(c)
					},
				},
				{
					Name: "find",
					Usage: "Search for project info",
					Action: func(c *cli.Context){
						f := &Product{}
						f.Find(c)
					},
				},
			},
		},
		{
			Name: "sites",
			Usage: "Publish an Atlas project to sites.oreilly.com",
			Subcommands: []cli.Command{
				{
					Name: "open",
					Usage: "Open a site",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name: "public",
							Usage: "Make project public",
						},
					},
					Action: func(c *cli.Context){
						s := &Sites{}
						s.Open(c)
					},
				},
				{
					Name: "publish",
					Usage: "Publish a site",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name: "public",
							Usage: "Make project public",
						},
					},
					Action: func(c *cli.Context){
						s := &Sites{}
						s.Publish(c)
					},
				},
			},
		},
	}

	app.Run(os.Args)

}
