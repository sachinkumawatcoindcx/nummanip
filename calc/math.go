package calc

func Add(numbers ...int) int { // variadic function
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}