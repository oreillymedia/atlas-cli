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
)


// These are the main auth credentials
type Credentials struct {
   User string
   Key  string
}


var HOME_DIR, _ =  homedir.Dir()
var CREDENTIAL_FILE = ".atlas.json"
var atlas_user Credentials


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

// List all the builds in a project
func list() {
	fmt.Println("I'm listing your builds...")
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