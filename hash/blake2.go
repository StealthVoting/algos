package hash

import (
	"github.com/ethereum/go-ethereum/crypto/blake2b"
)

func Hash(data []byte) []byte {
	hash := blake2b.Sum256(data)

	return hash[:]
}
