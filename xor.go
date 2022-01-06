package main

func getOneNum(slice []int) int {
	var result int
	for _, v := range slice {
		result ^= v
	}
	return result
}

func getTwoNum(slice []int) (int, int) {
	var resultA int
	var resultB int
	// ero1 两数 不相等的部分
	var ero int
	var ero2 int
	for _, v := range slice {
		ero ^= v
	}
	for _, v := range slice {
		if ero {
			ero2 ^= v
		}
	}
	// ero2
	return resultA, resultB
}
