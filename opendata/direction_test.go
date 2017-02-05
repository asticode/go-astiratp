package astiopendata_test

import (
	"testing"

	"github.com/asticode/go-astiratp/opendata"
	"github.com/stretchr/testify/assert"
)

func TestDirectionWay_UnmarshalText(t *testing.T) {
	var w astiopendata.DirectionWay
	err := w.UnmarshalText([]byte("A"))
	assert.NoError(t, err)
	assert.Equal(t, astiopendata.DirectionWayOneWay, w)
	err = w.UnmarshalText([]byte("invalid"))
	assert.Error(t, err)
}
