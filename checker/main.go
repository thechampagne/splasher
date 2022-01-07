package main

import (
	"context"
	"fmt"
	"github.com/hbagdi/go-unsplash/unsplash"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func main() {
	args := os.Args
	checker(args)
}

func checker(args []string) {
	if len(args) < 2 {
		log.Fatal("Error: Missing Token")
	} else {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: args[1]},
		)
		client := oauth2.NewClient(context.Background(), ts)

		splash := unsplash.New(client)
		opt :=  unsplash.RandomPhotoOpt{}
		_, _, err := splash.Photos.Random(&opt)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error: %s", err))
		}
		file, err := os.OpenFile("TOKEN", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error: %s", err))
		}
		stat, err := file.Stat()
		if err != nil {
			log.Fatal(fmt.Sprintf("Error: %s", err))
		}
		if stat.Size() != 0 {
			err := os.Truncate("TOKEN", 0)
			if err != nil {
				log.Fatal(fmt.Sprintf("Error: %s", err))
			}
		}
		_, er := file.WriteString(args[1])
		if er != nil {
			log.Fatal(fmt.Sprintf("Error: %s", er))
		}
		fmt.Print("Token saved Successfully.")
	}
}