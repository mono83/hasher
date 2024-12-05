package hasher

import "io"

// Hasher defines functions that can be used to produce hash
// This functions works with [io.Reader] for maximum versatility,
// for simpler cases like working with strings or bytes utility
// package [hash] can be used.
//
// Example:
// hash.String(sha256.SHA256, "Hello, world")
type Hasher[T any] func(io.Reader) (T, error)
