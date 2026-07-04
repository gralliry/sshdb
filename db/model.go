package db

import "time"

// Key is the GORM model — key data is stored entirely in the database.
type Key struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"uniqueIndex;not null"`
	Type        string    `gorm:"not null;default:ssh-ed25519"`
	Comment     string    `gorm:"not null;default:''"`
	Fingerprint string    `gorm:"index"`
	PrivateKey  []byte    `gorm:"not null"`
	PublicKey   []byte    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time
}

func (Key) TableName() string { return "keys" }

// SSHKeyInfo is the display-ready struct for listing.
type SSHKeyInfo struct {
	Name, Type, Comment, Fingerprint, CreatedAt string
}
