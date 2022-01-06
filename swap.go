package main

func swap(slice []int, a int, b int) bool {
	slice[a] = slice[a] ^ slice[b]
	slice[b] = slice[a] ^ slice[b]
	slice[a] = slice[a] ^ slice[b]
	return true
}
