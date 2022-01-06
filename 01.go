package main

import "fmt"

func main() {
	arr := [10]int{
		879, 456, 745, 235, 123, 5, 8, 3, 1, 634,
	}
	// fmt.Println(selectSort(arr[0:10]))
	fmt.Println(BubbleSort(arr[0:10]))
}
