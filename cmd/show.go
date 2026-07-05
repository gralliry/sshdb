package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func showFunc(_ *cobra.Command, args []string) error {
	name := args[0]
	rec, err := Get(name)
	if err != nil {
		return err
	}

	fmt.Printf("Private key:\n%s\n", string(rec.PrivateKey))
	fmt.Printf("Public key:\n%s\n", string(rec.PublicKey))
	return nil
}

var showCmd = &cobra.Command{
	Use:     "show <name>",
	Aliases: []string{"s"},
	Short:   "Show key contents",
	Long:    `Print the private and public key data for a given key.`,
	Args:    exactArgs(1),
	RunE:    showFunc,
}

func init() {
	rootCmd.AddCommand(showCmd)
}
