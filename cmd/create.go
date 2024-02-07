package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Manually create new resources in the database",
}

func init() {
	rootCmd.AddCommand(createCmd)
}
