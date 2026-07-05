package cmd

import (
	"errors"
	"fmt"

	"github.com/gralliry/sshdb/db"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func exactArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != n {
			cmd.Help()
			return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
		}
		return nil
	}
}

func Get(name string) (*db.Key, error) {
	var rec db.Key
	err := db.Conn().Where("name = ?", name).First(&rec).Error
	if err == nil {
		return &rec, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("key %q not found", name)
	}
	return nil, fmt.Errorf("look up key %q: %w", name, err)
}

func Exists(name string) (bool, error) {
	var count int64
	err := db.Conn().Model(&db.Key{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}
