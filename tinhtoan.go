package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int { // x int , y int = x,y int
	return x + y

}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum*3 - 20
	y = (x + sum) / 3
	fmt.Println("X va Y ", x, y)
	return x, y

}

func main() {
	fmt.Println("so ua thich la", rand.Intn(10))
	fmt.Println("Số lỗi của bài là", math.Sqrt(9))
	fmt.Println("So pi :", math.Pi)
	fmt.Println("x+y", add(30, 60))
	a, b := swap("halo", "hello")
	fmt.Println(a, b)
	fmt.Println(split(30))

}
