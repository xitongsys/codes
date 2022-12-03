package main

import (
	"fmt"
	"math/rand"
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
	for {
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
}

func (treap *Treap) Put(val interface{}) int {
	v := treap.root
	for {
		flag := treap.cmp(treap.vals[v], val)
		if flag > 0 {
			if treap.lefts[v] >= 0 {
				v = treap.lefts[v]
			} else {
				u := treap.newNode(val)
				treap.lefts[v] = u
				treap.pars[u] = v
				return u
			}

		} else if flag < 0 {
			if treap.rights[v] >= 0 {
				v = treap.rights[v]
			} else {
				u := treap.newNode(val)
				treap.rights[v] = u
				treap.pars[u] = v
				return u
			}

		} else {
			treap.vals[v] = val
			return v
		}
	}
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
			for p := treap.pars[u]; p >= 0 && treap.rights[u] == u; u = p {
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
			for p := treap.pars[u]; p >= 0 && treap.lefts[u] == u; u = p {
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

func (treap *Treap) newNode(val interface{}) int {
	key := rand.Int()
	u := treap.n
	if cd := len(treap.dels); cd > 0 {
		u = treap.dels[cd-1]
		treap.dels = treap.dels[:cd-1]
		treap.keys[u] = key
		treap.vals[u] = val
		treap.lefts[u] = -1
		treap.rights[u] = -1
		treap.pars[u] = -1
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
		treap.rights[u] = p
		if pp := treap.pars[p]; pp >= 0 {
			if treap.lefts[pp] == p {
				treap.lefts[pp] = u
			} else {
				treap.rights[pp] = u
			}
		}
	} else { // right rotate
		treap.rights[p] = treap.lefts[u]
		treap.lefts[u] = p
		if pp := treap.pars[p]; pp >= 0 {
			if treap.lefts[pp] == p {
				treap.lefts[pp] = u
			} else {
				treap.rights[pp] = u
			}
		}
	}
}

func (treap *Treap) rotateToTop(u int) {
	for p := treap.pars[u]; p >= 0 && treap.keys[p] <= treap.keys[u]; {
		treap.rotate(u)
	}
}

func (treap *Treap) rotateToBottom(u int) {
	for left, right := treap.lefts[u], treap.rights[u]; left >= 0 || right >= 0; {
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

///////////////////////////////////////

type Pair struct {
	num int
	ids []int
}

func cmp(a, b interface{}) int {
	return a.(Pair).num - b.(Pair).num
}

func twoSum(nums []int, target int) []int {
	mp := NewTreap(cmp)
	for i, a := range nums {
		p := Pair{
			num: a,
		}
		if mp.Get(p) < 0 {
			mp.Put(p)
		}

		p = mp.vals[mp.Get(p)].(Pair)
		p.ids = append(p.ids, i)
	}

	res := []int{-1, -1}

	for i, a := range nums {
		t := target - a
		p := Pair{
			num: t,
		}
		it := mp.Get(p)
		if it >= 0 {
			if t == a {
				ids := mp.vals[it].(Pair).ids
				if len(ids) > 1 {
					res[0] = ids[0]
					res[1] = ids[1]
					return res
				}

			} else {
				res[0] = i
				res[1] = mp.vals[it].(Pair).ids[0]
				return res
			}
		}
	}
	return res

}

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))

}
