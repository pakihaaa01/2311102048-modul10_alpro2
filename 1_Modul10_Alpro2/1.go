package main

import (
	"fmt"
)

func main() {
	var N int
	var weights [1000]float64

	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&N)

	fmt.Println("Masukkan berat kelinci:")
	for i := 0; i < N; i++ {
		fmt.Scan(&weights[i])
	}

	minWeight := weights[0]
	maxWeight := weights[0]

	for i := 1; i < N; i++ {
		if weights[i] < minWeight {
			minWeight = weights[i]
		}
		if weights[i] > maxWeight {
			maxWeight = weights[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat terbesar: %.2f\n", maxWeight)
}
