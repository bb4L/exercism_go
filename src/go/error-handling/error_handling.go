package erratum

import (
	"fmt"
)

// Use wrapper to use a given opener with the input
func Use(opener ResourceOpener, data string) (err error) {
	var value Resource

	for value, err = opener(); err != nil; value, err = opener() {
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}

	defer value.Close()
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)

			fErr, ok := e.(FrobError)
			if ok {
				value.Defrob(fErr.defrobTag)
				err = fErr
			}
			nErr, okE := e.(error)
			if okE {
				err = nErr
			}
		}
	}()

	value.Frob(data)

	return err
}
