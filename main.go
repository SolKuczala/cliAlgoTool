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

func saveToFile(filename string, records [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file:%v", err)
		return err
	}
	defer file.Close()

	// Write the processed records to the output CSV file
	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	if err != nil {
		log.Fatalf("Failed to write:%v", err)
		return err
	}
	return nil
}

func parseCliOptions() (string, string, bool, error) {
	// create CLI
	// Define command-line flags
	log.Info("Initializing cli tool")
	inputFlag := flag.String("input", "", "Input CSV file")
	outputFlag := flag.String("output", "", "Output CSV file")
	printFlag := flag.Bool("print", false, "Print the output CSV file to stdout")
	flag.Parse()

	// Check if input file is provided
	if *inputFlag == "" {
		return "", "", false, fmt.Errorf("Please provide input CSV file using -input flag")
	}
	inputFileName := *inputFlag
	outputFileName := *outputFlag
	printToStdout := *printFlag

	// Set the default output filename if not provided
	if outputFileName == "" {
		absInputFilePath, err := filepath.Abs(inputFileName)
		if err != nil {
			return "", "", false, fmt.Errorf("Failed to determine output file path:%v", err)
		}
		inputFilePath := strings.TrimSuffix(absInputFilePath, filepath.Ext(inputFileName))
		outputFileName = inputFilePath + "_output.csv"
	}

	return inputFileName, outputFileName, printToStdout, nil
}

func main() {
	inputFileName, outputFileName, printToStdout, err := parseCliOptions()
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Reading from file %v", inputFileName)
	// Read input CSV file
	openInputFile, err := os.Open(inputFileName)
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

	locationColumnIdx, err := getColumnIndex(records[0], "pick_location")
	if err != nil {
		log.Fatalf("getColumnIndex failed:%v", err)
		return
	}

	quantityColumnIdx, err := getColumnIndex(records[0], "quantity")
	if err != nil {
		log.Fatalf("getColumnIndex failed:%v", err)
		return
	}

	// call module that processes CSV
	customsort.SortByColumnIdx(locationColumnIdx, records[1:])

	results, err := customsort.MergeDuplicatesAsSums(locationColumnIdx, quantityColumnIdx, records[1:])
	if err != nil {
		log.Fatalf("Failed to merge duplicates")
	}

	if err = saveToFile(outputFileName, results); err != nil {
		log.Fatalf("Failed to save results")
	}
	log.Infof("Successfully saved results to file %v", outputFileName)

	if printToStdout {
		for _, result := range results {
			fmt.Printf("%+q\n", result)
		}
	}
}
