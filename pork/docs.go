package pork

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// DocsCmd is the docs command definition
var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Read the documentation for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatal("You must supply repository argument")
		}
		content := GetRepositoryReadme(args[0])
		fmt.Println(content)
	},
}

func GetRepositoryReadme(repository string) string {
	return repository
}
