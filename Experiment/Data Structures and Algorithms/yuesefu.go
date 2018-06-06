package main

import (
	"fmt"
)

type node struct {
	num  int
	data int
	next *node
}

func main() {
	var head node
	head.next = nil

	var p *node
	p = &head
	head.data = 1

	var n, m int
	fmt.Scanln(&n, &m)

	for i := 1; i < m ; i++ {
		var r node
		r.data = i + 1

		p.next = &r
		p = &r
	}
	p.next = &head
	q := p.next

	for a:=1; ; a++{
		if ((a+1) % n == 0) {
			fmt.Println(q.next.data)
			q.next = q.next.next
			m -= 1
			a = a+1
		}

		q = q.next

		if(m == 0) {
			break
		}
	}
}