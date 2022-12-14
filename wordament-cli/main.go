package main

import (
	"fmt"
	"os"
	"time"

	"bonggeek.com/wordament/solver"
)

const Size = 4 // Width and height of the Wordament matrix

func main() {

	size := 4
	if len(os.Args) < 2 {
		fmt.Printf("Usage: worament-cli <full sequennce of %v letters>\n", Size*Size)
		fmt.Printf("       worament-cli ZRFLPFUALINXAYEM\n")
		return
	}
	input := os.Args[1:][0]
	if len(input) < size*size {
		fmt.Printf("Usage: Input string has to be atleast %v long", Size*Size)
		return
	}

	tStart := time.Now()

	solver.LoadDictionary("../service/english0.dict")
	solver.LoadDictionary("../service/english2.dict")
	tDict := time.Since(tStart)

	tStart = time.Now()
	w := solver.NewWordament(size)
	solution, err := w.Solve(input)
	if err != nil {
		fmt.Printf("Error!! %v\n ", err)
		return
	}

	tEnd1 := time.Since(tStart)

	// Print the matrix
	fmt.Println("Input:")
	for _, ros := range solution.Input {
		for _, cos := range ros {
			fmt.Print(string(cos), " ")
		}
		fmt.Println()
	}
	fmt.Println()

	longestWordLen := 0
	for _, wordCells := range solution.Result {
		if len(wordCells) > longestWordLen {
			longestWordLen = len(wordCells)
		}

		s := w.WordFromCells(wordCells)
		fmt.Println(s, wordCells)
	}

	fmt.Println()
	fmt.Printf("Total %v words found\n", len(solution.Result))
	fmt.Printf("Longest word size is %v \n", longestWordLen)
	fmt.Printf("Took %vms to load dictionaries, %vms to compute\n", tDict.Milliseconds(), tEnd1.Milliseconds())

}
