package pork

import (
	"log"

	"github.com/spf13/cobra"
)

// ForkCmd is the fork command definition
var ForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Fork a Github repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatal("You must supply a repository")
		}
		if err := ForkRepository(args[0]); err != nil {
			log.Fatalln("Unable to fork repository:", err)
		}
	},
}

// ForkRepository creates a copy of a repository
func ForkRepository(repository string) error {
	return nil
}
