package cmd

import "github.com/spf13/cobra"

// RootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "taskx",
	Short: "Taskx is a CLI task manager to manage all your TODOs from the comfort of your CLI",
}
