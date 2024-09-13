package fixtures_test

import (
	"testing"

	"github.com/justtrackio/gosoline/pkg/fixtures"
	"github.com/justtrackio/gosoline/pkg/mdl"
	"github.com/stretchr/testify/assert"
)

func TestAutoNumbered(t *testing.T) {
	autoNumbered := fixtures.NewAutoNumbered()
	assert.Equal(t, mdl.Box[uint](1), autoNumbered.GetNext())
	assert.Equal(t, mdl.Box[uint](2), autoNumbered.GetNext())
	assert.Equal(t, mdl.Box[uint](3), autoNumbered.GetNext())
}

func TestAutoNumberedFrom(t *testing.T) {
	autoNumbered := fixtures.NewAutoNumberedFrom(5)
	assert.Equal(t, mdl.Box[uint](5), autoNumbered.GetNext())
	assert.Equal(t, mdl.Box[uint](6), autoNumbered.GetNext())
	assert.Equal(t, mdl.Box[uint](7), autoNumbered.GetNext())
}
