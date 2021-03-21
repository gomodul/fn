package fn

import (
	"log"
)

// Check godoc.
func Check(fn func() error) {
	if err := fn(); err != nil {
		log.Println(err.Error())
	}
}
