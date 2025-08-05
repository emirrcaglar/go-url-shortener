// cmd/shorten.go
package cmd

import (
	"fmt"
	"net/url"

	"github.com/emirrcaglar/go-url-shortener/db"
	"github.com/emirrcaglar/go-url-shortener/urlpkg"
	"github.com/spf13/cobra"
)

func isValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	return err == nil
}

var shortenCmd = &cobra.Command{
	Use:   "shorten [url]",
	Short: "Shorten a long URL",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		longURL := args[0]

		if !isValidURL(longURL) {
			fmt.Println("❌ Invalid URL. Please enter a valid one.")
			return
		}

		if !loggedIn || currentUser == nil {
			fmt.Println("❌ You must be logged in to shorten URLs.")
			fmt.Println("Run: go-url-shortener login")
			return
		}

		dbConn, err := db.Connect()
		if err != nil {
			fmt.Printf("❌ Database error: %v\n", err)
			return
		}
		defer dbConn.Close()

		url_ := &urlpkg.Url{}
		baseUrl := "short.ly/"
		shortURL, err := url_.GenerateShortUrl(dbConn, url_, longURL, baseUrl, currentUser.ID)
		if err != nil {
			fmt.Printf("❌ Error creating short URL: %v\n", err)
			return
		}

		fmt.Printf("✅ Short URL: %s\n", shortURL)
	},
}

func init() {
	rootCmd.AddCommand(shortenCmd)
}
