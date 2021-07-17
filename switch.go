package main

import "fmt"
		"time"

func main() {
	var dayOfWeek = 4
	switch dayOfWeek {
	case 1:
		fmt.Println("Thu 2")
	case 2:
		fmt.Println("Thu 3")
	case 3:
		fmt.Println("Thu 4")
	case 4:
		fmt.Println("Thu 5")
	case 5:
		fmt.Println("Thu 6")
	default:
		fmt.Println("Khong tồn tại ngày")

	}
}
