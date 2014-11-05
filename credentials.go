package main

import (
   "os"
   "github.com/mitchellh/go-homedir"
   "fmt"
   "io/ioutil"
   "encoding/json"
   "bufio"
)


var HOME_DIR, _ =  homedir.Dir()     //the users home directory where we'll put the credentials
var CREDENTIAL_FILE = ".atlas.json"  //the name of the credentials file


// Define API auth credentials structure
type Credentials struct {
   User string
   Key  string
}


// Get the users login credentials and save them to "~/.atlas.json" for the next time
func (c *Credentials) Query()  {
    // Prompt the user for his or her credentials
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Atlas user name: ")
	c.User, _ = reader.ReadString('\n')
	c.User = c.User[:len(c.User)-1]
	fmt.Print("Enter your API key: ")
	c.Key, _ = reader.ReadString('\n')
	c.Key = c.Key[:len(c.Key)-1]
}


// Save the credentials to ~/.atlas.json
func (c *Credentials) Save() {
   out, _ := json.Marshal(c)
   err := ioutil.WriteFile( HOME_DIR + "/" + CREDENTIAL_FILE, out, 0644)
   if err != nil {
      panic(err)
   }	
}


// retreive the credentials from ~/.atlas.json
func (c *Credentials) Load() error {	
   credsJSON, err := ioutil.ReadFile(HOME_DIR + "/" + CREDENTIAL_FILE)    	
   if err == nil {
      err = json.Unmarshal(credsJSON, &c)      
   } else {
      return err	
   }
   return nil
}


//do the login
func (c *Credentials) Login() {
	err := c.Load()
	if err != nil {
		c.Query()
		c.Save()
	}	
}


