// Package util provides SSH key parsing and file I/O utilities.
package util

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

// ParsePublicKey extracts key type, comment, and fingerprint from
// authorized_keys-format public key bytes.
// Only the first line is parsed.
func ParsePublicKey(pubBytes []byte) (keyType, comment, fingerprint string, err error) {
	line := string(pubBytes)
	if idx := strings.IndexAny(line, "\n\r"); idx >= 0 {
		line = line[:idx]
	}
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return "", "", "", fmt.Errorf("invalid public key: need key type and base64 data")
	}
	keyType = parts[0]
	if len(parts) >= 3 {
		comment = strings.Join(parts[2:], " ")
	}
	parsed, _, _, _, e := ssh.ParseAuthorizedKey(pubBytes)
	if e != nil {
		return "", "", "", fmt.Errorf("invalid SSH public key: %w", e)
	}
	fingerprint = ssh.FingerprintSHA256(parsed)
	return keyType, comment, fingerprint, nil
}
