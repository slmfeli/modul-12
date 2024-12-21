package main
import "fmt"

func main() {
	const maxDaerah = 100
	const maxRumah = 100

	var n int
	fmt.Print("Masukkan jumlah daerah kerabat Hercules: ")
	fmt.Scan(&n)

	var jumlahRumah [maxDaerah]int
	var masukan [maxDaerah][maxRumah]int

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan jumlah rumah di daerah %d: ", i+1)
		fmt.Scan(&jumlahRumah[i])

		fmt.Printf("Masukkan nomor rumah di daerah %d (pisahkan dengan spasi): ", i+1)
		for j := 0; j < jumlahRumah[i]; j++ {
			fmt.Scan(&masukan[i][j])
		}

		selectionSort(masukan[i][:jumlahRumah[i]])
	}

	fmt.Println("\nKeluarannya:")
	fmt.Printf("%-5s %-15s %s\n", "No", "Masukan", "Keluaran")
	for i := 0; i < n; i++ {
		fmt.Printf("%-5d %-15s %s\n", i+1, formatMasukan(masukan[i][:jumlahRumah[i]]), formatMasukan(masukan[i][:jumlahRumah[i]]))
	}
}

func selectionSort(arr []int) {
	for i := 0; i < lenArray(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < lenArray(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func formatMasukan(angka []int) string {
	var hasil string
	for i := 0; i < lenArray(angka); i++ {
		if i > 0 {
			hasil += " "
		}
		hasil += fmt.Sprint(angka[i])
	}
	return hasil
}

func lenArray(arr []int) int {
	count := 0
	for range arr {
		count++
	}
	return count
}