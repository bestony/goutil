package goutil

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstError(t *testing.T) {
	should := assert.New(t)

	const someError = ConstError("some error")
	f := func() error {
		return fmt.Errorf("in func: %w", someError)
	}
	err := f()

	should.Equal("in func: some error", err.Error())
	should.True(errors.Is(err, someError))
}
