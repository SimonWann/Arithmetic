package main

func MergeSort(slice []int, L int, R int) {
	if L == R {
		return
	}
	M := L + ((R - L) >> 1)
	MergeSort(slice, L, M)
	MergeSort(slice, M+1, R)
}

func Merge(slice []int, L int, M int, R int) {
	temp := make([]int, len(slice), cap(slice))
	i := 0
	pL := L
	pR := R
	for pL <= M && pR <= R {
		if slice[pL] <= slice[pR] {
			temp[i] = slice[pL]
			i++
			pL++
		} else {
			temp[i] = slice[pR]
			i++
			pR++
		}
	}
	for pL <= M {
		temp[i] = slice[pL]
		i++
		pL++
	}
	for pR <= R {
		temp[i] = slice[pR]
		i++
		pR++
	}
	for k, v := range temp {
		slice[L+k] = v
	}
}
