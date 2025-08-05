/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/emirrcaglar/go-url-shortener/db"
	"github.com/emirrcaglar/go-url-shortener/utils"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show user's url history",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := utils.CheckStatus()
		if err != nil {
			fmt.Printf("❌ %v\n", err)
			fmt.Println("Run: go-url-shortener login")
			return
		}

		dbConn, err := db.Connect()
		if err != nil {
			fmt.Printf("Error connecting to database: %v\n", err)
			return
		}
		defer dbConn.Close()

		rows, err := dbConn.Query("SELECT id, long_url, short_url FROM urls WHERE userID = ?", cfg.CurrentUser.ID)
		if err != nil {
			fmt.Printf("❌ Error querying URLs: %v\n", err)
			return
		}
		defer rows.Close()

		fmt.Printf("Your shortened URLs (Username: %s):\n", cfg.CurrentUser.UserName)

		found := false
		for rows.Next() {
			var id int
			var longURL string
			var shortURL string

			err := rows.Scan(&id, &longURL, &shortURL)
			if err != nil {
				fmt.Printf("❌ Error scanning row: %v\n", err)
				return
			}

			fmt.Printf(" - %s → %s\n", longURL, shortURL)
			found = true
		}

		if !found {
			fmt.Println("You haven't shortened any URLs yet.")
		}

		// Check for errors during iteration
		if err = rows.Err(); err != nil {
			fmt.Printf("❌ Row iteration error: %v\n", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
