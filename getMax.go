package main

func GetMax(slice []int, start int, end int) (result int) {
	if start == end {
		return slice[start]
	}
	middle := ((end - start) >> 2) + start
	return Max(GetMax(slice, start, middle), GetMax(slice, middle+1, end))
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
