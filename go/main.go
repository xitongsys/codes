package main

import "fmt"

type IntType int

func (a IntType) Compare(b RBKeyType) int {
	return int(a - b.(IntType))
}


func main() {
	t := NewRBTree()
	for a := 10; a >= 0; a-- {
		t.Add(IntType(a), 1)
		t.Remove(IntType(a+1))
		fmt.Println(t.Check())
	}

	fmt.Println(t.String())


	return 
}
