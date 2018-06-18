package main

import "fmt"

func Selection(queue []int) []int {
	for i := 0; i < len(queue); i++ {
		for j := i+1; j<len(queue); j++ {
			if queue[i] < queue[j] {
				queue[i], queue[j] = queue[j], queue[i]
			}
		}
	}
	return queue
}

func Insertion(queue []int) []int {
	for i := 1; i < len(queue); i++ {
		for j := i-1; j > 0; j-- {
			if queue[j] < queue[i] {
				queue[j], queue[i] = queue[i], queue[j]
			}
			break
		}
	}
	return queue
}

func Hill(queue []int) []int {
	h := 1
	for h < len(queue) / 3 {
		h = h * 3 + 1
	}
	for i := h; i < len(queue); i++ {
		for j := i; j >= h; j = j - h {
			if queue[j] < queue[j - h] {
				queue[j], queue[j - h] = queue[j - h], queue[j]
			}
		}
		h = h / 3
		if h == 0 {
			break
		}
	}
	return queue
}

func main() {
	//fmt.Println("enter queue: ")
	var queue []int
	//fmt.Scanln(&queue)

	queue = []int{1,3,5,6,4,8,9}

	fmt.Println(Hill(queue))
	fmt.Println(Selection(queue))
	fmt.Println(Insertion(queue))
}