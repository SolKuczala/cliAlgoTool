package main

import (
	"encoding/csv"
	"flag"
	"fmt"

	"os"
	"path/filepath"
	"strings"

	"clialgotool/customsort"
	"clialgotool/utils"

	log "github.com/sirupsen/logrus"
)

func init() {
	// set logrus as the default logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

// parseCliOptions initialize the cli tool, parses the command-line options and
// returns the input and output filenames.
func parseCliOptions() (string, string, bool, error) {
	// define command-line flags
	log.Info("Initializing cli tool")
	inputFlag := flag.String("input", "", "Input CSV file")
	outputFlag := flag.String("output", "", "Output CSV file")
	printFlag := flag.Bool("print", false, "Print the output CSV file to stdout")
	flag.Parse()

	if *inputFlag == "" {
		return "", "", false, fmt.Errorf("please provide input CSV file using '-input' flag")
	}
	inputFileName := *inputFlag
	outputFileName := *outputFlag
	printToStdout := *printFlag

	// set the default output filename if not provided
	if outputFileName == "" {
		absInputFilePath, err := filepath.Abs(inputFileName)
		if err != nil {
			return "", "", false, fmt.Errorf("failed to determine output file path:%v", err)
		}
		inputFilePath := strings.TrimSuffix(absInputFilePath, filepath.Ext(inputFileName))
		outputFileName = inputFilePath + "_output.csv"
	}

	return inputFileName, outputFileName, printToStdout, nil
}

// saveToFile saves the processed records to the output CSV file
func saveToFile(file *os.File, data utils.OrderAwareMap) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	var header []string
	header = append(header, data.GetHeader()...)
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing CSV header: %v", err)
	}

	for _, key := range data.GetOrder() {
		row := data.CSVdata[key]
		record := []string{row.ProductCode, fmt.Sprintf("%d", row.Quantity), row.PickLocation}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write row: %v", err)
		}
	}

	return nil
}

func main() {
	// init cli tool
	inputFileName, outputFileName, printToStdout, err := parseCliOptions()
	if err != nil {
		log.Fatalf("failed to init and parse cli tool:%v", err.Error())
		return
	}

	log.Infof("Reading from file %v", inputFileName)
	openInputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("failed to open input file:%v", err.Error())
		return
	}
	defer openInputFile.Close()

	reader := csv.NewReader(openInputFile)
	reader.FieldsPerRecord = 3
	optimalPathData, err := customsort.FindOptimalPath(reader, 1)
	if err != nil {
		log.Fatalf("failed to find optimal path:%v", err.Error())
	}

	if printToStdout {
		for _, key := range optimalPathData.GetOrder() {
			// print the results to stdout ordered
			if printToStdout {
				fmt.Printf("%+q\n", optimalPathData.CSVdata[key])
			}
		}
	}

	// create output file and save results
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("failed to create output file:%v", err)
		return
	}
	defer outputFile.Close()

	if err = saveToFile(outputFile, optimalPathData); err != nil {
		log.Fatalf("failed to save results:%v", err.Error())
		return
	}

	log.Infof("Successfully saved results to file %v", outputFileName)
}
