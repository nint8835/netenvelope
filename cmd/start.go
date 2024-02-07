package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/nint8835/netenvelope/pkg/server"
)

var defaultSessionSecret = "default"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the web server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		bindAddr := envVarOrFlagValue("BIND_ADDR", "bind-addr", cmd)
		sessionSecret := envVarOrFlagValue("SESSION_SECRET", "session-secret", cmd)

		if sessionSecret == defaultSessionSecret {
			log.Warn().Msg("Using the default session secret - this is insecure and should not be used in production!")
			log.Warn().Msg("Use --session-secret or the SESSION_SECRET environment variable to set a custom session secret.")
		}

		serverInst := server.New(
			server.Config{
				BindAddr:      bindAddr,
				SessionSecret: sessionSecret,
			},
		)

		err := serverInst.Start()
		checkError(err, "Error starting server")
	},
}

func init() {
	startCmd.Flags().String("bind-addr", ":8080", "The address to bind to")
	startCmd.Flags().String("session-secret", defaultSessionSecret, "The secret to use to sign session cookies")

	rootCmd.AddCommand(startCmd)
}
