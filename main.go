package main

import (
	"fmt"
	"os"
)

const Size = 4 // Width and height of the Wordament matrix

func parseMatrix(s string) [][]rune {
	matrix := make([][]rune, Size)
	for i := range matrix {
		matrix[i] = make([]rune, Size)
	}

	// we are just assuming these to be english chars
	k := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			ch := s[k]
			k++
			matrix[i][j] = rune(ch)
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

	// get the input string in a Size x Size 2D slice. We use string and not char (rune) because later
	// enhancement should also cover multi char per cell wordaments
	matrix := parseMatrix(input)

	// Print the matrix
	fmt.Println("Input:")
	for _, ros := range matrix {
		for _, cos := range ros {
			fmt.Print(string(cos), " ")
		}
		fmt.Println()
	}
	fmt.Println()

	w := NewWordament(Size)
	w.LoadDictionary("english.0") // todo: add other dicts as well

	solution := w.Solve(matrix)
	fmt.Println(solution)
}
