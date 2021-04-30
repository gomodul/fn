package fn

// Exception godoc.
type Exception interface{}

// Throw godoc.
func Throw(e Exception) {
	if e != nil {
		panic(e)
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
