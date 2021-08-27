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
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TOKEN"), os.Getenv("TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	prefix := os.Getenv("PREFIX")
	loc, _ := time.LoadLocation("Asia/Shanghai")
	days := int(time.Now().In(loc).Sub(time.Unix(1629003600, 0)).Hours() / 24)
	var name = fmt.Sprintf("%s%v", prefix, days)
	u, _, err := client.Accounts.UpdateProfile(&twitter.AccountUpdateProfileParams{
		Name: &name,
	})
	if err != nil {
		log.Println("update profile failed", err)
		return
	}
	log.Println(u.Name, u.Description, u.Location, u.Status)
}
