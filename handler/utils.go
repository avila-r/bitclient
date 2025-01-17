package handler

import "github.com/spf13/cobra"

// Handler defines a function type that handles commands with a cobra.Command
type Handler func(*cobra.Command, []string)

// subcommand searches for a subcommand by its name within the parent command's list of subcommands.
// It returns the matching subcommand or nil if no match is found.
func subcommand(cmd *cobra.Command, name string) *cobra.Command {
	for _, children := range cmd.Commands() {
		if children.Use == name || children.Name() == name {
			return children
		}
	}
	// Return nil if no subcommand with the name is found
	return nil
}
