package main

import (
	"log"
	"github.com/codegangsta/cli"
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
)

//Define a Product interfacce that will mostly have methods
type Product struct {
	Data []struct {
		EditionLabel string `json:"editionLabel"`
		OracleId string `json:"oracleId"`
		Title string `json:"title"`
		
	} `json:"data"`
}


// Define a janrain user
type JanRainUser struct {
	Results []struct {
		UUID string `json:"uuid"`
		Email string `json:"email"`
	} `json:"results"`
}

// pads a string to be N chars long
func pad(s string, N int) string {
	if len(s) > N {
		return s[:N]
	} else {
		return s + strings.Repeat(" ", N - len(s))
	}
	
}


func (j *JanRainUser) Find(c Credentials, email string) {
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

func (p *Product) Grant(c Credentials, args *cli.Context) {
	//product_id := args.String("oracle_id")
	email := args.String("email")
    
	fmt.Println("Working...")
	janrain := JanRainUser{}
	janrain.Find(c, email)
	fmt.Println(janrain)
	

}


func (product *Product) Find(c *cli.Context) {
	term := c.Args().First()
	if len(term) < 3 {
		log.Fatal("Your title must contain at least 3 characters.")
	}
    // Build the URL for the title search
    query := url.Values{
		"simpleSearch":    {term},
	}
	qry := "http://api.oreilly.com/product-service/v1/products?" + query.Encode()
	
	fmt.Println("Working...")
	resp, err := http.Get(qry)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &product)
	
	// Now print the results
	fmt.Printf("|%3s|  %-30s \t| %-12s | %-s\n", " ", pad("Title",30), "Edition", "Oracle Id")
	fmt.Printf("%s\n",strings.Repeat("-",75))
	for i,p := range product.Data {
		fmt.Printf("|%-3d|  %-30s \t| %12s | %s\n", i+1, pad(p.Title,30), p.EditionLabel, p.OracleId)
	}

}