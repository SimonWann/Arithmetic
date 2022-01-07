package main

func insertSort(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		for j := i - 1; j >= 0 && slice[j+1] < slice[j]; j-- {
			swap(slice, j+1, j)
		}
	}
	return slice
}
