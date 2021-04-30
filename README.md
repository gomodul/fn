# fn

### Try-Catch-Finally

```go
package main

import "github.com/gomodul/fn"

func main() {
	fn.Block{
		Try: func() {
			// try to run function...

			// throw exception where not null...
			fn.Throw()
		},
		Catch: func(e fn.Exception) {
			// catch exception from throw function...
		},
		Finally: func() {
			// optional...
		},
	}.Do()
}
```