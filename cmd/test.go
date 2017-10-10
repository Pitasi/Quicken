package cmd

import (
	"github.com/matteojoliveau/quicken/recipe"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Print the version number of Quicken",
	Long:  `All software has versions. This is Quicken's`,
	Run: func(cmd *cobra.Command, args []string) {
		recipe.ParseRecipe("test.yml")
	},
}
