package main

import "golang.org/x/tour/pic"

func imageFun(x, y int) uint8 {
	return uint8(x) ^ uint8(y)
}
func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		row := make([]uint8, dy)
		for j := 0; j < dy; j++ {
			row[j] = imageFun(i, j)
		}
		image[i] = row
	}
	return image
}

func main() {
	pic.Show(Pic)
}
