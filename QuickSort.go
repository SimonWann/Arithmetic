package main

import "math/rand"

func QuickSort(slice []int) {

}

func QuickSortHandler(slice []int, l int, r int) {
	if r >= l {
		return
	}
	tailIndex := len(slice) - 1
	mIndex := rand.Intn(len(slice))
	slice[tailIndex] = slice[tailIndex] ^ slice[mIndex]
	slice[mIndex] = slice[tailIndex] ^ slice[mIndex]
	slice[tailIndex] = slice[tailIndex] ^ slice[mIndex]
	ml, mr := partition(slice, l, r)
	QuickSortHandler(slice, l, ml-1)
	QuickSortHandler(slice, mr+1, r)
}

func partition(slice []int, l int, r int) (lR int, rR int) {
	less := l - 1 // < 右边的
	more := r     // > 左边的
	for l < more {
		if slice[l] < slice[r] {
			less += 1
			swap(slice, less, l)
			l++
		} else if slice[l] > slice[r] {
			more -= 1
			swap(slice, more, l)
			l++
		} else {
			l++
		}
	}
	swap(slice, r, more)
	return less + 1, more
}
