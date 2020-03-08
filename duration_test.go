package goutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDuration(t *testing.T) {
	tests := []struct {
		d    time.Duration
		want string
	}{
		{time.Second, "1s"},
		{5 * time.Minute, "5m0s"},
	}

	should := require.New(t)
	must := assert.New(t)

	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			duration := Duration(test.d)

			got, err := duration.MarshalText()
			must.Nil(err)
			should.Equal(test.want, string(got))

			err = duration.UnmarshalText(got)
			must.Nil(err)
			should.Equal(test.d, duration.Duration())
		})
	}
}
