package cmd

import (
	"fmt"
	"os"

	"github.com/gralliry/sshdb/db"

	"github.com/spf13/cobra"
)

func renameFunc(_ *cobra.Command, args []string) error {
	oldName, newName := args[0], args[1]
	if ok, err := Exists(oldName); err != nil {
		return err
	} else if !ok {
		fmt.Fprintf(os.Stderr, "Error: key %q not found\n", oldName)
		return nil
	}
	if ok, err := Exists(newName); err != nil {
		return err
	} else if ok {
		fmt.Fprintf(os.Stderr, "Error: key %q already exists\n", newName)
		return nil
	}

	result := db.Conn().Model(&db.Key{}).Where("name = ?", oldName).Update("name", newName)
	if result.Error != nil {
		return fmt.Errorf("rename in database: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		fmt.Fprintf(os.Stderr, "Error: key %q not found\n", oldName)
		return nil
	}
	fmt.Fprintf(os.Stderr, "Renamed %q → %q\n", oldName, newName)
	return nil
}

var renameCmd = &cobra.Command{
	Use:     "rename <old-name> <new-name>",
	Aliases: []string{"rn"},
	Short:   "Rename key",
	Long:    `Change the name of a key in the database.`,
	Args:    cobra.ExactArgs(2),
	RunE:    renameFunc,
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
