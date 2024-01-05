package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/nint8835/netenvelope/pkg/database"
)

var dbDropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Completely empty the current database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := database.NewMigrateInstance(dbPath)
		checkError(err, "Error setting up migrations")

		log.Info().Msg("Dropping database...")
		err = m.Drop()
		checkError(err, "Error dropping database")

		log.Info().Msg("Database dropped!")
	},
}

func init() {
	dbCmd.AddCommand(dbDropCmd)
}
