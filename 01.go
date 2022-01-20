package main

import (
	"fmt"
)

func main() {
	// arr := [10]int{
	// 	879, 456, 745, 235, 123, 5, 8, 3, 1, 634,
	// }
	// // [1 3 5 8 123 235 456 634 745 879]
	// arr2 := []int{
	// 	1, 2, 3, 4, 5, 5, 4, 3,
	// }
	// arr3 := []int{
	// 	1, 3, 4, 5, 5, 5, 6, 7, 7, 8, 8,
	// }
	// for i := 0; i < 1_000; i++ {
	// 	testSlice := make([]int, 10)
	// 	for j := 0; j < 10; j++ {
	// 		testSlice[j] = rand.Int()
	// 	}
	// 	var testSlice2 []int
	// 	testSlice2 = make([]int, 10)
	// 	copy(testSlice2, testSlice)
	// 	fmt.Println("i:", i, Equal(BubbleSort(testSlice), insertSort(testSlice2)))
	// 	// fmt.Println(testSlice, testSlice2)
	// }
	// fmt.Println(selectSort(arr[0:10]))
	// fmt.Println(BubbleSort(arr[0:10]))
	// fmt.Println(insertSort(arr[0:10]))
	// tempS1 := arr[0:10]
	// MergeSort(tempS1, 0, 9)
	// sliceS2 := []int{
	// 	1, 3, 4, 2, 5,
	// }
	// fmt.Println("MergeSorr: ", SmallSum(sliceS2))
	// fmt.Println("BinarySearch: ", binarySearchByComposer(arr3, searchComposer))
	// fmt.Println(getTwoNum(arr2))
	// fmt.Println("getMax: ", GetMax(arr3, 0, len(arr3)-1))
	var arr []int = []int{
		5, 8, 1, 3, 0, 9, 5,
	}
	var heap *Heap = &Heap{
		Arr: arr,
	}
	heap.Init()
	fmt.Println(heap.Arr, heap.HeapSize)
	heap.Add(100)
	fmt.Println(heap.Arr, heap.HeapSize)
	// heap.Heapify(heap.Arr, 0, len(heap.Arr)-1)
	// fmt.Println(heap.Arr, heap.HeapSize)
	heap.sort()
	fmt.Println(heap.Arr, heap.HeapSize)

}

func searchComposer(v int, k int, origin []int) int {
	if k-1 < 0 {
		return 0
	}
	fmt.Println("SearchByComposer", v, k, "left:", origin[k-1])
	if v >= 5 && origin[k-1] < 5 {
		return 0
	} else if v >= 5 && origin[k-1] >= 5 {
		return -1
	} else {
		return 1
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
