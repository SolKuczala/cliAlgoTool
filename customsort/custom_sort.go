package customsort

import (
	"clialgotool/utils"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// // sort the chars first and if duplicate, compare the number
// func SortByColumnIdx(column int, matrix [][]string) {
// 	// sort chars first and numbers second
// 	sort.Slice(matrix, func(i, j int) bool {
// 		// extract chars
// 		charI := strings.Split(matrix[i][column], " ")
// 		charJ := strings.Split(matrix[j][column], " ")

// 		// if chars are equal
// 		if charI[0] == charJ[0] {
// 			// sort by ints
// 			if len(charI[1]) != len(charJ[1]) {
// 				// normalize: add 0
// 				if len(charI[1]) < len(charJ[1]) {
// 					charI[1] = "0" + charI[1]
// 				} else {
// 					charJ[1] = "0" + charJ[1]
// 				}
// 				return charI[1] < charJ[1]
// 			}
// 			return charI[1] < charJ[1]

// 		}
// 		if len(charI[0]) == len(charJ[0]) {
// 			return charI[0] < charJ[0]
// 		} else {
// 			return len(charI[0]) < len(charJ[0])
// 		}
// 	})
// }

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
	// TODO: tendria que leer las columnas para determinar sus indices
	productCodeCol := 0
	quantityCol := 1
	locationCol := 2

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
			fmt.Printf("Skipping row: %+q\n", row)
			// skip the first X rows (headers)
			continue
		}

		// normalize product code and location
		normalizedProductCode := normalizeProductCode(row[productCodeCol])
		normalizedLocation := normalizeLocation(row[locationCol])
		combinedKey := normalizedLocation + normalizedProductCode

		// convert quantity to int
		quantity, err := strconv.Atoi(row[quantityCol])
		if err != nil {
			return results, fmt.Errorf("findOptimalPath failed: %v", err)
		}
		// if the key already exists, add the quantity to the existing value
		// and update row.
		if val, ok := results.CSVdata[combinedKey]; ok {
			// TODO: try val.Quantity += quantity?
			val.Quantity = val.Quantity + quantity
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
