package cmd

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/nint8835/netenvelope/pkg/database"
)

var dbVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Fetch the current database version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := database.NewMigrateInstance(dbPath)
		if err != nil {
			log.Fatal().Err(err).Msg("Error setting up migrations")
		}

		version, dirty, err := m.Version()
		if err != nil {
			log.Fatal().Err(err).Msg("Error getting database version")
		}

		log.Info().Bool("dirty", dirty).Msg(fmt.Sprintf("Version %d", version))
	},
}

func init() {
	dbCmd.AddCommand(dbVersionCmd)
}
