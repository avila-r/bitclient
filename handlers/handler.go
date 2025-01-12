package handlers

import "github.com/spf13/cobra"

type Handler func(*cobra.Command, []string)
