package erratum

func Use(o ResourceOpener, input string) (e error) {
	result, err := o()
	if err != nil {
		switch err.(type) {
		case TransientError:
			return Use(o, input)
		default:
			return err
		}
	}
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				e = r.(FrobError).inner
				result.Defrob(r.(FrobError).defrobTag)
			case error:
				e = r.(error)
			default:
				panic(r)
			}
		}
		result.Close()
	}()
	result.Frob(input)
	return nil
}
