package main

import (
	"fmt"
	)

func Transpose(M [][]float64) [][]float64 {
	Mt := make([][]float64, len(M[0]), len(M))
	for i := range(M) {
		for j := range(Mt) {
			Mt[j] = append(Mt[j], M[i][j])
		}
	}
	return Mt
}

func main() {
	M := [][]float64{[]float64{1., 2.}, []float64{3., 4.}, []float64{5., 6.}}

	Mt := Transpose(M)
	fmt.Println(M)
	fmt.Println(Mt)
}
