package cmd

import (
	"errors"
	"fmt"

	"github.com/gralliry/sshdb/db"

	"gorm.io/gorm"
)

func Get(name string) (*db.Key, error) {
	var rec db.Key
	err := db.Conn().Where("name = ?", name).First(&rec).Error
	if err == nil {
		return &rec, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("Key %q not found", name)
	}
	return nil, fmt.Errorf("Look up key %q: %w", name, err)
}

func Exists(name string) (bool, error) {
	var count int64
	err := db.Conn().Model(&db.Key{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}
