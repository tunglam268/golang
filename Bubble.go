package main

import "fmt"

func Swap(arr []int, index int) {
	arr[index], arr[index+1] = arr[index+1], arr[index]
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j = j + 1 {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func main() {
	var size, num int
	var arr []int
	fmt.Println("Nhap mang ( tu 1 den 10 ): ")
	fmt.Scan(&size)

	if size < 1 || size > 10 {
		panic("Size nen nhap tu 1 den 10") // giong try catch
	}
	fmt.Println("nhap thanh phan cua mang")

	for i := 0; i < size; i++ {
		fmt.Scan(&num) //nhap tu ban phim num cac phan tu trong mang
		arr = append(arr, num)
	}
	BubbleSort(arr)
	fmt.Println("Mang sau khi duoc sap xep", arr)
}
