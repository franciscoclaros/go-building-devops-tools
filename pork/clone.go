package pork

import (
	"log"

	"github.com/spf13/cobra"
)

// CloneCmd is the clone command definition
var CloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone repository from Github",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply the repository")
		}
		if err := CloneRepository(args[0], ref, create); err != nil {
			log.Fatalln("Error when cloning repository:", err)
		}
	},
}

// CloneRepository clones the repository
func CloneRepository(repository string, ref string, shouldCreate bool) error {
	return nil
}

var ref string
var create bool

func init() {
	CloneCmd.PersistentFlags().StringVar(&ref, "ref", "", "Specific reference to check out")
	CloneCmd.PersistentFlags().BoolVar(&create, "create", false, "Create the reference if it does not exist")
}
