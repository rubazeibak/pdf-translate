package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var keyAuthToken = "auth-token"

// rootCmd root level flags and commands
var rootCmd = &cobra.Command{
	Use:   "fb-event-locator",
	Short: "Facebook Event Locator",
	Long: `Facebook Event Locator is a tool to locate events according
   to distance from a given coordinate`,
}

//main adds all child commands to the root command sets flags appropriately.
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Cobra supports Persistent Flags which if defined here will be global for the application
	rootCmd.PersistentFlags().StringP(keyAuthToken, "t", "", "fb auth access token")
	viper.BindPFlag("token", rootCmd.Flags().Lookup("port"))
}

// Read in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_")) // from flag to Env format
	viper.AutomaticEnv()                                   // read in environment variables that match
}
