package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"github.com/toqueteos/webbrowser"
)


type Sites struct {

}


func (s *Sites) Publish(c *cli.Context) {
	project := c.Args().First()
	if len(project) == 0 {
		log.Fatal("You must supply a project name")
	}
	fmt.Printf("Publishing %s once Rune makes an API\n", project)
}

func (s *Sites) Open(c *cli.Context) {
	project := c.Args().First()
	if len(project) == 0 {
		log.Fatal("You must supply a project name")
	}
		
	url := fmt.Sprintf("http://orm-static-site-proxy.herokuapp.com/%s", project)
	if c.Bool("public") {
		url = fmt.Sprintf("http://sites.oreilly.com/%s", project)		
	}
	fmt.Printf("Opening %s\n", url)
	webbrowser.Open(url)
}