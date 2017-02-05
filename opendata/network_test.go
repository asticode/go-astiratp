package astiopendata_test

import (
	"testing"

	"github.com/asticode/go-astiratp/opendata"
	"github.com/stretchr/testify/assert"
)

func TestNetworkCode_UnmarshalText(t *testing.T) {
	var c astiopendata.NetworkCode
	err := c.UnmarshalText([]byte("rer"))
	assert.NoError(t, err)
	assert.Equal(t, astiopendata.NetworkCodeRER, c)
	err = c.UnmarshalText([]byte("invalid"))
	assert.Error(t, err)
}
