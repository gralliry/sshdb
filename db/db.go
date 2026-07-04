// Package db manages the SQLite database (via GORM) for SSH key metadata.
package db

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ─── Database ────────────────────────────────────────────────────────────────

var (
	gdb    *gorm.DB
	dbOnce sync.Once
	dbErr  error
)

// Init opens (or creates) the SQLite database and runs migrations.
func Init() error {
	dbOnce.Do(func() { dbErr = initDB() })
	return dbErr
}

func initDB() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("get home dir: %w", err)
	}
	dir := filepath.Join(home, ".ssh")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("create %s: %w", dir, err)
	}

	dbPath := filepath.Join(dir, "sshgen.db") +
		"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)"
	gdb, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	return gdb.AutoMigrate(&Key{})
}

// Conn returns the gorm.DB instance for use by the core package.
func Conn() *gorm.DB { return gdb }
