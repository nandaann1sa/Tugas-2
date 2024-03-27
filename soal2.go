package main

import (
	"fmt"
)

func main() {
	// Data 1
	markWeight1 := 78.0 // kg
	markHeight1 := 1.69 // m

	johnWeight1 := 92.0 // kg
	johnHeight1 := 1.95 // m

	// Menghitung BMI untuk Mark dan John menggunakan kedua versi rumus
	markBMI1 := markWeight1 / (markHeight1 * markHeight1)
	johnBMI1 := johnWeight1 / (johnHeight1 * johnHeight1)

	// Menampilkan hasil BMI
	fmt.Printf("BMI Mark (Data 1): %.2f\n", markBMI1)
	fmt.Printf("BMI John (Data 1): %.2f\n", johnBMI1)

	// Data 2
	markWeight2 := 95.0 // kg
	markHeight2 := 1.88 // m

	johnWeight2 := 85.0 // kg
	johnHeight2 := 1.76 // m

	// Menghitung BMI untuk Mark dan John menggunakan kedua versi rumus
	markBMI2 := markWeight2 / (markHeight2 * markHeight2)
	johnBMI2 := johnWeight2 / (johnHeight2 * johnHeight2)

	// Menampilkan hasil BMI
	fmt.Printf("\nBMI Mark (Data 2): %.2f\n", markBMI2)
	fmt.Printf("BMI John (Data 2): %.2f\n", johnBMI2)

	// Membuat variabel boolean untuk menentukan apakah Mark memiliki BMI lebih tinggi dari John
	markHigherBMI1 := markBMI1 > johnBMI1
	markHigherBMI2 := markBMI2 > johnBMI2

	// Menampilkan informasi apakah Mark memiliki BMI lebih tinggi dari John
	fmt.Printf("\nApakah Mark memiliki BMI lebih tinggi dari John (Data 1) ? %t\n", markHigherBMI1)
	fmt.Printf("Apakah Mark memiliki BMI lebih tinggi dari John (Data 2) ? %t\n", markHigherBMI2)
}
