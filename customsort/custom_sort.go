package customsort

import (
	"clialgotool/utils"
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// sort the chars first and if duplicate, compare the number
func SortByColumnIdx(column int, matrix [][]string) {
	// sort chars first and numbers second
	sort.Slice(matrix, func(i, j int) bool {
		// extract chars
		charI := strings.Split(matrix[i][column], " ")
		charJ := strings.Split(matrix[j][column], " ")

		// if chars are equal
		if charI[0] == charJ[0] {
			// sort by ints
			if len(charI[1]) != len(charJ[1]) {
				// normalize: add 0
				if len(charI[1]) < len(charJ[1]) {
					charI[1] = "0" + charI[1]
				} else {
					charJ[1] = "0" + charJ[1]
				}
				return charI[1] < charJ[1]
			}
			return charI[1] < charJ[1]

		}
		if len(charI[0]) == len(charJ[0]) {
			return charI[0] < charJ[0]
		} else {
			return len(charI[0]) < len(charJ[0])
		}
	})
}

func MergeDuplicatesAsSums(keyColumnIdx int, sumColumnIdx int, input [][]string, skip int) ([][]string, error) {
	results := [][]string{}

	if len(input) == 0 {
		return results, nil
	}

	// skip head rows
	for i := 0; i < skip; i++ {
		results = append(results, input[i])
	}

	// set initial value to compare against
	results = append(results, input[skip])

	for i := skip + 1; i < len(input); i++ {
		inputRow := input[i]
		inputKeyValue := inputRow[keyColumnIdx]
		resultsKeyValue := results[len(results)-1][keyColumnIdx]
		if inputKeyValue == resultsKeyValue {
			inputQuantity, err := strconv.Atoi(inputRow[sumColumnIdx])
			if err != nil {
				return results, err
			}

			resultsQuantity, err := strconv.Atoi(results[len(results)-1][sumColumnIdx])
			if err != nil {
				return results, err
			}

			results[len(results)-1][sumColumnIdx] = strconv.Itoa(resultsQuantity + inputQuantity)
		} else {
			results = append(results, inputRow)
		}
	}
	return results, nil
}

// // / my custom qs
// func swap(arr []int, idx1, idx2 int) {
// 	temp := arr[idx1]
// 	arr[idx1] = arr[idx2]
// 	arr[idx2] = temp
// }

// func quicksort(arr []int, lo, hi int) {
// 	if lo >= hi {
// 		return
// 	}
// 	pivot := arr[hi]
// 	idx := lo
// 	for i := lo; i < hi; i++ {
// 		if arr[i] < pivot && i != lo {
// 			swap(arr, i, idx)
// 			idx++
// 		}
// 		if arr[i] < pivot && i == lo {
// 			idx++
// 		}

// 	}
// 	if arr[idx] > arr[hi] {
// 		swap(arr, hi, idx)
// 	}

// 	quicksort(arr, lo, idx-1)
// 	quicksort(arr, idx+1, hi)
// }

func normalizeProductCode(productCode string) string {
	if len(productCode) < 6 {
		// add leading zeros to product code until it's 6 characters long
		return strings.Repeat("0", 6-len(productCode)) + productCode
	}
	return productCode
}

func normalizeLocation(location string) string {
	if len(location) < 5 {
		locationSplit := strings.Split(location, " ")
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

func FindOptimalPath(input *csv.Reader, skip int) (utils.OrderAwareMap, error) {
	productCodeCol := 0
	quantityCol := 1
	locationCol := 2

	// key = normalized location and product code, value = quantity
	// this map will contain the deduplicated rows with the total sum of quantities for each unique location and product code
	// order will contain the keys of the results map in sorted order
	results := utils.OrderAwareMap{
		Values: make(map[string]utils.CSVRow),
	}

	index := 0
	for {
		row, err := input.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return results, err
		}
		index++

		if index <= skip {
			// skip the first X rows (headers)
			fmt.Printf("Skipping row: %v\n", row)
			continue
		}

		normalizedProductCode := normalizeProductCode(row[productCodeCol])
		normalizedLocation := normalizeLocation(row[locationCol])
		key := normalizedLocation + normalizedProductCode

		quantity, err := strconv.Atoi(row[quantityCol])
		if err != nil {
			return results, err
		}
		if val, ok := results.Values[key]; ok {
			val.Quantity = val.Quantity + quantity
			results.Values[key] = val
		} else {
			results.Values[key] = utils.CSVRow{
				ProductCode:  row[productCodeCol],
				Quantity:     quantity,
				PickLocation: row[locationCol],
			}
		}
	}

	results.SortKeys()
	return results, nil
}
