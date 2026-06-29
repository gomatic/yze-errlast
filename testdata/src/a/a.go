package a

// good returns error last.
func good() (int, error) { return 0, nil }

// bad returns error not last.
func bad() (error, int) { return nil, 0 } // want `last return`

// noErr returns no error.
func noErr() int { return 0 }

// noResults returns nothing.
func noResults() {}

// named returns named results with error last.
func named() (n int, err error) { return 0, nil }

// unnamedBad returns unnamed results with error not last.
func unnamedBad() (error, string) { return nil, "" } // want `last return`
