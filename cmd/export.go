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
		return err
	}
	if privPath == "" {
		privPath = filepath.Join(outputDir, name)
	}
	if pubPath == "" {
		pubPath = filepath.Join(outputDir, name) + ".pub"
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
