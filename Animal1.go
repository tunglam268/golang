package main

import (
	"fmt"
	"os"
)

type Animal struct {
	food, locomation, noise string
}

func (ani Animal) Eat() {
	fmt.Println(ani.food)
}
func (ani Animal) Move() {
	fmt.Println(ani.locomation)
}
func (ani Animal) Speak() {
	fmt.Println(ani.noise)
}

func main() {
	cow := Animal{"co", "di", "woo"}
	snake := Animal{"thit", "truon", "shhh"}
	bird := Animal{"sau", "bay", "peep"}

	var name, info string
	var ani Animal
	for {
		fmt.Println("Nhap 2 bien ten va thong tin cua dong vat; Nhap \"exit\" de thoat ra ngoai")
		fmt.Scanln(&name, &info)
		if name == "exit" {
			os.Exit(0)

		}
		switch name {
		case "cow":
			ani = cow
		case "snake":
			ani = snake
		case "bird":
			ani = bird
		default:
			fmt.Println("chua nhap du lieu")
			continue
		}
		switch info {
		case "eat":
			ani.Eat()
		case "move":
			ani.Move()
		case "speak":
			ani.Speak()
		default:
			fmt.Println("chua nhap du lieu")
			continue

		}
	}
}
