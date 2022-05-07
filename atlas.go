package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// Declare package level vairables

func main() {

	// Start the "real" program

	app := cli.NewApp()
	app.Name = "atlas"
	app.Usage = "OºReilly Atlas command line tool"
	app.Version = "0.1.0-alpha"
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
			Name:  "info",
			Usage: "Display info about your Atlas project based on the git config file",
			Action: func(c *cli.Context) {
				fmt.Printf("Your project is %s\n", GetGitInfo())
			},
		},
		{
			Name:  "build",
			Usage: "Build a project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "project, p",
					Usage: "Project name",
				},
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
		{
			Name:  "open",
			Usage: "Open a site",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "project, p",
					Usage: "Project name",
				},
				cli.BoolFlag{
					Name:  "public",
					Usage: "Make project public",
				},
			},
			Action: func(c *cli.Context) {
				s := &Sites{}
				s.Open(c)
			},
		},
		{
			Name:  "publish",
			Usage: "Publish a site",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "project, p",
					Usage: "Project name",
				},
				cli.BoolFlag{
					Name:  "public",
					Usage: "Make project public",
				},
			},
			Action: func(c *cli.Context) {
				s := &Sites{}
				s.Publish(atlas_user, c)
			},
		},
	}

	app.Run(os.Args)

}
