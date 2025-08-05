// cmd/login.go
package cmd

import (
	"fmt"

	"github.com/emirrcaglar/go-url-shortener/auth"
	"github.com/emirrcaglar/go-url-shortener/db"
	"github.com/emirrcaglar/go-url-shortener/session"
	"github.com/emirrcaglar/go-url-shortener/types"
	"github.com/emirrcaglar/go-url-shortener/utils"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to your account",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := utils.CheckStatus()
		if err == nil {
			fmt.Println("Already logged in.")
			fmt.Println("Run: go-url-shortener logout")
			return
		}
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

		// Save session
		cfg = &session.Cfg{
			LoggedIn: true,
			CurrentUser: &types.User{
				ID:       user.ID,
				UserName: user.UserName,
			},
		}
		err = session.SaveConfig(cfg)
		if err != nil {
			fmt.Printf("❌ Could not save session: %v\n", err)
			return
		}
		fmt.Printf("✅ Logged in as %s\n", user.UserName)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
