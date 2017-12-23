package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	fmt.Println(l)

	lone := l.PushBack(4)
	fmt.Println(lone.Value)

	ltwo := l.PushFront(6)
	fmt.Println(ltwo.Value)

	l.InsertBefore(3, lone)
	fmt.Println(lone.Value)

	l.InsertAfter(2, ltwo)
	fmt.Println(ltwo.Value)

	for  e := l.Front(); e!= nil ; e=e.Next() {
		fmt.Println(e)
		fmt.Println(e.Value)
	}
}
