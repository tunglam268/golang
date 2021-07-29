package kata

func MakeNegative(x int) int {
	switch {
	case x > 0:
		return -x
	case x < 0:
		return x
	default:
		return 0
	}
}

// In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?

// Example:

// MakeNegative(1)    # return -1
// MakeNegative(-5)   # return -5
// MakeNegative(0)    # return 0
