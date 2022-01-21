package main

// 容器 -> 桶
// 通过不同进制数的位数进行排序

func baseSort(arr []int, digit int, L int, R int) {
	i, j := 0, 0
	radix := 10
	// 辅助空间
	bucket := make([]int, R-L+1)
	for d := 1; d <= digit; d++ {
		count := make([]int, radix)
		for i = L; i <= R; i++ {
			j = getDigit(arr[i], digit)
			count[j]++
		}
		for i := 1; i < len(count); i++ {
			count[i] = count[i] + count[i-1]
		}

		for i := R; i >= L; i-- {
			j := getDigit(arr[i], digit)
			bucket[count[j]-1] = arr[i]
			count[j]--
		}

		for i, j := L, 0; i <= R; i++ {
			arr[i] = bucket[j]
			j++
		}
	}
}

func getDigit(element int, digit int) int {
	unit := 10
	for i := 0; i < digit; i++ {
		unit *= 10
	}
	return element % unit
}
