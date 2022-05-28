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
	for a := 10; a >= 0; a-- {
		t.Add(IntType(a), 1)
		t.Remove(IntType(a+1))
		fmt.Println(t.Check())
	}

	fmt.Println(t.String())


	return 
}
