package customsort

import (
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

func MergeDuplicatesAsSums(keyColumnIdx int, sumColumnIdx int, input [][]string) ([][]string, error) {
	results := [][]string{}

	if len(input) == 0 {
		return results, nil
	}

	results = append(results, input[0])
	resultsIdx := 0

	for i := 1; i < len(input); i++ {
		inputRow := input[i]
		inputKeyValue := inputRow[keyColumnIdx]
		resultsKeyValue := results[resultsIdx][keyColumnIdx]
		if inputKeyValue == resultsKeyValue {
			inputQuantity, err := strconv.Atoi(inputRow[sumColumnIdx])
			if err != nil {
				return results, err
			}

			resultsQuantity, err := strconv.Atoi(results[resultsIdx][sumColumnIdx])
			if err != nil {
				return results, err
			}

			results[resultsIdx][sumColumnIdx] = strconv.Itoa(resultsQuantity + inputQuantity)
		} else {
			results = append(results, inputRow)
			resultsIdx++
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
