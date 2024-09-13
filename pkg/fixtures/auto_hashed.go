package fixtures

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/justtrackio/gosoline/pkg/uuid"
)

type AutoHashed struct {
	nextSeed []byte
}

var _ uuid.Uuid = &AutoHashed{}

// NewAutoHashed provides a way to generate random-looking values for fixtures from a given seed. A possible use case would be to provide the name of
// the fixture set you are creating as the seed to get distinct values from other fixture sets, but having deterministic values compared to just
// generating random UUIDs for example.
func NewAutoHashed(seed string) *AutoHashed {
	return &AutoHashed{
		nextSeed: []byte(seed),
	}
}

// GetNextHash provides a fresh 32-byte hash for a fixture in case you don't want to assign specific values to random data or similar.
// Keep in mind that the ids are unique only in the scope of the same *AutoHashed instance with the given seed.
func (h *AutoHashed) GetNextHash() []byte {
	hash := sha256.Sum256(h.nextSeed)
	h.nextSeed = hash[:]

	return hash[:]
}

// GetNextHashString is the same as GetNextHash, but encodes the hash using base 16.
func (h *AutoHashed) GetNextHashString() string {
	return hex.EncodeToString(h.GetNextHash())
}

// GetNextUuidV4 produces a version 4 UUID from the data returned by GetNextHash.
func (h *AutoHashed) GetNextUuidV4() string {
	hashed := h.GetNextHash()

	uuidBytes := make([]byte, 16)
	for i := range uuidBytes {
		uuidBytes[i] = hashed[i*2] ^ hashed[i*2+1]
	}

	// set version and variant
	uuidBytes[6] = (uuidBytes[6] & 0x0F) | 0x40
	uuidBytes[8] = (uuidBytes[8] & 0x3F) | 0x80

	return fmt.Sprintf(
		"%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		uuidBytes[0],
		uuidBytes[1],
		uuidBytes[2],
		uuidBytes[3],
		uuidBytes[4],
		uuidBytes[5],
		uuidBytes[6],
		uuidBytes[7],
		uuidBytes[8],
		uuidBytes[9],
		uuidBytes[10],
		uuidBytes[11],
		uuidBytes[12],
		uuidBytes[13],
		uuidBytes[14],
		uuidBytes[15],
	)
}

func (h *AutoHashed) NewV4() string {
	return h.GetNextUuidV4()
}
