package main

func SmallSum(slice []int) int {
	return MergeSort2(slice, 0, len(slice)-1)
}
func MergeSort2(slice []int, l int, r int) int {
	if l == r {
		return 0
	}
	m := l + ((r - l) >> 1)
	return MergeSort2(slice, l, m) + MergeSort2(slice, m+1, r) + Merge2(slice, l, r, m)
}
func Merge2(slice []int, l int, r int, m int) int {
	i := 0
	pl := l
	pr := m + 1
	temp := make([]int, len(slice), cap(slice))
	sum := 0
	for pl <= m && pr <= r {
		if slice[pl] < slice[pr] {
			temp[i] = slice[pl]
			i++
			cnt := r - pr + 1
			sum += (slice[pl] * cnt)
			pl++
		} else {
			temp[i] = slice[pr]
			i++
			pr++
		}
	}
	for pl <= m {
		temp[i] = slice[pl]
		i++
		pl++
	}
	for pr <= r {
		temp[i] = slice[pr]
		i++
		pr++
	}
	return sum
}
