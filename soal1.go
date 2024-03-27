package main

import "fmt"

func main() {
	// Data uji 1
	lumbaLumbaScores1 := []float64{96, 108, 89}
	koalaScores1 := []float64{88, 91, 110}

	// Data bonus 1
	lumbaLumbaScoresBonus1 := []float64{97, 112, 101}
	koalaScoresBonus1 := []float64{109, 95, 123}

	// Data bonus 2
	lumbaLumbaScoresBonus2 := []float64{97, 112, 101}
	koalaScoresBonus2 := []float64{109, 95, 106}

	// Menghitung skor rata-rata
	lumbaLumbaAvg1 := calculateAverage(lumbaLumbaScores1)
	koalaAvg1 := calculateAverage(koalaScores1)

	lumbaLumbaAvgBonus1 := calculateAverage(lumbaLumbaScoresBonus1)
	koalaAvgBonus1 := calculateAverage(koalaScoresBonus1)

	lumbaLumbaAvgBonus2 := calculateAverage(lumbaLumbaScoresBonus2)
	koalaAvgBonus2 := calculateAverage(koalaScoresBonus2)

	// Memeriksa hasil pertandingan untuk data uji 1
	fmt.Println("Data 1:")
	printWinner(lumbaLumbaAvg1, koalaAvg1)

	// Memeriksa hasil pertandingan untuk bonus 1
	fmt.Println("\nBonus 1:")
	printWinnerWithMinimumScore(lumbaLumbaAvgBonus1, koalaAvgBonus1, 100)

	// Memeriksa hasil pertandingan untuk bonus 2
	fmt.Println("\nBonus 2:")
	printWinnerWithMinimumScore(lumbaLumbaAvgBonus2, koalaAvgBonus2, 100)
}

// Fungsi untuk menghitung skor rata-rata
func calculateAverage(scores []float64) float64 {
	total := 0.0
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))
}

// Fungsi untuk mencetak pemenang
func printWinner(lumbaLumbaAvg, koalaAvg float64) {
	fmt.Printf("Skor rata-rata Lumba-lumba: %.2f\n", lumbaLumbaAvg)
	fmt.Printf("Skor rata-rata Koala: %.2f\n", koalaAvg)

	if lumbaLumbaAvg > koalaAvg {
		fmt.Println("Lumba-lumba memenangkan pertandingan!")
	} else if lumbaLumbaAvg < koalaAvg {
		fmt.Println("Koala memenangkan pertandingan!")
	} else {
		fmt.Println("Hasil seri!")
	}
}

// Fungsi untuk mencetak pemenang dengan skor minimum
func printWinnerWithMinimumScore(lumbaLumbaAvg, koalaAvg, minScore float64) {
	fmt.Printf("Skor rata-rata Lumba-lumba: %.2f\n", lumbaLumbaAvg)
	fmt.Printf("Skor rata-rata Koala: %.2f\n", koalaAvg)

	if lumbaLumbaAvg >= minScore && koalaAvg >= minScore {
		if lumbaLumbaAvg > koalaAvg {
			fmt.Println("Lumba-lumba memenangkan pertandingan!")
		} else if lumbaLumbaAvg < koalaAvg {
			fmt.Println("Koala memenangkan pertandingan!")
		} else {
			fmt.Println("Hasil seri!")
		}
	} else {
		fmt.Println("Tidak ada pemenang karena salah satu tim tidak mencapai skor minimum!")
	}
}
