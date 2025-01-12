package handlers

import "github.com/spf13/cobra"

type Handler func(*cobra.Command, []string)

func subcommand(cmd *cobra.Command, name string) *cobra.Command {
	for _, children := range cmd.Commands() {
		if children.Use == name {
			return children
		}
	}
	return nil
}
