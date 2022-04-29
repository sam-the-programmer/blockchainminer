package hash

import (
	sha "crypto/sha256"
	"fmt"
)

// SHA256 returns the SHA256 of as string as a string
func SHA256(s string) string {
	return fmt.Sprintf("%x", sha.Sum256([]byte(s)))
}
