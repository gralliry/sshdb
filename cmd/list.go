package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/gralliry/sshdb/db"

	"github.com/spf13/cobra"
)

func listFunc(_ *cobra.Command, _ []string) {
	var records []db.Key
	if err := db.Conn().Order("created_at DESC").Find(&records).Error; err != nil {
		fmt.Fprintf(os.Stderr, "query database: %v\n", err)
		return
	}
	if len(records) == 0 {
		fmt.Println("(no keys)")
		return
	}

	const fmtStr = "%-14s %-14s  %-14s  %-19s  %s"
	header := fmt.Sprintf(fmtStr, "Name", "Type", "Comment", "Created", "Fingerprint")
	fmt.Println(header)
	fmt.Println(strings.Repeat("─", len(header)))
	for _, r := range records {
		created := ""
		if !r.CreatedAt.IsZero() {
			created = r.CreatedAt.Local().Format("2006-01-02 15:04:05")
		}
		fmt.Printf(fmtStr+"\n", r.Name, r.Type, r.Comment, created, r.Fingerprint)
	}
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "List all keys",
	Long:    `List all SSH keys in the database.`,
	Args:    cobra.NoArgs,
	Run:    listFunc,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
