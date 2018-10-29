package cmd

import (
	"github.com/BeenVerifiedInc/dockerbox/registry"
	"github.com/BeenVerifiedInc/dockerbox/repo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update the repo from the registry configs",
	RunE: func(cmd *cobra.Command, args []string) error {
		reg, err := registry.New()
		if err != nil {
			return err
		}

		r := repo.New()

		r.Update(reg)
		if err != nil {
			return err
		}

		return nil
	},
}
