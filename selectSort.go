package main

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
	return arr
}
