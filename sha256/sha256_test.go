package sha256

import (
	"encoding/hex"
	"github.com/mono83/hasher/hash"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sha256Data = []struct {
	Expected string
	Given    string
}{
	{Expected: "19fa5b7e29e08b97bf399542642d76e4d11432982a4257943664b86fbf2da8c2", Given: "flüggåәnkб€čhiœßølįên"},
	{Expected: "4ae7c3b6ac0beff671efa8cf57386151c06e58ca53a78d83f36107316cec125f", Given: "Hello, world"},
	{Expected: "1cbec737f863e4922cee63cc2ebbfaafcd1cff8b790d8cfd2e6a5d550b648afa", Given: "Foo"},
	{Expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", Given: ""},
}

func TestCRC32(t *testing.T) {
	for _, datum := range sha256Data {
		t.Run(datum.Given, func(t *testing.T) {
			if v, err := hash.String(SHA256, datum.Given); assert.NoError(t, err) {
				assert.Equal(t, datum.Expected, hex.EncodeToString(v))
			}
		})
	}
}
