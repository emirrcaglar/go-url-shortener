package main

import (
	"database/sql"
	"fmt"
	"log"
	_url "net/url"

	"github.com/emirrcaglar/go-url-shortener/auth"
	"github.com/emirrcaglar/go-url-shortener/db"
	"github.com/emirrcaglar/go-url-shortener/types"
	"github.com/emirrcaglar/go-url-shortener/url"
)

func isValidUrl(inputUrl string) bool {
	_, err := _url.ParseRequestURI(inputUrl)
	return err == nil
}

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Println("error connecting to db")
		return
	}
	defer db.Close()

	var loggedIn bool
	var currentUser *types.User

	for {
		var inputUrl string
		var input string

		fmt.Printf("DEBUG - LOGGED IN: %v\n", loggedIn)

		if loggedIn {
			fmt.Println("1 - URL Shortener")
			fmt.Println("4 - Logout")
			fmt.Println("q - Exit")
			fmt.Scan(&input)

			if input == "q" {
				fmt.Println("goodbye.")
				return
			}

			if input == "1" {
				fmt.Println("enter the url you want to shorten")
				fmt.Scan(&inputUrl)
				shorten(db, inputUrl, currentUser.ID)
				continue
			}

			if input == "4" {
				loggedIn = false
				currentUser = nil
				fmt.Println("Logged out.")
				continue
			}
		}

		if !loggedIn {
			fmt.Println("2 - Login")
			fmt.Println("3 - Register")
			fmt.Println("q - Exit")
			fmt.Scan(&input)

			if input == "q" {
				fmt.Println("goodbye.")
				return
			}

			if input == "2" || input == "3" {
				var userName string
				var password string
				fmt.Println("Enter username")
				fmt.Scan(&userName)
				fmt.Println("Enter password")
				fmt.Scan(&password)

				if input == "2" {
					u, err := auth.Login(db, userName, password)
					if err != nil {
						log.Printf("error logging in: %v", err)
						continue
					}

					currentUser = u
					loggedIn = true

					fmt.Println("Successfully logged in.")
					fmt.Println("Welcome, ", u.UserName)
					continue
				}

				if input == "3" {
					err := auth.Register(db, userName, password)
					if err != nil {
						log.Printf("error registering: %v", err)
						continue
					}
					fmt.Println("Successfully registered.")
					continue
				}
			}
		}
	}
}

func shorten(db *sql.DB, input string, userId int) {

	if !isValidUrl(input) {
		fmt.Printf("Please enter a valid url.\n")
		return
	}

	url := &url.Url{}
	baseUrl := "short.ly/"

	shortUrl := url.GenerateShortUrl(db, url, input, baseUrl, userId)

	fmt.Printf("Your short url is: %s\n", shortUrl)
}
