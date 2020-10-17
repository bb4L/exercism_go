package erratum

// Use wrapper to use a given opener with the input
func Use(opener ResourceOpener, data string) (err error) {
	var value Resource

	value, err = opener()

	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		value, err = opener()
	}

	defer value.Close()
	defer func() {
		if e := recover(); e != nil {
			fErr, ok := e.(FrobError)
			if ok {
				value.Defrob(fErr.defrobTag)
				err = fErr
			}
			nErr, okE := e.(error)
			if okE {
				err = nErr
			}

			if !ok && !okE {

			}
		}
	}()

	value.Frob(data)

	return err
}
