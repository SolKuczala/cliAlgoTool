package main

import (
	"encoding/csv"
	"flag"
	"fmt"

	"os"
	"path/filepath"
	"strings"

	"clialgotool/customsort"

	log "github.com/sirupsen/logrus"
)

func init() {
	// set logrus as the default logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

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
	log.Info("Initializing cli tool")
	inputFile := flag.String("input", "", "Input CSV file")
	outputFile := flag.String("output", "", "Output CSV file")
	printToStdout := flag.Bool("print", false, "Print the output CSV file to stdout")
	flag.Parse()

	// Check if input file is provided
	if *inputFile == "" {
		log.Fatalf("Please provide input CSV file using -input flag")
		return
	}
	// Set the default output filename if not provided
	if *outputFile == "" {
		inputFileName := strings.TrimSuffix(filepath.Base(*inputFile), filepath.Ext(*inputFile))
		*outputFile = inputFileName + "_output.csv"
	}

	log.Infof("Reading file %v", *inputFile)

	// Read input CSV file
	openInputFile, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("OS open failed:%v", err)
		return
	}
	defer openInputFile.Close()

	// Parse the input CSV file
	reader := csv.NewReader(openInputFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Reader failed:%v", err)
		return
	}

	columnIdx, err := getColumnIndex(records[0], "pick_location")
	if err != nil {
		log.Fatalf("getColumnIndex failed:%v", err)
		return
	}

	// call module that processes CSV
	customsort.SortCSVbyColumnIdx(columnIdx, records[1:])

	// Create the output CSV file
	newCSV, err := os.Create(*outputFile)
	if err != nil {
		log.Fatalf("OS create failed:%v", err)
		return
	}
	defer newCSV.Close()

	// Write the processed records to the output CSV file
	writer := csv.NewWriter(newCSV)
	err = writer.WriteAll(records)
	if err != nil {
		log.Fatalf("Writer failed:%v", err)
		return
	}
	log.Infof("Successfully written to file %v", *outputFile)

	if *printToStdout {
		for _, record := range records {
			fmt.Printf("%+q\n", record)
		}
	}
}
