package cmd

import (
	"errors"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/nint8835/netenvelope/pkg/database"
)

var dbStepCmd = &cobra.Command{
	Use: "step [number of steps]",
	Example: `  Downgrade by one version: netenvelope db step -- -1
  Upgrade by 3 versions: netenvelope db step 3`,
	Short: "Perform a relative migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := database.NewMigrateInstance(dbPath)
		checkError(err, "Error setting up migrations")

		stepCountStr := args[0]

		stepCount, err := strconv.Atoi(stepCountStr)
		checkError(err, "Error parsing step count")

		log.Info().Msg("Performing database migration...")
		err = m.Steps(stepCount)
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Info().Msg("Database is already at desired version.")
				return
			} else {
				log.Fatal().Err(err).Msg("Error upgrading database")
			}
		}

		log.Info().Msg("Database migrations complete!")
	},
}

func init() {
	dbCmd.AddCommand(dbStepCmd)
}
