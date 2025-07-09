package main

import (
	"fmt"
	"net/url"

	"github.com/emirrcaglar/go-url-shortener/utils"
)

func isValidUrl(inputUrl string) bool {
	_, err := url.ParseRequestURI(inputUrl)
	return err == nil
}

func main() {
	var inputUrl string

	fmt.Println("Enter the url you want to shorten.")
	fmt.Scan(&inputUrl)

	if !isValidUrl(inputUrl) {
		fmt.Printf("Please enter a valid url.")
		return
	}

	url := &utils.Url{}
	baseUrl := "short.ly/"

	shortUrl := utils.GenerateShortUrl(url, inputUrl, baseUrl)

	fmt.Printf("Your short url is: %s\n", shortUrl)
}
