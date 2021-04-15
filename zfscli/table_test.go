package zfscli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanTable(t *testing.T) {
	assert := require.New(t)

	expected := [][]string{
		{"NAME", "PROPERTY", "VALUE", "SOURCE"},
		{"foo", "bar", "x  y", "z"},
		{"fizz", "buzz", "a", "bcdefgh"},
		nil,
	}

	err := ScanTable([]byte(""+
		"NAME   PROPERTY   VALUE   SOURCE\n"+
		"foo    bar        x  y    z\n"+
		"fizz   buzz       a       bcdefgh\n",
	), func(i int, row []string) error {
		assert.Equal(expected[i], row)
		return nil
	})
	assert.NoError(err)
}
