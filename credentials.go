package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
)

var HOME_DIR, _ = homedir.Dir()     //the users home directory where we'll put the credentials
var CREDENTIAL_FILE = ".atlas.json" //the name of the credentials file

// Define API auth credentials structure
type Credentials struct {
	User string
	Key  string
	JanrainEmail string
	JanrainClientID string
	JanrainClientSecret string
}

// Fetch a value
func prompt(s string) string {
	// Prompt the user for his or her credentials
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", s)
	val, _ := reader.ReadString('\n')
	val = val[:len(val)-1]
	return val
}

// Get the users login credentials and save them to "~/.atlas.json" for the next time
func (c *Credentials) Query() {
	c.User = prompt("Enter your Atlas user name")
	c.Key = prompt("Enter your Atlas key")
	c.JanrainEmail = prompt("Enter your JanRain email address (hit enter to skip)")
	c.JanrainClientID = prompt("Enter your JanRain client ID (hit enter to skip)")
	c.JanrainClientSecret = prompt("Enter your JanRain client secret (hit enter to skip)")
}

// Save the credentials to ~/.atlas.json
func (c *Credentials) Save() {
	out, _ := json.Marshal(c)
	err := ioutil.WriteFile(HOME_DIR+"/"+CREDENTIAL_FILE, out, 0644)
	if err != nil {
		panic(err)
	}
}

// retreive the credentials from ~/.atlas.json
func (c *Credentials) Load() error {
	credsJSON, err := ioutil.ReadFile(HOME_DIR + "/" + CREDENTIAL_FILE)
	if err == nil {
		err = json.Unmarshal(credsJSON, &c)
		if err != nil { return err }
	} else {
		return err
	}
	return err
}

//do the login
func (c *Credentials) Login() {
	err := c.Load()
	if err != nil {
		c.Query()
		c.Save()
	}
}
