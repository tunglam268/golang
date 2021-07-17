package main

import "fmt"

func main() {
	var weight int = 65
	if weight < 50 {
		fmt.Println("Ban thieu can nang")
	} else if weight >= 55 && weight <= 70 {
		fmt.Println("Can nang du chi tieu")
	} else if weight >= 71 && weight <= 80 {
		fmt.Println("Ban bi thua can")
	} else {
		fmt.Println("ban bi thua can")
	}

}
