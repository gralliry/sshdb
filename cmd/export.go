package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gralliry/sshdb/util"

	"github.com/spf13/cobra"
)

func exportFunc(cmd *cobra.Command, args []string) error {
	name := args[0]
	outputDir, _ := cmd.Flags().GetString("output")
	if outputDir == "" {
		outputDir = "."
	}
	privPath, _ := cmd.Flags().GetString("priv")
	pubPath, _ := cmd.Flags().GetString("pub")

	rec, err := Get(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return nil
	}
	if privPath == "" {
		privPath = filepath.Join(outputDir, name)
	}
	if pubPath == "" {
		pubPath = filepath.Join(outputDir, name) + ".pub"
	}

	existing := []string{}
	if _, err := os.Stat(privPath); err == nil {
		existing = append(existing, privPath)
	}
	if _, err := os.Stat(pubPath); err == nil {
		existing = append(existing, pubPath)
	}
	if len(existing) > 0 {
		fmt.Fprintf(os.Stderr, "File already exists:\n")
		for _, f := range existing {
			fmt.Fprintf(os.Stderr, "  %s\n", f)
		}
		fmt.Fprintf(os.Stderr, "Overwrite? [y/N] ")
		var answer string
		fmt.Scanln(&answer)
		if answer != "y" && answer != "Y" {
			fmt.Fprintln(os.Stderr, "Cancelled")
			return nil
		}
	}

	if err := util.WriteFile(privPath, rec.PrivateKey, 0600); err != nil {
		return fmt.Errorf("write private key: %w", err)
	}
	if err := util.WriteFile(pubPath, rec.PublicKey, 0644); err != nil {
		return fmt.Errorf("write public key: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Exported to:\n  %s\n  %s\n", privPath, pubPath)
	return nil
}

var exportCmd = &cobra.Command{
	Use:     "export <name>",
	Aliases: []string{"e"},
	Short:   "Export key to directory",
	Long: `Write key pair files to disk.

Default output paths: <output>/<name> and <output>/<name>.pub.
Use --priv / --pub to specify custom paths.`,
	Args: cobra.ExactArgs(1),
	RunE: exportFunc,
}

func init() {
	exportCmd.Flags().StringP("output", "o", ".", "Output directory")
	exportCmd.Flags().String("priv", "", "Private key output path (default <dir>/<name>)")
	exportCmd.Flags().String("pub", "", "Public key output path (default <priv-path>.pub)")
	rootCmd.AddCommand(exportCmd)
}
