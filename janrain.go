package main

import (
	"log"
	"encoding/json"
	"net/http"
	"net/url"
	"io/ioutil"
)


// Define a janrain user
type JanRainUser struct {
	Results []struct {
		UUID string `json:"uuid"`
		Email string `json:"email"`
	} `json:"results"`
}


func (j *JanRainUser) Find(email string) {
	
	c := &Credentials{}
	c.Load()
	
	//find the users UUID in JanRain
    query := url.Values{
		"client_id":    {c.JanrainClientID},
		"client_secret": {c.JanrainClientSecret},
		"type_name": {"user"},
		"filter": {"email='" + email + "'"},
	}	
	
	resp, err := http.PostForm("https://oreilly.janraincapture.com/entity.find", query)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &j)
	
}