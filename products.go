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



// pads a string to be N chars long
func pad(s string, N int) string {
	if len(s) > N {
		return s[:N]
	} else {
		return s + strings.Repeat(" ", N - len(s))
	}
	
}



func (p *Product) Grant(args *cli.Context) {
	
	if len(args.Args()) != 2 {
		log.Fatal("You must supply an oracle id and a user email")
	}
	
	oracle_id := args.Args().Get(0)
	email := args.Args().Get(1)
    
	fmt.Printf("Granting %s to %s...\n", oracle_id, email)
	
	// Load the credentials that we'll use to set the grantor in the ownership service
	creds := Credentials{}
	creds.Load()
	
	// Find the grantor's Janrain UUID
	grantor := JanrainUser{}
	grantor.Find(creds.JanrainEmail)
	fmt.Println(grantor.Results[0].UUID)
	
	// Find the grantee's Janrain UUID
	grantee := JanrainUser{}
	grantee.Find(email)
	fmt.Println(grantee.Results[0].UUID)

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