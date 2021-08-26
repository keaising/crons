package main

import (
	"log"

	"github.com/dghubble/oauth1"
	"github.com/keaising/go-twitter/twitter"
)

func main() {
	config := oauth1.NewConfig("", "")
	token := oauth1.NewToken("", "")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	var name = "Taiga"
	var url = "www.bing.com"
	u, _, err := client.Accounts.UpdateProfile(&twitter.AccountUpdateProfileParams{
		Name:             &name,
		URL:              &url,
		Location:         &name,
		ProfileLinkColor: &name,
		Description:      &name,
		IncludeEntities:  twitter.Bool(false),
		SkipStatus:       twitter.Bool(true),
	})
	if err != nil {
		log.Println("update profile failed", err)
		return
	}
	log.Println(u.Name, u.Description, u.Location, u.Status)
}
