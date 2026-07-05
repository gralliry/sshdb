package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gralliry/sshdb/db"
	"github.com/gralliry/sshdb/util"

	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"

	"github.com/spf13/cobra"
)

func importFunc(cmd *cobra.Command, args []string) {
	name := args[0]
	inputDir, _ := cmd.Flags().GetString("input")
	if inputDir == "" {
		inputDir = "."
	}
	privPath, _ := cmd.Flags().GetString("priv")
	pubPath, _ := cmd.Flags().GetString("pub")

	if privPath == "" {
		privPath = filepath.Join(inputDir, name)
	}
	if pubPath == "" {
		pubPath = filepath.Join(inputDir, name) + ".pub"
	}

	privBytes, err := os.ReadFile(privPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read private key: %v\n", err)
		return
	}
	if _, err := ssh.ParsePrivateKey(privBytes); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid private key: %v\n", err)
		return
	}

	pubBytes, err := os.ReadFile(pubPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read public key: %v\n", err)
		return
	}
	keyType, comment, fingerprint, err := util.ParsePublicKey(pubBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid public key: %v\n", err)
		return
	}

	if err := db.Conn().Create(&db.Key{
		Name: name, Type: keyType, Comment: comment,
		Fingerprint: fingerprint, PrivateKey: privBytes, PublicKey: pubBytes,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "UNIQUE") {
			fmt.Fprintf(os.Stderr, "Name %q already exists\n", name)
			return
		}
		fmt.Fprintf(os.Stderr, "Write to database: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stderr, "Imported: %s (%s)\n", name, fingerprint)
}

var importCmd = &cobra.Command{
	Use:     "import <name>",
	Aliases: []string{"i"},
	Short:   "Import key pair",
	Long: `Import an existing SSH key pair into the database.

By default it looks for <name> and <name>.pub in the current directory.
Use --priv / --pub to specify custom paths.`,
	Args: cobra.ExactArgs(1),
	Run: importFunc,
}

func init() {
	importCmd.Flags().StringP("input", "i", ".", "Input directory")
	importCmd.Flags().String("priv", "", "Private key path (default <dir>/<name>)")
	importCmd.Flags().String("pub", "", "Public key path (default <dir>/<name>.pub)")
	rootCmd.AddCommand(importCmd)
}
