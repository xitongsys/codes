package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// big heap
type Treap struct {
	n, root             int
	keys                []int
	vals                []interface{}
	lefts, rights, pars []int
	dels                []int
	cmp                 func(a, b interface{}) int
}

func NewTreap(cmp func(a, b interface{}) int) *Treap {
	treap := &Treap{
		root: -1,
		cmp:  cmp,
	}
	return treap
}

func (treap *Treap) Get(val interface{}) int {
	v := treap.root
	for v >= 0 {
		flag := treap.cmp(treap.vals[v], val)
		if flag > 0 {
			if treap.lefts[v] >= 0 {
				v = treap.lefts[v]
			} else {
				return -1
			}

		} else if flag < 0 {
			if treap.rights[v] >= 0 {
				v = treap.rights[v]
			} else {
				return -1
			}

		} else {
			return v
		}
	}
	return -1
}

func (treap *Treap) Put(val interface{}, key interface{}) int {
	v := treap.root
	for v >= 0 {
		flag := treap.cmp(treap.vals[v], val)
		if flag > 0 {
			if treap.lefts[v] >= 0 {
				v = treap.lefts[v]
			} else {
				u := treap.newNode(val, key)
				treap.lefts[v] = u
				treap.pars[u] = v
				return u
			}

		} else if flag < 0 {
			if treap.rights[v] >= 0 {
				v = treap.rights[v]
			} else {
				u := treap.newNode(val, key)
				treap.rights[v] = u
				treap.pars[u] = v
				return u
			}

		} else {
			treap.vals[v] = val
			return v
		}
	}

	u := treap.newNode(val, key)
	treap.root = u
	return u
}

func (treap *Treap) Remove(val interface{}) {
	u := treap.Get(val)
	if u < 0 {
		return
	}

	// set key = -1 to rotate it to leaf
	treap.keys[u] = -1
	treap.rotateToBottom(u)
	p := treap.pars[u]
	if p >= 0 {
		if treap.lefts[p] == u {
			treap.lefts[p] = -1
		} else {
			treap.rights[p] = -1
		}
	} else {
		treap.root = -1
	}
	treap.delNode(u)
}

func (treap *Treap) Next(u int) int {
	if right := treap.rights[u]; right >= 0 {
		u = right
		for treap.lefts[u] >= 0 {
			u = treap.lefts[u]
		}
		return u
	}

	if p := treap.pars[u]; p >= 0 {
		if treap.lefts[p] == u {
			return p
		} else {
			for p = treap.pars[u]; p >= 0 && treap.rights[p] == u; u, p = p, treap.pars[p] {
			}
			return p
		}
	} else {
		return -1
	}
}

func (treap *Treap) Prev(u int) int {
	if left := treap.lefts[u]; left >= 0 {
		u = left
		for treap.rights[u] >= 0 {
			u = treap.lefts[u]
		}
		return u
	}

	if p := treap.pars[u]; p >= 0 {
		if treap.rights[p] == u {
			return p
		} else {
			for p = treap.pars[u]; p >= 0 && treap.lefts[p] == u; u, p = p, treap.pars[p] {
			}
			return p
		}

	} else {
		return -1
	}
}

func (treap *Treap) LowerBound(val interface{}) int {
	u := treap.root
	for u >= 0 {
		flag := treap.cmp(treap.vals[u], val)
		if flag < 0 {
			u = treap.rights[u]
		} else if flag > 0 {
			if treap.lefts[u] <= 0 {
				return u
			}
			u = treap.lefts[u]
		} else {
			return u
		}
	}
	return -1
}

func (treap *Treap) UpperBound(val interface{}) int {
	u := treap.LowerBound(val)
	if u < 0 {
		return u
	}

	if treap.cmp(treap.vals[u], val) == 0 {
		u = treap.Next(u)
	}
	return u
}

func (treap *Treap) First() int {
	u := treap.root
	for u >= 0 && treap.lefts[u] >= 0 {
		u = treap.lefts[u]
	}
	return u
}

func (treap *Treap) Last() int {
	u := treap.root
	for u >= 0 && treap.rights[u] >= 0 {
		u = treap.rights[u]
	}
	return u
}

func (treap *Treap) newNode(val interface{}, keyi interface{}) int {
	key := 0
	if keyi == nil {
		key = rand.Int()
	} else {
		key = keyi.(int)
	}

	u := treap.n
	if cd := len(treap.dels); cd > 0 {
		u = treap.dels[cd-1]
		treap.dels = treap.dels[:cd-1]
		treap.keys[u] = key
		treap.vals[u] = val
		treap.lefts[u] = -1
		treap.rights[u] = -1
		treap.pars[u] = -1
		treap.n++
	} else {
		treap.keys = append(treap.keys, key)
		treap.vals = append(treap.vals, val)
		treap.lefts = append(treap.lefts, -1)
		treap.rights = append(treap.rights, -1)
		treap.pars = append(treap.pars, -1)
		treap.n++
	}

	return u
}

func (treap *Treap) delNode(u int) {
	treap.dels = append(treap.dels, u)
	treap.vals[u] = nil
	treap.n--
}

func (treap *Treap) rotate(u int) {
	p := treap.pars[u]
	if p < 0 {
		treap.root = u
		return
	}

	if treap.keys[p] >= treap.keys[u] {
		return
	}

	// left rotate
	if treap.lefts[p] == u {
		treap.lefts[p] = treap.rights[u]
		if treap.rights[u] >= 0 {
			treap.pars[treap.rights[u]] = p
		}
		treap.rights[u] = p
		if pp := treap.pars[p]; pp >= 0 {
			if treap.lefts[pp] == p {
				treap.lefts[pp] = u
			} else {
				treap.rights[pp] = u
			}
			treap.pars[u] = pp
			treap.pars[p] = u
		} else {
			treap.pars[u] = pp
			treap.pars[p] = u
			treap.root = u
		}

	} else { // right rotate
		treap.rights[p] = treap.lefts[u]
		if treap.lefts[u] >= 0 {
			treap.pars[treap.lefts[u]] = p
		}
		treap.lefts[u] = p
		if pp := treap.pars[p]; pp >= 0 {
			if treap.lefts[pp] == p {
				treap.lefts[pp] = u
			} else {
				treap.rights[pp] = u
			}
			treap.pars[u] = pp
			treap.pars[p] = u
		} else {
			treap.pars[u] = pp
			treap.pars[p] = u
			treap.root = u
		}
	}
}

func (treap *Treap) rotateToTop(u int) {
	for p := treap.pars[u]; p >= 0 && treap.keys[p] < treap.keys[u]; p = treap.pars[u] {
		treap.rotate(u)
	}
}

func (treap *Treap) rotateToBottom(u int) {
	for left, right := treap.lefts[u], treap.rights[u]; left >= 0 || right >= 0; left, right = treap.lefts[u], treap.rights[u] {
		if left >= 0 && right >= 0 {
			if treap.keys[left] >= treap.keys[right] { //left is bigger
				treap.rotate(left)
			} else { //right is bigger
				treap.rotate(right)
			}
		} else if left >= 0 {
			treap.rotate(left)
		} else if right >= 0 {
			treap.rotate(right)
		}
	}
}

func (treap *Treap) print(u int) {
	if u < 0 {
		fmt.Print("nil", ",")
		return
	}

	treap.print(treap.lefts[u])
	fmt.Print(u, "-", treap.vals[u], ",")
	treap.print(treap.rights[u])
}

///////////////////////////////////////

type Pair struct {
	a, i int
}

func findRelativeRanks(score []int) []string {
	mp := NewTreap(func(a, b interface{}) int { return b.(*Pair).a - a.(*Pair).a })
	for i, a := range score {
		p := &Pair{
			a: a,
			i: i,
		}

		it := mp.Get(p)

		if it < 0 {
			it = mp.Put(p, nil)
		}
	}

	n := len(score)
	res := make([]string, n)
	rank := 0
	it := mp.First()
	for it >= 0 {
		i := mp.vals[it].(*Pair).i
		if rank == 0 {
			res[i] = "Gold Medal"
		} else if rank == 1 {
			res[i] = "Silver Medal"
		} else if rank == 2 {
			res[i] = "Bronze Medal"
		} else {
			res[i] = strconv.FormatInt(int64(rank+1), 10)
		}
		rank++
		it = mp.Next(it)
	}
	return res
}

func main() {
	score := []int{5, 4, 3, 2, 1}
	fmt.Println(findRelativeRanks(score))
}
