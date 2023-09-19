package strings

func internalFunc() {
	// In Go, a name is exported (visible outside the package) if it begins with a capital letter.
}

func CountOddEven(s string) (odds, evens int) {
	odds, evens = 0, 0
	for _, c := range s {
		if int(c)%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return
}
