package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	s = s[1:4] // in phan tu 1 den 3
	printSlice(s)

	s = s[:4] // mo rong chieu dai
	printSlice(s)

	s = s[1:] // cat di 1 gia tri
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap =%d %v\n", len(s), cap(s), s)

}
