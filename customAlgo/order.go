package customAlgo

import (
	"sort"
	"strings"
)

// order the chars first and if duplicate, compare the number
func OrderCSVbyColumnIdx(column int, matrix [][]string) {
	// order chars first and numbers second
	sort.Slice(matrix, func(i, j int) bool {
		// extract chars
		charI := strings.Split(matrix[i][column], " ")
		charJ := strings.Split(matrix[j][column], " ")

		// if chars are equal
		if charI[0] == charJ[0] {
			// sort  by ints
			if len(charI[1]) != len(charJ[1]) {
				// add 0
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
