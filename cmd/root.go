/*
Copyright © 2020 Panagiotis Georgiadis <drpaneas@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var AccessToken string
var ClientID string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "igdb-go",
	Short: "Command-line tool for accessing the IGDB API",
	Long: `A command-line tool for accesing the IGDB API:

Requires Twitch developer account (access token & client id)
     https://api-docs.igdb.com/?shell#account-creation`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Allow the user to choose their own configuration file
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.igdb-go.yaml)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".igdb-go" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".igdb-go")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Read variables from ENV and the configuration file, if they exist
	// Note: ENV variables overwrite the configuration file values
	if viper.GetString("IGDB_TOKEN") != "" {
		AccessToken = viper.GetString("IGDB_TOKEN")
	} else {
		log.Fatalf("[DEBUG]\tNo IGDB_TOKEN environment variable set")
	}

	if viper.GetString("IGDB_CLIENT_ID") != "" {
		ClientID = viper.GetString("IGDB_CLIENT_ID")
	} else {
		log.Fatal("[DEBUG]\tNo IGDB_CLIENT_ID environment variable set")
	}

}
