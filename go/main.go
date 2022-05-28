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
	for a := 0; a < 2; a++ {
		t.Add(IntType(a), 1)
	}
	t.Remove(IntType(0))

	fmt.Println(t.String())


	return 
}
