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
	data, err := ioutil.ReadFile("../../data/" + filename)
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

func main() {
	// penghitung waktu
	start := time.Now()

	fmt.Printf("\n==============================================================\n")
	fmt.Println("====================== INSERTION SORT ========================")
	fmt.Printf("==============================================================\n")

	// my data
	data2 := ReadData("Number.txt")

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
