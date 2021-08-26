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
	u, _, err := client.Accounts.UpdateProfile(&twitter.AccountUpdateProfileParams{
		Name: &name,
	})
	if err != nil {
		log.Println("update profile failed", err)
		return
	}
	log.Println(u.Name, u.Description, u.Location, u.Status)
}
