package quicksort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	var less []int
	var greater []int

	for _, v := range arr[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	sortedLess := QuickSort(less)
	sortedGreater := QuickSort(greater)

	return append(append(sortedLess, pivot), sortedGreater...)
}
