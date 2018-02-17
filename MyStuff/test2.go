package shiv

func Wrapper() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
