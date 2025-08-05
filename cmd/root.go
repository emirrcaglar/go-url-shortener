// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/emirrcaglar/go-url-shortener/types"
	"github.com/spf13/cobra"
)

var loggedIn bool
var currentUser *types.User // Make sure types.User is accessible

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-url-shortener",
	Short: "A simple URL shortener CLI",
	Long: `A CLI tool to shorten URLs and manage them with login support.
You need to log in to create short links.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Go URL Shortener!")
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
