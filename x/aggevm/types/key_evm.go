package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EvmKeyPrefix is the prefix to retrieve all Evm
	EvmKeyPrefix = "Evm/value/"
)

// EvmKey returns the store key to retrieve a Evm from the index fields
func EvmKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
