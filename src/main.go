package main

import "fmt"

// =================== SELECTION SORT =======================

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

// =================== END SELECTION SORT =====================

// =================== INSERTION SORT =========================

func InsertionSort(arr []int, i, j int) []int {
	if i < j {
		k := i
		InsertionSort(arr, k+1, j)
		Merge(arr, i, k, j)
	}

	return arr
}

func Merge(arr []int, i, k, j int) []int {
	// merge hasil pembagian
	for i <= k && k <= j {
		if arr[i] > arr[k] {
			arr[i], arr[k] = arr[k], arr[i]
			i++
			k++
		} else {
			k++
		}
	}
	return arr
}

// =================== END INSERTION SORT =====================

func main() {

	fmt.Printf("\n==============================================================\n")
	fmt.Println("====================== SELECTION SORT ========================")
	fmt.Printf("==============================================================\n")

	// inisiasi array
	arr := []int{5, 2, 4, 6, 1, 3, 90, 98}

	// print data sebelum diurutkan
	fmt.Printf("Data sebelum diurutkan: %v\n", arr)

	// urutkan data
	arr = SelectionSort(arr, 0, len(arr)-1)

	// print data setelah diurutkan
	fmt.Printf("Data setelah diurutkan: %v\n\n\n\n", arr)

	fmt.Printf("\n==============================================================\n")
	fmt.Println("====================== INSERTION SORT ========================")
	fmt.Printf("==============================================================\n")

	// inisiasi array
	arr2 := []int{5, 2, 4, 6, 1, 3, 90, 98}

	// print data sebelum diurutkan
	fmt.Printf("Data sebelum diurutkan: %v\n", arr2)

	// urutkan data
	res := InsertionSort(arr2, 0, len(arr2)-1)

	// print data setelah diurutkan
	fmt.Printf("Data setelah diurutkan: %v\n\n\n\n", res)

}
