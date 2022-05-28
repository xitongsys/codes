package main

import "fmt"

type IntType int

func (i IntType) Less(b RBKeyType) bool {
	return i < b.(IntType)
}

func (i IntType) Equal(b RBKeyType) bool {
	return i == b.(IntType)
}


func main() {
	t := NewRBTree()
	for a := 0; a < 10; a++ {
		t.Add(IntType(a), 1)
	}

	node := t.LowerBound(IntType(5))

	fmt.Println(node.Key, node.Value)

	return 
}
