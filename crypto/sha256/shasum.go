package sha256

import (
	"crypto/sha256"
	"fmt"
)

func shasum(raw []byte) string {
	sum := sha256.Sum256(raw)
	return fmt.Sprintf("%x", sum)
}
