package main

import "fmt"

func SelectionSort(arr []int, i, j int) []int {
	if i < j {
		p := Partisi3(arr, i, j)
		arr = SelectionSort(arr, i, p-1)
		arr = SelectionSort(arr, p+1, j)
	}
	return arr
}

func Partisi3(arr []int, i, j int) int {
	for k := i + 1; k <= j; k++ {
		nilaiMin := i
		if arr[k] < arr[nilaiMin] {
			nilaiMin = k
		}
		arr[nilaiMin], arr[i] = arr[i], arr[nilaiMin]
	}

	return i
}

func main() {
	// inisiasi array
	arr := []int{5, 2, 4, 6, 1, 3, 90, 98}

	// print data sebelum diurutkan
	fmt.Printf("Data sebelum diurutkan: %v\n", arr)

	// urutkan data
	arr = SelectionSort(arr, 0, len(arr)-1)

	// print data setelah diurutkan
	fmt.Printf("Data setelah diurutkan: %v\n", arr)

}
