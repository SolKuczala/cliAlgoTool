package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	// TODO: fix broken import
	"clialgotool/customAlgo"
)

func getColumnIndex(header []string, columnName string) (int, error) {
	for i, col := range header {
		if col == columnName {
			return i, nil
		}
	}
	return -1, fmt.Errorf("column not found")
}

func main() {
	// create CLI
	// Define command-line flags
	inputFile := flag.String("input", "", "Input CSV file")
	outputFile := flag.String("output", "", "Output CSV file")
	flag.Parse()

	// TODO, make a proper cli, that doesn't return when encountering errors
	// Check if input file is provided
	if *inputFile == "" {
		log.Println("Please provide input CSV file using -input flag")
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
		log.Println("OS open failed: ", err)
		return
	}
	defer openInputFile.Close()

	// Parse the input CSV file
	reader := csv.NewReader(openInputFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Reader failed: ", err)
		return
	}

	columnIdx, err := getColumnIndex(records[0], "pick_location")
	if err != nil {
		return
	}
	// call module that processes CSV
	customAlgo.OrderCSVbyColumnIdx(columnIdx, records[1:])

	// Create the output CSV file
	newCSV, err := os.Create(*outputFile)
	if err != nil {
		log.Println("OS create failed: ", err)
		return
	}
	defer newCSV.Close()

	// Write the processed records to the output CSV file
	writer := csv.NewWriter(newCSV)
	// TODO: format it correctly before writing
	err = writer.WriteAll(records)
	if err != nil {
		log.Println("Writer failed: ", err)
		return
	}
}
