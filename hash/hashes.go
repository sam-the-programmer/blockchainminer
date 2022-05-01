package hash

import (
	md5 "crypto/md5"
	s1 "crypto/sha1"
	s256 "crypto/sha256"
	s512 "crypto/sha512"
	"fmt"
)

// SHA1 returns the SHA1 hash of the inputted string as a string
func SHA1(s string) string {
	return fmt.Sprintf("%x", s1.Sum([]byte(s)))
}

// MD5 returns the MD5 hash of the inputted string as a string
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// SHA224 returns the SHA224 hash of the inputted string as a string
func SHA224(s string) string {
	return fmt.Sprintf("%x", s256.Sum224([]byte(s)))
}

// SHA256 returns the SHA256 hash of the inputted string as a string
func SHA256(s string) string {
	return fmt.Sprintf("%x", s256.Sum256([]byte(s)))
}

// SHA384 returns the SHA384 hash of the inputted string as a string
func SHA384(s string) string {
	return fmt.Sprintf("%x", s512.Sum384([]byte(s)))
}

// SHA512 returns the SHA512 hash of the inputted string as a string
func SHA512(s string) string {
	return fmt.Sprintf("%x", s512.Sum512([]byte(s)))
}
