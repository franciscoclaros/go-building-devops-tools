package pork

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SearchCmd is the search command definition
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for Github respositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		repositoryList := SearchByKeyword(args)
		for _, repository := range repositoryList {
			fmt.Println(repository)
		}
	},
}

// SearchByKeyword search repositories by keyword
func SearchByKeyword(keywords []string) []string {
	return []string{"myrepository"}
}
