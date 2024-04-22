package customsort

import (
	"clialgotool/utils"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	// ProductCodeCol is the column index of the product code
	productCode = "product_code"
	// QuantityCol is the column index of the quantity
	quantity = "quantity"
	// LocationCol is the column index of the location
	pickLocation = "pick_location"
)

// Normalize product code to 6 characters to separate codes at combination key
func normalizeProductCode(productCode string) string {
	if len(productCode) < 6 {
		// add leading zeros to product code until it's 6 characters long
		return strings.Repeat("0", 6-len(productCode)) + productCode
	}
	return productCode
}

// Normalize location to 4 characters to separate locations at combination key.
func normalizeLocation(location string) string {
	if len(location) < 6 {
		locationSplit := strings.Split(location, " ")
		for i, loc := range locationSplit {
			locationSplit[i] = strings.Trim(loc, " ")
		}

		// TODO: no se si el primero hace diff
		if len(locationSplit[0]) < 2 {
			// add leading spaces to location until it's 2 characters long
			locationSplit[0] = strings.Repeat(" ", 2-len(locationSplit[0])) + locationSplit[0]
		}
		if len(locationSplit[1]) < 2 {
			// add leading zeros to location until it's 2 characters long
			locationSplit[1] = strings.Repeat("0", 2-len(locationSplit[1])) + locationSplit[1]
		}
		return locationSplit[0] + locationSplit[1]
	}
	return location
}

// normalizeLine the row line by trimming whitespace from each field
func normalizeLine(line []string) {
	// Trim whitespace from each field
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
			fmt.Printf("Skipping row: %+q\n", row)
			// skip the first X rows (headers)
			continue
		}

		// normalize product code and location to create key
		normalizedProductCode := normalizeProductCode(row[productCodeCol])
		normalizedLocation := normalizeLocation(row[locationCol])
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
