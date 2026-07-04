// Package cmd defines the CLI commands for sshgen using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/gralliry/sshgen/db"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sshgen",
	Short: "SSH key management tool",
	Long: `sshgen — SSH key lifecycle management

Create, list, delete, import, and export SSH key pairs.
All key metadata is stored in ~/.ssh/sshgen.db (SQLite).`,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return db.Init()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// Execute is the entry point for the CLI, called from main.go.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
