package main

import "fmt"

func main() {
	t := NewRBTree(func(a, b interface{})int{return a.(int) - b.(int)})
	pq := NewPriorityQueue(func(a, b interface{})int{return a.(int) - b.(int)})

	for a := 10; a >= 0; a-- {
		t.Put(a, a)
		pq.Put(a)
	}

	fmt.Println(t.String())

	fmt.Println(pq.Get())
	fmt.Println(pq.Get())
	fmt.Println(pq.Get())




	return 
}
