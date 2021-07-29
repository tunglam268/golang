package kata

import "strings"

func AbbrevName(name string) string {
	words := strings.Split(name, " ")

	return strings.ToUpper(string(words[0][0]) + "." + string(words[1][0]))
}

// Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

// The output should be two capital letters with a dot separating them.

// It should look like this:

// Sam Harris => S.H

// Patrick Feeney => P.F
