package kata

import "fmt"

func MultiTable(number int) string {
	result := ""
	for i := 1; i <= 10; i++ {
		result += fmt.Sprintf("%d * %d = %d", i, number, number*i)
		if i != 10 {
			result += "\n"
		}
	}
	return result
}

// Your goal is to return multiplication table for number that is always an integer from 1 to 10.

// For example, a multiplication table (string) for number == 5 looks like below:
