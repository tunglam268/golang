package main

import (
	"fmt"
)

func main() {
	i, j := 43, 2701

	p := &i         // tro toi diem i = 43
	fmt.Println(*p) // doc i thong qua con tro
	*p = 21         // set diem i moi thong qua con tro
	fmt.Println(i)  // diem i moi

	p = &j
	*p = *p / 37 // chia j thong qua con tro
	fmt.Println(j)
}
