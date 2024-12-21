package main
import "fmt"

func main() {
	const maxData = 1000000
	var data [maxData]int
	var count int

	fmt.Println("Masukkan angka (akhiri dengan -5313):")

	for {
		var input int
		fmt.Scan(&input)

		if input == -5313 {
			break
		}

		if input == 0 {
			insertionSort(data[:count])

			median := findMedian(data[:count])
			fmt.Println("Median saat ini:", median)
		} else {
			data[count] = input
			count++
		}
	}
}

func findMedian(arr []int) float64 {
	n := len(arr)
	if n%2 == 1 {
		return float64(arr[n/2])
	} else {
		mid1, mid2 := arr[(n/2)-1], arr[n/2]
		return float64(mid1+mid2) / 2
	}
}

func insertionSort(arr []int) {
	n := lenArray(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func lenArray(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			break
		}
		count++
	}
	return count
}