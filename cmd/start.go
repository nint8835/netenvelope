package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nint8835/netenvelope/pkg/server"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the web server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		serverInst := server.New()

		bindAddr, _ := cmd.Flags().GetString("bind-addr")

		err := serverInst.Start(bindAddr)
		checkError(err, "Error starting server")
	},
}

func init() {
	startCmd.Flags().String("bind-addr", ":8080", "The address to bind to")

	rootCmd.AddCommand(startCmd)
}
