package fixtures_test

import (
	"testing"

	"github.com/justtrackio/gosoline/pkg/fixtures"
	"github.com/justtrackio/gosoline/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAutoHashed_DifferentSeed(t *testing.T) {
	autoHashed1 := fixtures.NewAutoHashed("test")
	autoHashed2 := fixtures.NewAutoHashed("something else")

	for i := 0; i < 100; i++ {
		assert.NotEqual(t, autoHashed1.GetNextHash(), autoHashed2.GetNextHash())
		assert.NotEqual(t, autoHashed1.GetNextHashString(), autoHashed2.GetNextHashString())
		assert.NotEqual(t, autoHashed1.GetNextUuidV4(), autoHashed2.GetNextUuidV4())
		assert.NotEqual(t, autoHashed1.NewV4(), autoHashed2.NewV4())
	}
}

func TestAutoHashed_SameSeed(t *testing.T) {
	autoHashed1 := fixtures.NewAutoHashed("test")
	autoHashed2 := fixtures.NewAutoHashed("test")

	for i := 0; i < 100; i++ {
		assert.Equal(t, autoHashed1.GetNextHash(), autoHashed2.GetNextHash())
		assert.Equal(t, autoHashed1.GetNextHashString(), autoHashed2.GetNextHashString())
		assert.Equal(t, autoHashed1.GetNextUuidV4(), autoHashed2.GetNextUuidV4())
		assert.Equal(t, autoHashed1.NewV4(), autoHashed2.NewV4())
	}
}

func TestAutoHashed_ValidResults(t *testing.T) {
	autoHashed := fixtures.NewAutoHashed("test")

	for i := 0; i < 100; i++ {
		assert.Len(t, autoHashed.GetNextHash(), 32)
		assert.Len(t, autoHashed.GetNextHashString(), 64)
		assert.True(t, uuid.ValidV4(autoHashed.GetNextUuidV4()))
		assert.True(t, uuid.ValidV4(autoHashed.NewV4()))
	}
}
