package calc

import "errors"

func Add(numbers ...int) (error, int) { // variadic function
	result := 0

	if len(numbers) < 2 {
		return errors.New("provide more than 1 number"), result
	} else {
		for _, number := range numbers {
			result += number
		}
		return nil, result
	}
}