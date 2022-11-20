package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"bonggeek.com/wordament/solver"
)

const (
	WordamentSize = 4 // number of rows or colms of the mordament
)

var wordament solver.Wordament

func main() {
	port := flag.Int("port", 8080, "Port server will listen on")
	flag.Parse()

	log.Println("Creating wordament solver")
	wordament = *solver.NewWordament(WordamentSize)

	log.Println("Loading dictionaries")

	// load from current or the solver paths
	found := loadAllDictionaries(&wordament, "../solver/*.dict")
	found = found + loadAllDictionaries(&wordament, "./*.dict")
	if found == 0 {
		log.Fatal("No dictionaries were found")
	}

	http.HandleFunc("/", handler) // each request calls handler

	addr := fmt.Sprintf(":%v", *port)
	log.Println("Listening on ", addr)

	log.Fatal(http.ListenAndServe(addr, nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	log.Printf("Remote=%v;Request=%v %v,UA=%v", r.RemoteAddr, r.Method, r.URL.String(), r.UserAgent())
	query := r.URL.Query()
	input := query.Get("input")

	if len(input) != WordamentSize*WordamentSize {
		errorStr := "Error! Input is too long or too short"
		log.Println(errorStr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorStr))
		return
	}

	result, err := wordament.Solve(input)
	if err != nil {
		log.Println("Solve error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//j, _ := json.MarshalIndent(result, "", "  ") // use this for pretty printing on the client directly
	j, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

	timeEnd := time.Since(timeStart)
	log.Printf("Solved=%v;time=%vms;words=%v", input, timeEnd.Milliseconds(), len(result.Result))
}

func loadAllDictionaries(wordament *solver.Wordament, path string) int {
	// load from current or the solver paths
	found := 0
	log.Println("Loading dictionary from", path)

	files, err := filepath.Glob(path)
	if err == nil {
		for _, file := range files {
			log.Println("Loading dictionary", file)
			err := wordament.LoadDictionary(file)
			if err != nil {
				log.Fatal(err)
			}
			found++
		}
	}

	return found
}
