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
   "net/url"
   "strings"
)


// Define API auth credentials structure
type Credentials struct {
   User string
   Key  string
}

// Defines the payload of the build API
type Builds struct {
	Build_url string `json:"build_url"`
	Message string `json:message`
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


// kick off a build for the project
func buildAndPoll(c *cli.Context) {
	
   // the project must be the first argument
   project := c.Args().First()
   if len(project) == 0 {
      log.Fatal("You must supply a project name.")
   }
	
   //process the format flags they've requested
   formats := make([]string, 0)
   for _,f := range []string{"pdf", "html", "epub", "mobi"} {
      if c.Bool(f) { formats = append(formats, f) } 		
   }
   // if they haven't entered any formats, the build a PDF by default
   if len(formats) == 0 {formats = append(formats, "pdf") }


   fmt.Print("Working")

		
   resp, err := http.PostForm("https://atlas.oreilly.com/api/builds",
      url.Values{
        "project": {project}, 
        "auth_token": {atlas_user.Key},
        "branch": {c.String("branch")},
        "formats": {strings.Join(formats,",")},
   })
   if err != nil {
      log.Fatal(err)
   }
   defer resp.Body.Close()


   // Read the results from the build request
   body, err := ioutil.ReadAll(resp.Body)
   var builds Builds
   err = json.Unmarshal(body, &builds)

   build_url := fmt.Sprintf("https://atlas.oreilly.com%s?auth_token=%s", builds.Build_url, atlas_user.Key)

   // now poll the build and see how it's doing
   var build_status Builds
   for {
	   resp, err = http.Get(build_url);
	   if err != nil {
	      log.Fatal(err)
	   }
	   defer resp.Body.Close()
	   // Read the results from the build request
	   body, err = ioutil.ReadAll(resp.Body)
	   err = json.Unmarshal(body, &build_status)
	   // now process the reuturn value.  We should stop polling when all status values
	   // are either "completed" or "failed"
	   completed_count := 0
       for _,s := range build_status.Status {
          if (s.Status == "completed") || (s.Status == "failed") {
             completed_count += 	1
          }
       } 
       if completed_count == len(build_status.Status) {
          break;
       }
       fmt.Print(".")     
	   time.Sleep(500 * time.Millisecond)	
   }
   fmt.Println();

   // Now print the build info
   for _,s := range build_status.Status {
      if s.Status == "completed" {
         fmt.Printf("%s => %s\n", s.Format, s.Download_url)
      }
   }

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
            cli.BoolFlag{
		       Name: "epub",
		       Usage: " build epub format",
		    },
            cli.BoolFlag{
		       Name: "mobi",
		       Usage: " build mobi format",
		    },
            cli.StringFlag{
		       Name: "branch, b",
		       Value: "master",
		       Usage: "branch to build",
		    },
	     },
	     Action: func(c *cli.Context) {
		    buildAndPoll(c)
	     },
      },
   }

  app.Run(os.Args)

}