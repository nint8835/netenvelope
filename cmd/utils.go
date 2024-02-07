package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func checkError(err error, message string) {
	if err != nil {
		log.Fatal().Err(err).Msg(message)
	}
}

func envVarOrFlagValue(envVar string, flagName string, cmd *cobra.Command) string {
	envValue, envValueExists := os.LookupEnv(envVar)
	if envValueExists {
		return envValue
	}

	flagValue, _ := cmd.Flags().GetString(flagName)
	return flagValue
}
