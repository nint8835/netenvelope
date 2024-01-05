package cmd

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/nint8835/netenvelope/pkg/database"
)

var dbUpgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"up"},
	Short:   "Perform all database upgrades",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := database.NewMigrateInstance(dbPath)
		checkError(err, "Error setting up migrations")

		log.Info().Msg("Performing database upgrades...")
		err = m.Up()
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Info().Msg("Database is already up to date!")
				return
			} else {
				log.Fatal().Err(err).Msg("Error upgrading database")
			}
		}

		log.Info().Msg("Database upgrades complete!")
	},
}

func init() {
	dbCmd.AddCommand(dbUpgradeCmd)
}
