// Package cmd defines the CLI commands for sshdb using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/gralliry/sshdb/db"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sshdb",
	Short: "SSH key management",
	Long: `sshdb — SSH key lifecycle management

Create, list, delete, import, and export SSH key pairs.
All key metadata is stored in ~/.ssh/sshgen.db (SQLite).`,
	SilenceErrors: true,
	SilenceUsage:  true,
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
