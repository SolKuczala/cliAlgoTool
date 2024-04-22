package customsort

import (
	"clialgotool/utils"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	// ProductCodeCol is the column index of the product code
	productCode = "product_code"
	// QuantityCol is the column index of the quantity
	quantity = "quantity"
	// LocationCol is the column index of the location
	pickLocation = "pick_location"
)

// normalizeProductCode add leading spaces to product code to 6 characters long
// in order to be comparable.
func normalizeProductCode(productCode string) string {
	if len(productCode) < 6 {
		return strings.Repeat("0", 6-len(productCode)) + productCode
	}
	return productCode
}

// normalizePickLocation to 4 characters to be able to separate locations at combination key.
func normalizePickLocation(location string) string {
	if len(location) < 6 {
		locationSplit := strings.Split(location, " ")
		for i, loc := range locationSplit {
			locationSplit[i] = strings.Trim(loc, " ")
		}
		if len(locationSplit[0]) < 2 {
			// add leading spaces to location until it's 2 characters long
			// this will allow to sort characters in the order of A-AZ
			locationSplit[0] = strings.Repeat(" ", 2-len(locationSplit[0])) + locationSplit[0]
		}
		if len(locationSplit[1]) < 2 {
			// add leading zeros to location until it's 2 characters long
			// this will allow to sort numbers in the order of 01-10
			locationSplit[1] = strings.Repeat("0", 2-len(locationSplit[1])) + locationSplit[1]
		}
		return locationSplit[0] + locationSplit[1]
	}
	return location
}

// normalizeLine the row line by trimming whitespace from each field
func normalizeLine(line []string) {
	for i := range line {
		line[i] = strings.TrimSpace(line[i])
	}
}

// FindOptimalPath reads the input CSV file and returns a map with the optimal path,
// the order and the headers of the CSV file in the OrderAwareMap struct.
func FindOptimalPath(input *csv.Reader, skip int) (utils.OrderAwareMap, error) {
	// we are assuming that the input CSV file has the following columns:
	var productCodeCol int
	var quantityCol int
	var locationCol int

	// key = normalized location and product code, value = quantity
	// this map will contain the deduplicated rows with the total sum of quantities for
	// each unique location and product code order will contain the keys of the results
	// map in sorted order
	results := utils.OrderAwareMap{
		CSVdata: make(map[string]utils.CSVRow),
	}

	index := 0
	for {
		row, err := input.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return results, fmt.Errorf("failed to read input:%v", err.Error())
		}

		normalizeLine(row)
		index++
		if index <= skip {
			results.SetHeader(row)

			for i, col := range row {
				if col == productCode {
					productCodeCol = i
				}
				if col == quantity {
					quantityCol = i
				}
				if col == pickLocation {
					locationCol = i
				}
			}

			// and skip the first X rows (headers)
			log.Infof("Skipping row: %+q\n", row)
			continue
		}

		// normalize product code and location to create key
		normalizedProductCode := normalizeProductCode(row[productCodeCol])
		normalizedLocation := normalizePickLocation(row[locationCol])
		combinedKey := normalizedLocation + normalizedProductCode

		// convert quantity to int
		quantity, err := strconv.Atoi(row[quantityCol])
		if err != nil {
			return results, fmt.Errorf("findOptimalPath failed: %v", err)
		}
		// if the key already exists, add the quantity to the existing value
		// and update row. Otherwise, create a new row.
		if val, ok := results.CSVdata[combinedKey]; ok {
			val.Quantity += quantity
			results.CSVdata[combinedKey] = val
		} else {
			results.CSVdata[combinedKey] = utils.CSVRow{
				ProductCode:  row[productCodeCol],
				Quantity:     quantity,
				PickLocation: row[locationCol],
			}
		}
	}
	results.SortKeys()
	return results, nil
}
