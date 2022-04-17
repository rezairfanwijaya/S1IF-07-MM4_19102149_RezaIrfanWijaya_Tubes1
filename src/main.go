package main

import "fmt"

// =================== SELECTION SORT =======================

// function selection sort dengan parameter array, awal, akhir dan balikan berupa slice int
func SelectionSort(arr []int, i, j int) []int {
	// jika awal lebih kecil dari akhir
	if i < j {
		// lakukan partisi pada array
		p := Partisi3(arr, i, j)
		// urutkan array dengan partisi yang telah di lakukan
		arr = SelectionSort(arr, i, p-1)
		// urutkan array dengan partisi yang telah di lakukan
		arr = SelectionSort(arr, p+1, j)
	}
	// kembalikan array
	return arr
}

// function artisi dengan parameter array, awal, akhir denagan balikan berupa int
func Partisi3(arr []int, i, j int) int {
	// looping semua data pada array
	for k := i + 1; k <= j; k++ {
		// tetap kan nilai minimal
		nilaiMin := i
		// jika array k lebih kecil dari array k+1
		if arr[k] < arr[nilaiMin] {
			// nilai minimal di ganti dengan k
			nilaiMin = k
		}
		// lakukan pertukaran posisi
		arr[nilaiMin], arr[i] = arr[i], arr[nilaiMin]
	}

	// kembalikan nilai minimal
	return i
}

// =================== END SELECTION SORT =====================

// =================== INSERTION SORT =========================

// function insertion sort dengan parameter array, awal, akhir dan balikan berupa slice int
func InsertionSort(arr []int, i, j int) []int {
	// jika awal lebih kecil dari akhir
	if i < j {
		// tetapkan nilai awal
		k := i
		// rekursif
		InsertionSort(arr, k+1, j)
		// lakukan penggabungan
		Merge(arr, i, k, j)
	}

	return arr
}

// function merge dengan parameter array, awal, tengah, akhir dan balikan berupa slice int
func Merge(arr []int, i, k, j int) []int {
	// merge hasil pembagian
	for i <= k && k <= j {
		// jika nilai awal lebih besar dari nilai tengah
		if arr[i] > arr[k] {
			// lakukan pertukaran posisi
			arr[i], arr[k] = arr[k], arr[i]
			// increment i dan k
			i++
			k++
		} else {
			// incerment k
			k++
		}
	}
	// kembalikan array
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
