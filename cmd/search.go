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

	"github.com/drpaneas/igdb-go/igdbclient"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search based on name, results are sorted by similarity to the given search string",
	Long: `Search based on name, results are sorted by similarity to the given search string
Examples:
	- igdb-go seach -n Super Mario    # Searches only for "Super" and ignores "Mario"
	- igdb-go search -n "Super Mario" # Searches for the exact string
`,
	Run: func(cmd *cobra.Command, args []string) {
		cl, err := igdbclient.NewClient(AccessToken, ClientID)
		name, _ := cmd.Flags().GetString("name")
		gameSearch, err := cl.SearchGame(name)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range *gameSearch {
			fmt.Printf("%v\n", v.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	searchCmd.Flags().StringP("name", "n", "", "Name of the game")
}
