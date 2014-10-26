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
)


// These are the main auth credentials
type Credentials struct {
   User string
   Key  string
}


var HOME_DIR, _ =  homedir.Dir()
var CREDENTIAL_FILE = ".atlas.json"


// Set the user credentials
func setCredentials(user,key string) {
	c := Credentials{user,key}
	out, _ := json.Marshal(c)

    err := ioutil.WriteFile( HOME_DIR + "/" + CREDENTIAL_FILE, out, 0644)
    if err != nil {
       panic(err)
    }

	fmt.Printf("Credentials written to %s.\n", HOME_DIR + "/" + CREDENTIAL_FILE)
}

// Get the credentials
func getCredentials() Credentials {
   var c Credentials
   // Grab the contents of the section
   creds, err := ioutil.ReadFile(HOME_DIR + "/" + CREDENTIAL_FILE)
   if err != nil {
      log.Fatal(err)
   }
   err = json.Unmarshal(creds, &c)
   return c
}


// Kicks off a build
func build() {
	fmt.Println("I'm building!")
}

// Reads the users login credentials and saves them in ~/.atlas
func login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Atlas user name: ")
	user, _ := reader.ReadString('\n')
	fmt.Print("Enter your API key: ")
	key, _ := reader.ReadString('\n')
	setCredentials(user[:len(user)-1],key[:len(key)-1])
}

// Displays the user's info
func whoami() {
	c := getCredentials()
	fmt.Printf("You are %s.\n", c.User)
}


// key actions are login and build
// login should ask you for your token and then save it to a file where you can get it later
// build should build -- it needs flags for format, project name, and branch


func main() {
	
   fmt.Println(HOME_DIR + "/" + CREDENTIAL_FILE)

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
	        login()
         },
      },   
      {
         Name:  "whoami",
         Usage: "Display your login/API credentials",
         Action: func(c *cli.Context) {
	        whoami()
         },
      },      {
	     Name: "build",
	     Usage: "Build a project",
         Flags: []cli.Flag {
            cli.StringFlag{
		       Name: "project, p",
		       Value: "project",
		       Usage: " project to build",
		    },
		    cli.StringFlag{
		       Name: "branch, b",
		       Value: "master",
		       Usage: "branch to build",
		    },
	     },
	     Action: func(c *cli.Context) {
		    build()
	     },
      },	
   }

  app.Run(os.Args)

}