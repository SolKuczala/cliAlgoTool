package customalgo

func Order([][]string) {
	// TODO : how to apply quicksort to a matrix?

}

func swap(arr []int, idx1, idx2 int) {
	temp := arr[idx1]
	arr[idx1] = arr[idx2]
	arr[idx2] = temp
}

func quicksort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := arr[hi]
	idx := lo
	for i := lo; i < hi; i++ {
		if arr[i] < pivot && i != lo {
			swap(arr, i, idx)
			idx++
		}
		if arr[i] < pivot && i == lo {
			idx++
		}

	}
	if arr[idx] > arr[hi] {
		swap(arr, hi, idx)
	}

	quicksort(arr, lo, idx-1)
	quicksort(arr, idx+1, hi)
}
