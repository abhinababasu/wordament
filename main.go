package main

import (
	"fmt"
	"os"
)

const Size = 4 // Width and height of the Wordament matrix

func parseMatrix(s string) [][]string {
	matrix := make([][]string, Size)
	for i := range matrix {
		matrix[i] = make([]string, Size)
	}

	fmt.Println(s)

	k := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			ch := s[k]
			k++
			matrix[i][j] = string(ch)
		}
	}

	return matrix
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: wordament <full sequennce of %v letters>\n", Size*Size)
		fmt.Printf("       wordament ZRFLPFUALINXAYEM\n", Size*Size)
		return
	}
	input := os.Args[1:][0]
	if len(input) < Size*Size {
		fmt.Printf("Usage: Input string has to be atleast %v long", Size*Size)
		return
	}

	matrix := parseMatrix(input)

	fmt.Println(matrix)
}
