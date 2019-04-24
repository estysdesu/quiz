package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"log"
	"flag"
)

func main() {

	fileLoc := flag.String("csv", "./problems.csv", "the path to the csv problems file")
	flag.Parse()

	// open the file and defer close
	csvFile, err := os.Open(*fileLoc)
	if err != nil {
		log.Println(err)
	}
	defer csvFile.Close()

	// create new csv reader then read the files content
	reader := csv.NewReader(csvFile)
	csvData, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(csvData)
	fmt.Println(len(csvData))

	// make the q&a slices
	q := make([]string, 0, len(csvData))
	a := make([]string, 0, len(csvData))
	for _, data := range csvData {
		q = append(q, data[0])
		a = append(a, data[1])
	}

	
}
