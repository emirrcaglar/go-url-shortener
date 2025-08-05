// cmd/register.go
package cmd

import (
	"fmt"

	"github.com/emirrcaglar/go-url-shortener/auth"
	"github.com/emirrcaglar/go-url-shortener/db"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Create a new account",
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

		err = auth.Register(dbConn, username, password)
		if err != nil {
			fmt.Printf("❌ Registration failed: %v\n", err)
			return
		}

		fmt.Println("✅ Account created! You can now log in.")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
