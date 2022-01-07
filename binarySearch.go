package main

func binarySearch(slice []int, target int) int {
	currentIndex := len(slice) / 2
	for slice[currentIndex] != target && (currentIndex > 0 || currentIndex < len(slice)) {
		currentIndex /= 2
	}
	return currentIndex
}

func binarySearchByComposer(slice []int, composer func(target int, index int, orginSlice []int) int) int {
	currentIndex := len(slice) / 2
	compare := composer(slice[currentIndex], currentIndex, slice)
	for (compare != 0) && (currentIndex > 0 && currentIndex < len(slice)) {
		// fmt.Println(currentIndex)
		// currentIndex /= 2
		if compare > 0 {
			currentIndex = currentIndex + (currentIndex / 2)
		} else {
			currentIndex = currentIndex - currentIndex/2
		}
		compare = composer(slice[currentIndex], currentIndex, slice)
	}
	return currentIndex
}
