package main

import (
	"fmt"
	"time"
)

// function insertion sort dengan parameter array, awal, akhir dan balikan berupa slice int
func InsertionSort(arr []int, i, j int) []int {
	// jika awal lebih kecil dari akhir
	if i < j {
		// tetapkan nilai awal
		k := i
		// rekursif
		Merge(arr, i, k, i)
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

func main() {
	// penghitung waktu
	start := time.Now()

	// my data
	data2 := []int{5, 2, 4, 6, 1, 3}

	fmt.Printf("\n==============================================================\n")
	fmt.Println("====================== INSERTION SORT ========================")
	fmt.Printf("==============================================================\n")

	// print data sebelum diurutkan
	fmt.Printf("Data sebelum diurutkan: %v\n", data2)

	// urutkan data
	res := InsertionSort(data2, 0, len(data2)-1)

	// print data setelah diurutkan
	fmt.Printf("Data setelah diurutkan: %v\n\n\n\n", res)

	// durasi ekseskusi program
	fmt.Printf("Durasi eksekusi program: %v Detik \n", time.Since(start).Seconds())

	var space string
	fmt.Printf("Tekan enter untuk keluar")
	fmt.Scanln(&space)
}
