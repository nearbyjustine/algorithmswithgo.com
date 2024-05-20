package module01

// Sum will sum up all of the numbers passed
// in and return the result
// func Sum(numbers []int) int {
// 	result := 0
// 	for _, val := range numbers {
// 		result += val;
// 	}
// 	return result
// }

func Sum(numbers []int) int {
	result := 0
	if len(numbers) == 0 {
		return 0
	}
	// first, rest := numbers[0], numbers[1:]
	// result = first + Sum(rest)
	// or
	result = numbers[0] + Sum(numbers[1:])
	return result;
}
