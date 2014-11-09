package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
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
	fmt.Printf("Opening %s\n", project)
}