package main

import (
	"fmt"
	"os"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
	name string
}
type Snake struct {
	name string
}
type Bird struct {
	name string
}

func (Cow) Eat() {
	fmt.Println("Grass")
}
func (Snake) Eat() {
	fmt.Println("Meats")
}
func (Bird) Eat() {
	fmt.Println("Worms")
}
func (Cow) Move() {
	fmt.Println("Walk")
}
func (Snake) Move() {
	fmt.Println("Slither")
}
func (Bird) Move() {
	fmt.Println("Fly")
}
func (Cow) Speak() {
	fmt.Println("Wooo")
}
func (Snake) Speak() {
	fmt.Println("Shhhhh")

}
func (Bird) Speak() {
	fmt.Println("hssss")
}
func main() {
	var animalMap map[string]Animal
	animalMap = make(map[string]Animal)
	for {
		var commandType, animalType, animalName, animalInfo string
		var newAnimal Animal
		fmt.Println("Enter your command; Enter 'quit' to quit")
		fmt.Scan(&commandType)
		switch commandType {
		case "newanimal":
			fmt.Println("Nhap ten")
			fmt.Scan(&animalName)
			fmt.Println("Nhap loai")
			fmt.Scan(&animalType)

			switch animalType {
			case "cow":
				newAnimal = Cow{animalName}
			case "bird":
				newAnimal = Bird{animalName}
			case "snake":
				newAnimal = Snake{animalName}
			default:
				fmt.Println("wrong animal type , please try again with cow/snake/bird")
				continue

			}
			if _, ok := animalMap[animalName]; ok {
				fmt.Println("Sorry , an animal with the same namr already exists, try again with a different name")
				continue
			} else {
				animalMap[animalName] = newAnimal
				fmt.Println("Animal creation successful")
			}
		case "query":
			fmt.Println("nhap ten loai")
			fmt.Scan(&animalName)
			fmt.Println("nhap thong tin ")
			fmt.Scan(&animalInfo)

			if tempAnimal, ok := animalMap[animalName]; ok {
				switch animalInfo {
				case "eat":
					tempAnimal.Eat()
				case "move":
					tempAnimal.Move()
				case "speak":
					tempAnimal.Speak()
				default:
					fmt.Println("Wrong animal info queried , please try again with eat/move/speak")
					continue

				}
			} else {
				fmt.Println("Animal with given name does not exist.Please try again")
				continue
			}
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Wrong command, please try again with newanimal/query/quit")
		}
	}
}
