package main

import (
	"context"
	"fmt"
	"github.com/hbagdi/go-unsplash/unsplash"
	"github.com/reujab/wallpaper"
	"golang.org/x/oauth2"
	"log"
	"os"
	"tawesoft.co.uk/go/dialog"
	"time"
)

var (
	query        string
	wallpaperMode string
)

func main() {
	args := os.Args
	app(args)
}

func app(args []string) {
	if len(args) > 2 {
		if args[1] == "--query" {
			query = args[2]
		} else if args[1] == "--mode" {
			wallpaperMode = args[2]
		}
	} else if len(args) > 4 {
		if args[1] == "--query" {
			query = args[2]
		} else if args[1] == "--mode" {
			wallpaperMode = args[2]
		}
		if args[3] == "--query" {
			query = args[4]
		} else if args[3] == "--mode" {
			wallpaperMode = args[4]
		}
	}

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		dialog.Alert("%s", err)
		os.Exit(1)
	}
	log.SetOutput(file)
	token, err := os.ReadFile("TOKEN")
	if err != nil {
		log.Println(err)
		dialog.Alert("%s", err)
		os.Exit(1)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: string(token)},
	)
	client := oauth2.NewClient(context.Background(), ts)
	splash := unsplash.New(client)
	opt := unsplash.RandomPhotoOpt{}
	if query != "" {
		opt.SearchQuery = query
	}
	opt.Count = 1
	opt.Orientation = unsplash.Landscape

	for {
		select {
		case <-time.NewTicker(120 * time.Second).C:
			photos, _, err := splash.Photos.Random(&opt)
			if err != nil {
				log.Println(err)
			}
			var url *unsplash.URL
			for _, p := range *photos {
				url = p.Urls.Full
			}
			er := wallpaper.SetFromURL(fmt.Sprintf("%s", url))
			if er != nil {
				log.Println(er)
			}

			if wallpaperMode == "stretch" {
				err := wallpaper.SetMode(wallpaper.Stretch)
				if err != nil {
					log.Println(err)
				}
			} else if wallpaperMode == "center" {
				err := wallpaper.SetMode(wallpaper.Center)
				if err != nil {
					log.Println(err)
				}
			} else if wallpaperMode == "crop" {
				err := wallpaper.SetMode(wallpaper.Crop)
				if err != nil {
					log.Println(err)
				}
			} else if wallpaperMode == "fit" {
				err := wallpaper.SetMode(wallpaper.Fit)
				if err != nil {
					log.Println(err)
				}
			} else if wallpaperMode == "span" {
				err := wallpaper.SetMode(wallpaper.Span)
				if err != nil {
					log.Println(err)
				}
			} else if wallpaperMode == "tile" {
				err := wallpaper.SetMode(wallpaper.Tile)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}