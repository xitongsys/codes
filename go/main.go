package main

import "fmt"

func main() {
	t := NewRBTree(func(a, b interface{})int{return a.(int) - b.(int)})
	for a := 10; a >= 0; a-- {
		t.Put(a, a)
	}

	fmt.Println(t.String())


	return 
}
