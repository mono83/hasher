package hash

import (
	"bytes"
	"github.com/mono83/hasher"
)

func Bytes[T any](hash hasher.Hasher[T], bts []byte) (T, error) {
	return hash(bytes.NewReader(bts))
}
