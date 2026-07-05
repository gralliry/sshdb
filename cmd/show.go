package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func showFunc(_ *cobra.Command, args []string) {
	name := args[0]
	rec, err := Get(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	fmt.Printf("Private key:\n%s\n", string(rec.PrivateKey))
	fmt.Printf("Public key:\n%s\n", string(rec.PublicKey))
}

var showCmd = &cobra.Command{
	Use:     "show <name>",
	Aliases: []string{"s"},
	Short:   "Show key contents",
	Long:    `Print the private and public key data for a given key.`,
	Args:    cobra.ExactArgs(1),
	Run:    showFunc,
}

func init() {
	rootCmd.AddCommand(showCmd)
}
