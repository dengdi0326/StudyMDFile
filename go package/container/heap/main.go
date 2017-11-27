package main

import (
	"fmt"
	"container/heap"
)

type IntHeap []int

func(h IntHeap) Len() int {return len(h)}
func(h IntHeap) Less(i,l int) bool {return h[i] < h[l]}
func(h IntHeap) Swap(i,l int) {h[i], h[l] = h[l], h[i]}

func(h *IntHeap) Push(x interface{}) {
	 *h = append(*h, x.(int))
}

func(h *IntHeap) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func Heap() {
	h := &IntHeap{2,3,4,5,5}
	fmt.Println(h)

	//heap.Init(h)

	//h.Push 和 heap.Push ，前者向末尾加一个元素，后者会有一定规律排序
	h.Push(3)
	//heap.Push(h, 3)
	fmt.Println(h)

	//h.Pop 和 heap.Pop ，前者向末尾删除一个元素，后者删除最小的数
	fmt.Println(h.Pop())
	fmt.Println(h)

	fmt.Println("package example:")
	w := &IntHeap{5,6,9,5,5}
	fmt.Println(w)

	heap.Init(w)
	fmt.Println(w)

	heap.Push(w, 9)
	fmt.Println(w)

	fmt.Println(heap.Pop(w))
}

type Item struct {
	value     string
	priority  int
	index     int
}

type PriorityQueue []*Item

func(pq PriorityQueue) Len() int {return len(pq) }
func(pq PriorityQueue) Less(i,j int) bool {return pq[i].priority > pq[j].priority}
func(pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func(pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func(pq *PriorityQueue) Pop() interface{} {
	o := *pq
	n := len(o)
	item := o[n-1]
	item.index = -1
	*pq = o[0:n-1]
	return item
}

func(pq *PriorityQueue) Update(item *Item, priority int, value string) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:value,
			priority:priority,
			index:i,
		}
		i++
	}
	heap.Init(&pq)

	item := &Item{
		value:"apple",
		priority:1,
	}
	heap.Push(&pq, item)
	fmt.Println(pq)

	pq.Update(item, 5, item.value)
	fmt.Println(pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Println(item)
	}
}