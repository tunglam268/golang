package main

import "fmt"

var c, python, java bool
var a, b int = 6, 8

const world = "the gioi" // constant = chua

func main() {
	var i int = 5
	var c, python, java = true, true, "Yes" // gan bien
	fmt.Println(a, b, i, c, python, java)

	fmt.Println("Xin chao", world)
}
