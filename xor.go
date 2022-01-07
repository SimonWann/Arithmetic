package main

func getOneNum(slice []int) int {
	var result int
	for _, v := range slice {
		result ^= v
	}
	return result
}

func getTwoNum(slice []int) (int, int) {
	// ero1 两数 不相等的部分
	var ero int
	var ero2 int
	for _, v := range slice {
		ero ^= v
	}
	var rightOne = ero & ^ero + 1
	for _, v := range slice {
		if (v & rightOne) == 0 {
			ero2 ^= v
		}
	}
	// ero2
	return ero2, ero2 ^ ero
}
