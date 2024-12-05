package hash

import (
	"github.com/mono83/hasher"
	"strings"
)

func String[T any](hash hasher.Hasher[T], s string) (T, error) {
	return hash(strings.NewReader(s))
}
