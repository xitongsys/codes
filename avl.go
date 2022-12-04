package main

import "fmt"

type AVLTree struct {
	N, Root int
	Vals    []interface{}

	heights             []int
	lefts, rights, pars []int
	dels                []int
	cmp                 func(a, b interface{}) int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewAVLTree(cmp func(a, b interface{}) int) *AVLTree {
	tree := &AVLTree{
		Root: -1,
		cmp:  cmp,
	}
	return tree
}

func (tree *AVLTree) Put(val interface{}) int {
	u := tree.Root
	if u < 0 { // empty tree
		u = tree.newNode(val)
		tree.Root = u
		return u
	}

	return tree.put(u, val)
}

func (tree *AVLTree) put(u int, val interface{}) int {
	f := tree.cmp(tree.Vals[u], val)
	v := u
	if f > 0 { //val < node.val,to left
		if left := tree.lefts[u]; left >= 0 {
			v = tree.put(left, val)
		} else {
			v = tree.newNode(val)
			tree.lefts[u] = v
			tree.pars[v] = u
		}
	} else if f < 0 {
		if right := tree.rights[u]; right >= 0 {
			v = tree.put(right, val)
		} else {
			v = tree.newNode(val)
			tree.rights[u] = v
			tree.pars[v] = u
		}
	} else {
		tree.Vals[u] = val
		return u
	}

	tree.fixHeight(u)
	tree.balance(u)
	return v
}

func (tree *AVLTree) Get(val interface{}) int {
	v := tree.Root
	for v >= 0 {
		flag := tree.cmp(tree.Vals[v], val)
		if flag > 0 {
			if tree.lefts[v] >= 0 {
				v = tree.lefts[v]
			} else {
				return -1
			}

		} else if flag < 0 {
			if tree.rights[v] >= 0 {
				v = tree.rights[v]
			} else {
				return -1
			}

		} else {
			return v
		}
	}
	return -1
}

func (tree *AVLTree) Remove(val interface{}) {
	if u := tree.Get(val); u >= 0 {
		tree.remove(u)
	}
}

func (tree *AVLTree) remove(u int) {
	if right := tree.rights[u]; right >= 0 {
		tree.Vals[u] = tree.Vals[right]
		tree.remove(right)

	} else if left := tree.lefts[u]; left >= 0 {
		tree.Vals[u] = tree.Vals[left]
		tree.remove(left)

	} else {
		p := tree.pars[u]
		if p < 0 {
			tree.Root = -1
		} else {
			if tree.lefts[p] == u {
				tree.lefts[p] = -1
			} else {
				tree.rights[p] = -1
			}
		}
		tree.delNode(u)
		return
	}

	tree.fixHeight(u)
	tree.balance(u)
}

func (tree *AVLTree) Next(u int) int {
	if right := tree.rights[u]; right >= 0 {
		u = right
		for tree.lefts[u] >= 0 {
			u = tree.lefts[u]
		}
		return u
	}

	if p := tree.pars[u]; p >= 0 {
		if tree.lefts[p] == u {
			return p
		} else {
			for p = tree.pars[u]; p >= 0 && tree.rights[p] == u; u, p = p, tree.pars[p] {
			}
			return p
		}
	} else {
		return -1
	}
}

func (tree *AVLTree) Prev(u int) int {
	if left := tree.lefts[u]; left >= 0 {
		u = left
		for tree.rights[u] >= 0 {
			u = tree.lefts[u]
		}
		return u
	}

	if p := tree.pars[u]; p >= 0 {
		if tree.rights[p] == u {
			return p
		} else {
			for p = tree.pars[u]; p >= 0 && tree.lefts[p] == u; u, p = p, tree.pars[p] {
			}
			return p
		}

	} else {
		return -1
	}
}

func (tree *AVLTree) LowerBound(val interface{}) int {
	u := tree.Root
	for u >= 0 {
		flag := tree.cmp(tree.Vals[u], val)
		if flag < 0 {
			u = tree.rights[u]
		} else if flag > 0 {
			if tree.lefts[u] <= 0 {
				return u
			}
			u = tree.lefts[u]
		} else {
			return u
		}
	}
	return -1
}

func (tree *AVLTree) UpperBound(val interface{}) int {
	u := tree.LowerBound(val)
	if u < 0 {
		return u
	}

	if tree.cmp(tree.Vals[u], val) == 0 {
		u = tree.Next(u)
	}
	return u
}

func (tree *AVLTree) First() int {
	u := tree.Root
	for u >= 0 && tree.lefts[u] >= 0 {
		u = tree.lefts[u]
	}
	return u
}

func (tree *AVLTree) Last() int {
	u := tree.Root
	for u >= 0 && tree.rights[u] >= 0 {
		u = tree.rights[u]
	}
	return u
}

func (tree *AVLTree) balance(u int) {
	if u < 0 {
		return
	}
	leftH, rightH := 0, 0
	left, right := tree.lefts[u], tree.rights[u]
	if left >= 0 {
		leftH = tree.heights[left]
	}
	if right >= 0 {
		rightH = tree.heights[right]
	}

	if leftH-rightH > 1 {
		leftLeft, leftRight := tree.lefts[left], tree.rights[left]
		if ((leftLeft >= 0 && leftRight >= 0) && tree.heights[leftLeft] > tree.heights[leftRight]) || leftRight < 0 { // LL
			tree.rotateRight(u)

		} else if ((leftLeft >= 0 && leftRight >= 0) && tree.heights[leftLeft] < tree.heights[leftRight]) || leftLeft < 0 { // LR
			tree.rotateLeft(left)
			tree.rotateRight(u)
		}

	} else if rightH-leftH > 1 {
		rightLeft, rightRight := tree.lefts[right], tree.rights[right]
		if ((rightLeft >= 0 && rightRight >= 0) && tree.heights[rightRight] > tree.heights[rightLeft]) || rightLeft < 0 { // RR
			tree.rotateLeft(u)

		} else if ((rightLeft >= 0 && rightRight >= 0) && tree.heights[rightRight] < tree.heights[rightLeft]) || rightRight < 0 { // RL
			tree.rotateRight(right)
			tree.rotateLeft(u)
		}
	}
}

func (tree *AVLTree) fixChildPar(u int) {
	if u < 0 {
		return
	}
	left, right := tree.lefts[u], tree.rights[u]
	if left >= 0 {
		tree.pars[left] = u
	}
	if right >= 0 {
		tree.pars[right] = u
	}
}

func (tree *AVLTree) fixHeight(u int) {
	if u < 0 {
		return
	}
	leftH, rightH := 0, 0
	left, right := tree.lefts[u], tree.rights[u]
	if left >= 0 {
		leftH = tree.heights[left]
	}
	if right >= 0 {
		rightH = tree.heights[right]
	}

	tree.heights[u] = max(leftH, rightH) + 1
}

func (tree *AVLTree) rotateLeft(u int) {
	p, v := tree.pars[u], tree.rights[u]
	tree.rights[u] = tree.lefts[v]
	tree.lefts[v] = u

	if p >= 0 {
		if tree.lefts[p] == u {
			tree.lefts[p] = v
		} else {
			tree.rights[p] = v
		}
	} else {
		tree.Root = v
		tree.pars[v] = -1
	}
	tree.fixChildPar(u)
	tree.fixChildPar(v)
	tree.fixChildPar(p)
	tree.fixHeight(u)
	tree.fixHeight(v)
	tree.fixHeight(p)
}

func (tree *AVLTree) rotateRight(u int) {
	p, v := tree.pars[u], tree.lefts[u]
	tree.lefts[u] = tree.rights[v]
	tree.rights[v] = u
	if p >= 0 {
		if tree.lefts[p] == u {
			tree.lefts[p] = v
		} else {
			tree.rights[p] = v
		}
	} else {
		tree.Root = v
		tree.pars[v] = -1
	}
	tree.fixChildPar(u)
	tree.fixChildPar(v)
	tree.fixChildPar(p)
	tree.fixHeight(u)
	tree.fixHeight(v)
	tree.fixHeight(p)
}

func (tree *AVLTree) newNode(val interface{}) int {
	u := tree.N
	if cd := len(tree.dels); cd > 0 {
		u = tree.dels[cd-1]
		tree.Vals[u] = val
		tree.dels = tree.dels[:cd-1]
		tree.lefts[u] = -1
		tree.rights[u] = -1
		tree.pars[u] = -1
		tree.heights[u] = 1
		tree.N++

	} else {
		tree.Vals = append(tree.Vals, val)
		tree.lefts = append(tree.lefts, -1)
		tree.rights = append(tree.rights, -1)
		tree.pars = append(tree.pars, -1)
		tree.heights = append(tree.heights, 1)
		tree.N++
	}

	return u
}

func (tree *AVLTree) delNode(u int) {
	tree.dels = append(tree.dels, u)
	tree.Vals[u] = nil
	tree.pars[u] = -1
	tree.N--
}

func (tree *AVLTree) print(u int) {
	fmt.Print("[")
	defer fmt.Print("]")

	if u < 0 {
		fmt.Print("nil")
		return
	}
	tree.print(tree.lefts[u])
	fmt.Print("-", tree.Vals[u], "(", tree.pars[u], ".", u, ")", "-")
	tree.print(tree.rights[u])
}

///////////////////////////

type Pair struct {
	val int
	cnt int
}

func maxResult(nums []int, k int) int {
	mp := NewAVLTree(func(a, b interface{}) int { return a.(*Pair).val - b.(*Pair).val })
	mpc := 0

	add := func(val int) {
		p := &Pair{
			val: val,
			cnt: 0,
		}

		it := mp.Get(p)

		if it < 0 {
			it = mp.Put(p)
		}

		mp.Vals[it].(*Pair).cnt++
		mpc++
	}

	del := func(val int) {
		p := &Pair{
			val: val,
			cnt: 0,
		}

		it := mp.Get(p)
		if it < 0 {
			return
		}

		p = mp.Vals[it].(*Pair)
		p.cnt--
		if p.cnt == 0 {
			mp.Remove(p)
		}
		mpc--
	}

	n := len(nums)
	dp := make([]int, n)

	for i := n - 1; i >= 0; i-- {

		dp[i] = nums[i]
		if mpc > 0 {
			mx := mp.Vals[mp.Last()].(*Pair).val
			dp[i] = nums[i] + mx
		}

		add(dp[i])

		if mpc > k {

			del(dp[i+k])
		}
	}

	//fmt.Println(dp)

	return dp[0]
}

func main() {
	nums := []int{1, -1, -2, 4, -7, 3}
	k := 2

	fmt.Println(maxResult(nums, k))

}
