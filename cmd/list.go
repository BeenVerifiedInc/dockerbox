package cmd

import (
	"fmt"

	"github.com/BeenVerifiedInc/dockerbox/repo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all available applets in the repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := repo.New()
		r.Init()

		for _, a := range r.Applets {
			fmt.Println(a.Name)
		}

		return nil
	},
}
