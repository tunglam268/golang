package main

import "fmt"

func main() {
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	p := []bool{true, true, false, false, true}
	fmt.Println(p)

	s := []struct {
		x int
		y bool
	}{
		{2, true}, {3, true}, {1, false}, {4, false}, {5, true},
	}
	fmt.Println(s)
}
