package main

import (
   "log"
   "github.com/codegangsta/cli"
   "fmt"
)

type Grant struct {
	Email string
	OracleID string
}

func (g *Grant) ParseArgs(c *cli.Context ) {
	
	if len(c.Args()) != 2 {
		log.Fatal("You must supply an email address and an oracle id of the project to grant")
	}
	g.Email = c.Args().Get(0)
	g.OracleID = c.Args().Get(1)
	
}


func (g *Grant) Grant() {
	fmt.Printf("Granting %s permission to %s\n", g.Email, g.OracleID)
}