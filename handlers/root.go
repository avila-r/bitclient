package handlers

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/assets"
	"github.com/avila-r/bitclient/config"
	"github.com/avila-r/bitclient/logger"
)

// Root is the handler for the root command of the CLI application. It presents an interactive form to the user,
// allowing them to choose one of the available options and execute the corresponding command.
var Root = func(cmd *cobra.Command, args []string) {
	// Declare a variable to hold the user-selected command
	var command string

	// Define a map of available commands for the root menu
	commands := map[string]*cobra.Command{
		"config": subcommand(cmd, "config"),
	}

	// Check if the '--debug' flag is set and update the configuration accordingly
	if flag, err := cmd.PersistentFlags().GetBool("debug"); err == nil {
		config.Get().Advanced.Debug = flag
	}

	// Create a form for the user to select an option from a list
	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().
			Title("Choose an option").
			Options(
				huh.NewOption("See current configuration", "config"),
			).
			Validate(func(s string) error {
				// Validate the user's choice against the available commands
				for cmd := range commands {
					if s == cmd {
						return nil
					}
				}

				// If the selected option is invalid, return an error
				return fmt.Errorf("invalid command: %s (valid commands are: %v)", s, commands)
			}).
			Value(&command),
	)).WithTheme(assets.FormTheme)

	// Run the form and handle errors if any
	if err := form.Run(); err != nil {
		logger.Error(err.Error())
	} else {
		// If a valid command is selected, execute it
		if command, exists := commands[command]; exists {
			command.Run(cmd, args)
		}
	}
}
