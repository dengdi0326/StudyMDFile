package main

import (
	"fmt"
)

func BinarySearch(key int, queue []int) int{
	head := 0
	tail := len(queue) - 1
	for {
		index := (head + tail) / 2
		if key > queue[index] {
			head = index + 1
			continue
		}
		if key < queue[index] {
			tail = index - 1
			continue
		}
		if key == queue[index] {
			return index
		}

		return -1
	}
}

func main() {
	var queue []int
	var key int
	queue = []int{1,3,4,5,6,8,9}
	fmt.Scanln(&key)
	fmt.Println("index is :", BinarySearch(key, queue))
}
