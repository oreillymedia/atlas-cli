package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)



// Defines the payload of the build API
type Builds struct {
	Build_url string `json:"build_url"`
	Message   string `json:"message"`
	Status    []struct {
		Format       string `json:"format"`
		Status       string `json:"status"`
		Download_url string `json:"download_url"`
		Message      struct {
			Error	string `json:"error"`
		} `json:"message"`
	} `json:"status"`
}

type ProjectBuilds []struct{
	Id        int    `json:"id"`
    CreatedAt string  `json:"created_at"`
    Status    []struct {
	   Format string `json:"format"`
    } `json:"status"`
}


// Hold the parameters for a build
type BuildArgs struct {
	Project string
	Formats string
	Branch  string
}


// This function returns the builds for the project
// It uses the structure defined in builds.go

func (builds *ProjectBuilds) Get(user *Credentials, project string) {
	
	qry := url.Values{
			"project":    {project},
			"auth_token":  {user.Key},
	}
	
	resp, err := http.Get("https://atlas.oreilly.com/api/builds?" + qry.Encode() )
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	// Read the results from the build request
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &builds)
	if err != nil {
		log.Fatal(err)
	}
	
}


func (args *BuildArgs) Parse(c *cli.Context) {

	// the project must be the first argument
	
	if len(c.String("project")) > 0 {
		args.Project = c.String("project")
	} else {
		args.Project = GetGitInfo()
	}


	//process the format flags they've requested
	formats := make([]string, 0)
	for _, f := range []string{"pdf", "html", "epub", "mobi"} {
		if c.Bool(f) {
			formats = append(formats, f)
		}
	}
	// if they haven't entered any formats, the build a PDF by default
	if len(formats) == 0 {
		formats = append(formats, "pdf")
	}
	args.Formats = strings.Join(formats, ",")

	// Now specifiy the branch
	args.Branch = "master"
	if len(c.String("branch")) > 0 {
		args.Branch = c.String("branch")
	}
}

// start a build for the project
// Returns the URL to poll if successful
func (builds *Builds) Start(c Credentials, a BuildArgs) {

	resp, err := http.PostForm("https://atlas.oreilly.com/api/builds",
		url.Values{
			"project":    {a.Project},
			"auth_token": {c.Key},
			"branch":     {a.Branch},
			"formats":    {a.Formats},
		})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the results from the build request
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &builds)

}

//start the build and then query it till it's done

func (builds *Builds) Build(c Credentials, a BuildArgs) {

	fmt.Printf("Building %s",a.Project)

	builds.Start(c, a)

	build_url := fmt.Sprintf("https://atlas.oreilly.com%s?auth_token=%s", builds.Build_url, c.Key)
	var build_status Builds

	// now poll the build and see how it's doing
	for {
		resp, err := http.Get(build_url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		// Read the results from the build request
		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &build_status)
		// now process the reuturn value.  We should stop polling when all status values
		// are either "completed" or "failed"
		completed_count := 0
		for _, s := range build_status.Status {
			if (s.Status == "completed") || (s.Status == "failed") {
				completed_count += 1
			}
		}
		if completed_count == len(build_status.Status) {
			break
		}
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()

	// Now print the build info
	for _, s := range build_status.Status {
		if s.Status == "completed" {
			fmt.Printf("%s => %s\n", s.Format, s.Download_url)
		} else {
			fmt.Printf("%s => %s\n %s\n", s.Format, "Failed to build", s.Message.Error)
		}
	}

}
