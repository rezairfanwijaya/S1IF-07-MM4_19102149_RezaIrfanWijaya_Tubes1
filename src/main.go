// Kelas	: MM4
// Nama		: REZA IRFAN WIJAYA
// NIM		: 19102149

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// ambil data dari folder data
func ReadData(filename string) []int {
	// ambil path dari file
	data, err := ioutil.ReadFile("../data/" + filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	// ubah data byte menjadi string
	byteString := strings.Split(string(data), " ")
	res := []int{}
	for _, v := range byteString {
		// ubah string menjadi int
		i, _ := strconv.Atoi(v)
		res = append(res, i)
	}

	return res

}

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

	// penghitung waktu
	start := time.Now()

	// my data
	data := ReadData("Number.txt")

	fmt.Printf("\n==============================================================\n")
	fmt.Println("====================== SELECTION SORT ========================")
	fmt.Printf("==============================================================\n")

	// print data sebelum diurutkan
	fmt.Printf("Data sebelum diurutkan: %v\n", data)

	// urutkan data
	data = SelectionSort(data, 0, len(data)-1)

	// print data setelah diurutkan
	fmt.Printf("Data setelah diurutkan: %v\n\n\n\n", data)

	// durasi ekseskusi program
	fmt.Printf("Durasi eksekusi program selection-sort: %v Detik \n", time.Since(start).Seconds())

	fmt.Printf("\n==============================================================\n")
	fmt.Println("====================== INSERTION SORT ========================")
	fmt.Printf("==============================================================\n")

	// referesh time
	start = time.Now()
	
	// my data
	data2 := ReadData("Number.txt")

	// print data sebelum diurutkan
	fmt.Printf("Data sebelum diurutkan: %v\n", data2)

	// urutkan data
	res := InsertionSort(data2, 0, len(data2)-1)

	// print data setelah diurutkan
	fmt.Printf("Data setelah diurutkan: %v\n\n\n\n", res)

	// durasi ekseskusi program
	fmt.Printf("Durasi eksekusi program insertion-sort: %v Detik \n", time.Since(start).Seconds())

	var space string
	fmt.Printf("Tekan enter untuk keluar")
	fmt.Scanln(&space)
}
