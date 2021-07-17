package main

import "fmt"

func main() {
	name := [4]string{
		"Nguyen", //0
		"Tung",   //1
		"Lam",    //2
		"Dz",     //3
	}
	fmt.Println(name)
	a := name[0:3] // tru phan tu 3 la dz
	b := name[0:4]
	fmt.Println(a, b)

	b[2] = "X"        // phan tu b tu 1 -> 4 thi co 3 phan tu la tung lam dz
	fmt.Println(a, b) //b[2] thay the = X thay the cho Lam
	fmt.Println(name)
}
