package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"github.com/toqueteos/webbrowser"
//	"net/url"
//	"io/ioutil"
//	"net/http"
//	"encoding/json"
)


type Sites struct {

}




func (s *Sites) Publish(user *Credentials, c *cli.Context) {
	project := c.Args().First()
	if len(project) == 0 {
		log.Fatal("You must supply a project name")
	}
	
	builds := &ProjectBuilds{}
	
	builds.Get(user, project)
	
	fmt.Printf("Publishing %s for %s once Rune makes an API\n", project, user.User)
	
	//Find the first build with an HTML build
	build_id := -1
	for _, b := range *builds {
		for _,s := range b.Status {
			//If we've found an HTML build, then set the build ID and break
			if s.Format == "html" {
				build_id = b.Id
				break
			}
		}
		//If we've found the HTML build, then break
		if build_id != -1 {
			break
		}
	}
	fmt.Println(build_id)
	
}

func (s *Sites) Open(c *cli.Context) {
	project := c.Args().First()
	if len(project) == 0 {
		log.Fatal("You must supply a project name")
	}
		
	url := fmt.Sprintf("http://orm-static-site-proxy.herokuapp.com/%s/ch01.html", project)
	if c.Bool("public") {
		url = fmt.Sprintf("http://sites.oreilly.com/%s/ch01.html", project)		
	}
	fmt.Printf("Opening %s\n", url)
	webbrowser.Open(url)
}