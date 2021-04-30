package fn

// Exception godoc.
type Exception interface{}

// Throw godoc.
func Throw(e ...Exception) {
	if len(e) == 0 {
		panic("something wrong")
	}

	if len(e) > 0 && e[0] != nil {
		panic(e[0])
	}
}

// Block godoc.
type Block struct {
	Try     func()
	Catch   func(e Exception)
	Finally func()
}

// Do godoc.
func (b Block) Do() {
	if b.Finally != nil {
		defer b.Finally()
	}

	if b.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				b.Catch(r)
			}
		}()
	}

	b.Try()
}
