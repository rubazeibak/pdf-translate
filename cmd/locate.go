package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var keyLongitude = "longtitude"
var keyLatitude = "latitude"

// locateCmd respresents the add command
var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "locate events",
	Long:  `locate events according to location and filers`,
	Run:   locateRun,
}

func locateRun(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(locateCmd)

	locateCmd.PersistentFlags().Float32(keyLongitude, 0, "Location longtitude")
	locateCmd.PersistentFlags().Float32(keyLatitude, 0, "Location longtitude")
	viper.BindPFlags(locateCmd.PersistentFlags())
}
