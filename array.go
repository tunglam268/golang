package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[1], a[0])
	fmt.Println(a) // in tu 0 den 1

	primes := [6]int{2, 3, 4, 5, 6, 7}
	fmt.Println(primes)
	var s []int = primes[0:6] // low -> high-1
	fmt.Print(s)
}
