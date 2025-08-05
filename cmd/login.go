// cmd/login.go
package cmd

import (
	"fmt"

	"github.com/emirrcaglar/go-url-shortener/auth"
	"github.com/emirrcaglar/go-url-shortener/db"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to your account",
	Run: func(cmd *cobra.Command, args []string) {
		var username, password string

		fmt.Print("Username: ")
		fmt.Scanln(&username)
		fmt.Print("Password: ")
		fmt.Scanln(&password)

		dbConn, err := db.Connect()
		if err != nil {
			fmt.Printf("Error connecting to database: %v\n", err)
			return
		}
		defer dbConn.Close()

		user, err := auth.Login(dbConn, username, password)
		if err != nil {
			fmt.Printf("❌ Login failed: %v\n", err)
			return
		}

		currentUser = user
		loggedIn = true
		fmt.Printf("✅ Logged in as %s\n", user.UserName)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
