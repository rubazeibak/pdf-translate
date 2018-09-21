package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sruba90z/pdf-translate/internal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var keySourceLng = "source-lng"
var keyTargetLng = "target-lng"
var keyPath = "path"

var rootCmd = &cobra.Command{
	Use:   "translate",
	Short: "translate pdf",
	Long:  `translate pdf according to source and target languages`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sl := viper.GetString(keySourceLng)
		tl := viper.GetString(keyTargetLng)
		path := viper.GetString(keyPath)
		p := internal.Processor{
			Reader:     newReader(path),
			Translator: newTranslator(sl, tl),
		}
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String(keySourceLng, "", "Source language")
	rootCmd.PersistentFlags().String(keyTargetLng, "", "Target language")
	rootCmd.PersistentFlags().String(keyPath, "", "Path of pdf file")
	viper.BindPFlags(rootCmd.PersistentFlags())
}

func initConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_")) // from flag to Env format
	viper.AutomaticEnv()                                   // read in environment variables that match
}
