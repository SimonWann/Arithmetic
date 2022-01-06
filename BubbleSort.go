package main

func BubbleSort(slice []int) []int {
	for i := len(slice); i >= 0; i-- {
		for j := 0; j < i-1; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j, j+1)
			}
		}
	}
	return slice
}
