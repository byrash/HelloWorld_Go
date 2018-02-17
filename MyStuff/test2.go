package shiv

//Wrapper to incrementer
func Wrapper() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
