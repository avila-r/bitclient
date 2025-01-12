package handlers

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/assets"
	"github.com/avila-r/bitclient/config"
	"github.com/avila-r/bitclient/logger"
)

var Root = func(cmd *cobra.Command, args []string) {
	var command string
	commands := map[string]*cobra.Command{
		"config": subcommand(cmd, "config"),
	}

	if flag, err := cmd.PersistentFlags().GetBool("debug"); err == nil {
		config.Get().Advanced.Debug = flag
	}

	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().
			Title("Choose your burger").
			Options(
				huh.NewOption("See current configuration", "config"),
			).
			Validate(func(s string) error {
				for cmd := range commands {
					if s == cmd {
						return nil
					}
				}

				return fmt.Errorf("invalid command: %s (valid commands are: %v)", s, commands)
			}).
			Value(&command),
	)).WithTheme(assets.FormTheme)

	if err := form.Run(); err != nil {
		logger.Error(err.Error())
	} else {
		if command, exists := commands[command]; exists {
			command.Run(cmd, args)
		}
	}
}
