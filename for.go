package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"

	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //math.Pow luy thua x^n
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) //\n xuong dong
	}
	return lim
}

func main() {
	fmt.Println(sqrt(81), sqrt(-3))
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		fmt.Println("can phan tu i bao gom:", i)
	}
	fmt.Println("Done")
	fmt.Println(pow(2, 3, 8), pow(3, 2, 4))
}
