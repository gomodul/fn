package fn_test

import (
	"testing"

	"github.com/gomodul/fn"
)

func TestBlock_Do(t *testing.T) {
	called := true

	fn.Block{
		Try: func() {
			fn.Throw(nil)
			fn.Throw("errors")
			called = false
		},
		Catch: func(e fn.Exception) {
			if !called {
				t.Fatal("something wrong")
			}
		},
		Finally: func() {
			if !called {
				t.Fatal("something wrong")
			}
		},
	}.Do()
}
