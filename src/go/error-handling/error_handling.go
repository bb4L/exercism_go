package erratum

import "fmt"

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
			switch t := e.(type) {
			case FrobError:
				value.Defrob(t.defrobTag)
				err = t

			case error:
				err = t

			default:
				err = fmt.Errorf("%s", e)
			}
		}
	}()

	value.Frob(data)

	return err
}
