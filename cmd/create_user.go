package cmd

import (
	"context"

	"github.com/charmbracelet/huh"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"

	"github.com/nint8835/netenvelope/pkg/database"
	"github.com/nint8835/netenvelope/pkg/database/queries"
)

var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Create a new user",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		dbInst, err := database.New(dbPath)
		checkError(err, "Error setting up database")

		var username, password string

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title("Username").Value(&username),
				huh.NewInput().Password(true).Title("Password").Value(&password),
			),
		)

		err = form.Run()
		checkError(err, "Error getting user input")

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		checkError(err, "Error hashing password")

		user, err := dbInst.CreateUser(
			context.Background(),
			queries.CreateUserParams{
				Username:     username,
				PasswordHash: hash,
			},
		)
		checkError(err, "Error creating user")

		log.Info().Msgf("Successfully created user %s with ID %d", user.Username, user.ID)
	},
}

func init() {
	createCmd.AddCommand(createUserCmd)
}
