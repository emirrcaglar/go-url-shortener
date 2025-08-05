// cmd/shorten.go
package cmd

import (
	"fmt"
	"net/url"

	"github.com/emirrcaglar/go-url-shortener/db"
	"github.com/emirrcaglar/go-url-shortener/urlpkg"
	"github.com/emirrcaglar/go-url-shortener/utils"
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
		cfg, err := utils.CheckStatus()
		if err != nil {
			fmt.Printf("error shortening url: %v", err)
			return
		}

		dbConn, err := db.Connect()
		if err != nil {
			fmt.Printf("❌ Database error: %v\n", err)
			return
		}
		defer db.Close(dbConn)

		url_ := &urlpkg.Url{}
		baseUrl := "short.ly/"

		longURL := args[0]
		custom, _ := cmd.Flags().GetString("custom")

		if custom != "" {
			err := urlpkg.GenerateCustomUrl(dbConn, url_, longURL, custom, baseUrl, cfg.CurrentUser.ID)
			if err != nil {
				fmt.Printf("error generating custom url: %v\n", err)
				return
			}
			fmt.Println("✅ Custom url created.")
			return
		}

		if !isValidURL(longURL) {
			fmt.Println("❌ Invalid URL. Please enter a valid one.")
			return
		}

		shortURL, err := url_.GenerateShortUrl(dbConn, url_, longURL, baseUrl, cfg.CurrentUser.ID)
		if err != nil {
			fmt.Printf("❌ Error creating short URL: %v\n", err)
			return
		}

		fmt.Printf("✅ Short URL: %s\n", shortURL)
	},
}

func init() {
	rootCmd.AddCommand(shortenCmd)

	shortenCmd.Flags().StringP("custom", "c", "", "Custom code")
}
