package cmd

import (
	"fmt"
	"strings"

	"github.com/gralliry/sshgen/db"

	"github.com/spf13/cobra"
)

func listFunc(_ *cobra.Command, _ []string) error {
	var records []db.Key
	if err := db.Conn().Order("created_at DESC").Find(&records).Error; err != nil {
		return fmt.Errorf("query db: %w", err)
	}
	if len(records) == 0 {
		fmt.Println("(no keys)")
		return nil
	}

	const fmtStr = "%-14s %-14s  %-14s  %-10s  %s"
	header := fmt.Sprintf(fmtStr, "Name", "Type", "Comment", "Created", "Fingerprint")
	fmt.Println(header)
	fmt.Println(strings.Repeat("─", len(header)))
	for _, r := range records {
		created := ""
		if !r.CreatedAt.IsZero() {
			created = r.CreatedAt.Local().Format("2006-01-02")
		}
		fmt.Printf(fmtStr+"\n", r.Name, r.Type, r.Comment, created, r.Fingerprint)
	}
	return nil
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "List all keys",
	Long:    `List all SSH keys in the database.`,
	Args:    cobra.NoArgs,
	RunE:    listFunc,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
