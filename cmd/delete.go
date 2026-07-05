package cmd

import (
	"fmt"
	"os"

	"github.com/gralliry/sshdb/db"

	"github.com/spf13/cobra"
)

func deleteFunc(_ *cobra.Command, args []string) error {
	name := args[0]
	result := db.Conn().Where("name = ?", name).Delete(&db.Key{})
	if result.Error != nil {
		return fmt.Errorf("delete from database: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		fmt.Fprintf(os.Stderr, "Error: key %q not found\n", name)
		return nil
	}
	fmt.Fprintf(os.Stderr, "Key %q deleted\n", name)
	return nil
}

var deleteCmd = &cobra.Command{
	Use:     "delete <name>",
	Aliases: []string{"del", "d", "rm", "remove"},
	Short:   "Delete key",
	Long:    `Remove an SSH key from the database.`,
	Args:    cobra.ExactArgs(1),
	RunE:    deleteFunc,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
