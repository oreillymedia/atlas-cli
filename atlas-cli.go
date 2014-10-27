package main

import (
   "os"
   "github.com/codegangsta/cli"
   "fmt"
   "bufio"
   "encoding/json"
   "github.com/mitchellh/go-homedir"
   "io/ioutil"
   "log"
   "time"
   "net/http"
)


// Define API auth credentials structure
type Credentials struct {
   User string
   Key  string
}

// Defines the payload of the build API
type Builds struct {
	Build_url string `json:"build_url"`
	Status []struct {
		Format         string `json:"format"`
		Status         string `json:"status"`
		Download_url   string `json:"download_url"`
		Message        string `json:"message"`
	} `json:"status"`
}


// Declare package level vairables

var HOME_DIR, _ =  homedir.Dir()     //the users home directory where we'll put the credentials
var CREDENTIAL_FILE = ".atlas.json"  //the name of the credentials file
var atlas_user Credentials           //the data structure w/the users API credentials


// Get the users login credentials and save them to "~/.atlas.json" for the next time
func login() Credentials {

    // Prompt the user for his or her credentials
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Atlas user name: ")
	user, _ := reader.ReadString('\n')
	fmt.Print("Enter your API key: ")
	key, _ := reader.ReadString('\n')
	
	// Save the credentials to a ~/.atlas.json
	c := Credentials{user[:len(user)-1],key[:len(key)-1]}
	out, _ := json.Marshal(c)
    err := ioutil.WriteFile( HOME_DIR + "/" + CREDENTIAL_FILE, out, 0644)
    if err != nil {
       panic(err)
    }

    //let the user know!
	fmt.Printf("Credentials written to %s.\n", HOME_DIR + "/" + CREDENTIAL_FILE)
	
	return c
	
}


// Kicks off a build
func build( c *cli.Context ) {
	fmt.Printf("I'm building a file for %s...\n", atlas_user.User)
	fmt.Println(c.String("branch"))
	for i:=1; i < 13; i++ {
       fmt.Print(".")	
	   time.Sleep(250 * time.Millisecond)
    }
	fmt.Println()

	fmt.Printf("Now I'm done.\n")
}

// thanks for this great post: http://www.codingcookies.com/2013/03/21/consuming-json-apis-with-go/
func getBuildStatus(url string) (Builds, error) {
    // At this point we're done 
    var builds Builds

    // Build the request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
      return builds, err
    }
    // Send the request via a client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
      return builds, err
    }
    // Defer the closing of the body
    defer resp.Body.Close()
    // Read the content into a byte array
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return builds, err
    }

    err = json.Unmarshal(body, &builds)
    return builds, err
}



// List all the builds in a project
func list() {
	fmt.Println("I'm listing your builds!")
	url := fmt.Sprintf("https://atlas.oreilly.com/api/builds/1?auth_token=%s", atlas_user.Key)
	builds, _ := getBuildStatus(url)
	fmt.Println(builds)
}


func main() {
	
   // First we try to fetch the users credentials from "~/.atlas.json".  If they are in there, we use those.
   //  Otherwise, we prompt the user and set them
   credsJSON, err := ioutil.ReadFile(HOME_DIR + "/" + CREDENTIAL_FILE)
   if err == nil {
      err = json.Unmarshal(credsJSON, &atlas_user)
      // doublecheck to make sure nothing went wrong
      if err != nil {
	     log.Fatal(err)
      }
   } else {
      atlas_user = login()
   }
  
   // Start the "real" program

   app := cli.NewApp()
   app.Name = "atlas-cli"
   app.Usage = "Atlas commandline API!"
   app.Action = func(c *cli.Context) {
     println("Hello friend!")
   }

   app.Commands = []cli.Command{
      {
         Name:  "login",
         Usage: "Provide your login/API credentials",
         Action: func(c *cli.Context) {
	        atlas_user = login()
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
	     Name: "build",
	     Usage: "Build a project",
         Flags: []cli.Flag {
            cli.StringFlag{
		       Name: "project, p",
		       Usage: " project to build",
		    },
            cli.BoolFlag{
		       Name: "pdf",
		       Usage: " build a pdf",
		    },
            cli.BoolFlag{
		       Name: "html",
		       Usage: " build html format",
		    },
		    cli.StringFlag{
		       Name: "branch, b",
		       Value: "master",
		       Usage: "branch to build",
		    },
	     },
	     Action: func(c *cli.Context) {
		    build(c)
	     },
      },
      {
	     Name: "list",
	     Usage: "List builds of a project",
         Action: func(c *cli.Context) {
            list()
         },
	  },
	
   }

  app.Run(os.Args)

}