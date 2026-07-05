package cmd

import (
	"fmt"
	"os"

	"github.com/gralliry/sshdb/db"

	"github.com/spf13/cobra"
)

func renameFunc(_ *cobra.Command, args []string) {
	oldName, newName := args[0], args[1]
	if ok, err := Exists(oldName); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	} else if !ok {
		fmt.Fprintf(os.Stderr, "Name %q not found\n", oldName)
		return
	}
	if ok, err := Exists(newName); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	} else if ok {
		fmt.Fprintf(os.Stderr, "Name %q already exists\n", newName)
		return
	}

	result := db.Conn().Model(&db.Key{}).Where("name = ?", oldName).Update("name", newName)
	if result.Error != nil {
		fmt.Fprintf(os.Stderr, "Rename in database: %v\n", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		fmt.Fprintf(os.Stderr, "Name %q not found\n", oldName)
		return
	}
	fmt.Fprintf(os.Stderr, "Renamed %q → %q\n", oldName, newName)
}

var renameCmd = &cobra.Command{
	Use:     "rename <old-name> <new-name>",
	Aliases: []string{"rn"},
	Short:   "Rename key",
	Long:    `Change the name of a key in the database.`,
	Args:    cobra.ExactArgs(2),
	Run:    renameFunc,
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
