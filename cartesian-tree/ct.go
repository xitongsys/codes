package main

import (
	"fmt"
	"strconv"
)

var N int = 5

var root = -1
var vals []int = make([]int, N)
var lefts, rights []int = make([]int, N), make([]int, N)

// big heap
func create() {
	sk := make([]int, 0)
	for i := 0; i < N; i++ {
		rights[i] = -1
		lefts[i] = -1

		v := vals[i]
		c := -1
		for len(sk) > 0 && vals[sk[len(sk)-1]] <= v {
			c = sk[len(sk)-1]
			sk = sk[:len(sk)-1]
		}

		if len(sk) > 0 {
			p := sk[len(sk)-1]
			rights[p] = i
		} else {
			root = i
		}

		lefts[i] = c
		sk = append(sk, i)
	}
}

func toString(u int) string {
	if u < 0 {
		return "nil"
	}
	lefts := toString(lefts[u])
	rights := toString(rights[u])

	mids := "(" + strconv.FormatInt(int64(u), 10) + "," + strconv.FormatInt(int64(vals[u]), 10) + ")"

	res := "[" + lefts + "]-" + mids + "-[" + rights + "]"
	return res
}

// use 1: range max
func rangeMax(bgn, end int) int {
	u := root
	for {
		if u >= bgn && u <= end {
			return vals[u]
		}

		if u < bgn {
			u = rights[u]
		} else {
			u = lefts[u]
		}
	}
}

func main() {
	for i := 0; i < N; i++ {
		vals[i] = N - i
	}

	create()
	fmt.Println(toString(root))
	fmt.Println(rangeMax(1, 2))
}
