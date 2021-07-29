package kata

import "strings"

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

// Write a function called repeatStr which repeats the given string string exactly n times.

// repeatStr(6, "I") // "IIIIII"
// repeatStr(5, "Hello") // "HelloHelloHelloHelloHello"
