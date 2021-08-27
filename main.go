package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dghubble/oauth1"
	"github.com/keaising/go-twitter/twitter"
)

func main() {
	// Twitter client
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TOKEN"), os.Getenv("TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	days := int(time.Now().In(loc).Sub(time.Unix(1629003600, 0)).Hours() / 24)

	var name = fmt.Sprintf("%s%v", os.Getenv("PREFIX"), days)

	u, _, err := client.Accounts.UpdateProfile(&twitter.AccountUpdateProfileParams{
		Name: &name,
	})
	if err != nil {
		log.Println("update profile failed", err)
		return
	}
	log.Println(u.Name)
	log.Println(u.Description)
	log.Println(u.Location)
	log.Println(u.Status.FullText)
}
