package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/emirrcaglar/go-url-shortener/db"
	"github.com/emirrcaglar/go-url-shortener/utils"
)

func isValidUrl(inputUrl string) bool {
	_, err := url.ParseRequestURI(inputUrl)
	return err == nil
}

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Println("error connecting to db")
		return
	}

	for {
		var inputUrl string

		fmt.Println("Enter the url you want to shorten.")

		fmt.Scan(&inputUrl)
		if inputUrl == "q" {
			fmt.Println("goodbye.")
			return
		}

		if !isValidUrl(inputUrl) {
			fmt.Printf("Please enter a valid url.\n")
			continue
		}

		url := &utils.Url{}
		baseUrl := "short.ly/"

		shortUrl := utils.GenerateShortUrl(db, url, inputUrl, baseUrl)

		fmt.Printf("Your short url is: %s\n", shortUrl)
	}
}
