package crc32

import (
	"github.com/mono83/hasher/hash"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ieeeData = []struct {
	Expected uint32
	Given    string
}{
	{Expected: 3850467951, Given: "flüggåәnkб€čhiœßølįên"},
	{Expected: 3885672898, Given: "Hello, world"},
	{Expected: 3023971265, Given: "Foo"},
	{Expected: 0, Given: ""},
}

func TestCRC32(t *testing.T) {
	for _, datum := range ieeeData {
		t.Run(datum.Given, func(t *testing.T) {
			if v, err := hash.String(CRC32, datum.Given); assert.NoError(t, err) {
				assert.Equal(t, datum.Expected, v)
			}
		})
	}
}
