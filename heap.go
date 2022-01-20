package main

import "fmt"

// 堆是完全二叉树
// 数组可以转换为完全二叉树
// 完全二叉树中 节点 与 数组 的对应关系:
/*
	左子节点: i * 2 + 1
	右子节点: i * 2 + 2
	父节点:		i / 2 -1
*/
// func

type Heap struct {
	Arr      []int
	HeapSize int
}

func (heap *Heap) Init() {
	heap.HeapSize = 0
	for i := 0; i < len(heap.Arr); i++ {
		// fmt.Println("---start---")

		heap.Insert(heap.Arr, i)
		heap.HeapSize++
		// fmt.Println("---end---")
	}
}

// 从尾到头
func (heap *Heap) Insert(arr []int, index int) {
	for index > 0 {
		parent := arr[(index-1)/2]
		if arr[index] > parent {
			swap(arr, index, (index-1)/2)
		}
		// fmt.Println(arr)
		index = (index - 1) / 2
	}
}
func (heap *Heap) Add(value int) {
	len := len(heap.Arr)
	heap.Arr = append(heap.Arr, value)
	heap.Insert(heap.Arr, len)
	heap.HeapSize++
}

// 从头到尾
func (heap *Heap) Heapify(arr []int, index int, rIndex int) {
	heap.HeapSize--
	// removeOne := arr[index]
	// lastOne := rIndex
	// swap(arr, lastOne, index)

	for (index*2 + 1) < rIndex {
		var maxOne int
		if arr[index*2+1] > arr[index*2+2] {
			maxOne = index*2 + 1
		} else {
			maxOne = index*2 + 2

		}
		if arr[maxOne] < arr[index] {
			maxOne = index
		}
		if maxOne != index {
			swap(arr, maxOne, index)
		}
		index = (index*2 + 1)
	}
}
func (heap *Heap) sort() {
	for i := len(heap.Arr) - 1; i >= 0; i-- {
		heap.Heapify(heap.Arr, 0, i)
		fmt.Println(heap.Arr)
		swap(heap.Arr, 0, i)
	}
}
func (heap *Heap) removeHead() {

}
