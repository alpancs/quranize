package quran

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadQuranBadXMLFormat(t *testing.T) {
	assert.Panics(t, func() { parseQuran("") })
}
