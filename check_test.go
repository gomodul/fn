package fn_test

import (
	"errors"
	"testing"

	"github.com/gomodul/fn"
)

func TestCheck(t *testing.T) {
	fn.Check(func() error {
		return errors.New("error")
	})
}
