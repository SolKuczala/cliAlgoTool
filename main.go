package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// create CLI
	// Define command-line flags
	inputFile := flag.String("input", "", "Input CSV file")
	outputFile := flag.String("output", "", "Output CSV file")
	flag.Parse()

	// TODO, make a proper cli, that doesn't return when encountering errors
	// Check if input file is provided
	if *inputFile == "" {
		fmt.Println("Please provide input CSV file using -input flag")
		return
	}
	// Set the default output filename if not provided
	if *outputFile == "" {
		inputFileName := strings.TrimSuffix(filepath.Base(*inputFile), filepath.Ext(*inputFile))
		*outputFile = inputFileName + "_output.csv"
	}

	// Read input CSV file
	openInputFile, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer openInputFile.Close()

	// Parse the input CSV file
	reader := csv.NewReader(openInputFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// call module that processes CSV
	customalgo.Order(*records) // should i return it or use a pointer?

	// Check if output file is provided
	if *outputFile == "" {

		fmt.Println("Please provide output CSV file using -output flag")
		return
	}
	// Create the output CSV file
	newCSV, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer newCSV.Close()

	// Write the processed records to the output CSV file
	writer := csv.NewWriter(newCSV)
	// TODO: format it correctly before writing
	err = writer.WriteAll(orderedRecords)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
