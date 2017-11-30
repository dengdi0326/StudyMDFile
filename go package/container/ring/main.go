package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(6)
	for  i:= 0; i<6; i++ {
		r.Value = i
		r = r.Next()
	}

	s := 0
	r.Do(func(p interface{}) {
		fmt.Println(p,p.(int))
		s += p.(int)
		r = r.Next()
	})

	fmt.Println(s)
}
