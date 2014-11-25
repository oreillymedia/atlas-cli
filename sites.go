package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"github.com/toqueteos/webbrowser"
	"net/url"
	"io/ioutil"
	"net/http"
	"strconv"
	"encoding/json"
	"time"
)


type Sites struct {
	Id int `json:"id"`
	URL string `json:"url"`
	State string `json:"state"`
}


//Find the first build with an HTML build
func getBuildId(user *Credentials, project string) int {
	
	// Fetch the build data itself form the atlas api
	builds := &ProjectBuilds{}	
	builds.Get(user, project)
		
	// loop though each build  and build-> status till we find the first html build
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
	// if we don't find the build id then we need to error out
	if build_id == -1 {
		log.Fatal("Cannot find an html build for " + project)
	}
	return build_id
}


// Gets the status of the build with the given id
func get_state(id int) string {
	
	url := fmt.Sprintf("http://web-publisher.atlas.oreilly.com/deploy/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// Read the results from the build request
	body, err := ioutil.ReadAll(resp.Body)

	var s Sites
	err = json.Unmarshal(body, &s)
	
	return s.State
		
}


func (s *Sites) Publish(user *Credentials, c *cli.Context) {

	// Get the project name
	project := c.Args().First()
	if len(project) == 0 {
		log.Fatal("You must supply a project name")
	}
	
	fmt.Print("Working")

	// Get the visibility level
	bucket_type := "private"
	if c.Bool("public") {
		bucket_type = "public"
	}
	
	// get the build ID
	build_id := getBuildId(user,project)
	
	// Now hit the API endpoint to publish the build to sites
	resp, err := http.PostForm("http://web-publisher.atlas.oreilly.com/deploy",
		url.Values{
			"build_id": {strconv.Itoa(build_id)},
			"s3_path": {project},
			"bucket_type": {bucket_type},
			"token": {user.Key},
		})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	// Read the results from the build request
	body, err := ioutil.ReadAll(resp.Body)
	
	var response Sites
	json.Unmarshal(body, &response)
	
	// Now poll until the build is complete
	for {
		fmt.Print(".")
		state := get_state(response.Id)
		if (state == "complete") {
			break
		}
		time.Sleep(500 * time.Millisecond)		
	}

	s.Open(c)
	
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