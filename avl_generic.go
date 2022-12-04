package main

import "fmt"

type AVLTree[TVal any] struct {
	N, Root int
	Vals    []TVal

	heights             []int
	lefts, rights, pars []int
	dels                []int
	cmp                 func(a, b TVal) int
}

func NewAVLTree[TVal any](cmp func(a, b TVal) int) *AVLTree[TVal] {
	tree := &AVLTree[TVal]{
		Root: -1,
		cmp:  cmp,
	}
	return tree
}

func (tree *AVLTree[TVal]) Put(val TVal) int {
	u := tree.Root
	if u < 0 { // empty tree
		u = tree.newNode(val)
		tree.Root = u
		return u
	}

	return tree.put(u, val)
}

func (tree *AVLTree[TVal]) put(u int, val TVal) int {
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

func (tree *AVLTree[TVal]) Size() int {
	return tree.N
}

func (tree *AVLTree[TVal]) Get(val TVal) int {
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

func (tree *AVLTree[TVal]) Remove(val TVal) {
	if u := tree.Get(val); u >= 0 {
		v := u
		if left := tree.lefts[u]; left >= 0 {
			v = left
			for tree.rights[v] >= 0 {
				v = tree.rights[v]
			}

		} else if right := tree.rights[u]; right >= 0 {
			v = right
			for tree.lefts[v] >= 0 {
				v = tree.lefts[v]
			}
		}

		tree.Vals[u] = tree.Vals[v]

		if p := tree.pars[v]; p >= 0 {
			if tree.lefts[p] == v {
				tree.lefts[p] = -1
			} else {
				tree.rights[p] = -1
			}

			for p >= 0 {
				tree.fixHeight(p)
				tree.balance(p)
				p = tree.pars[p]
			}

		} else {
			tree.Root = -1
		}

		tree.delNode(v)
	}
}

func (tree *AVLTree[TVal]) Next(u int) int {
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

func (tree *AVLTree[TVal]) Prev(u int) int {
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

func (tree *AVLTree[TVal]) LowerBound(val TVal) int {
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

func (tree *AVLTree[TVal]) UpperBound(val TVal) int {
	u := tree.LowerBound(val)
	if u < 0 {
		return u
	}

	if tree.cmp(tree.Vals[u], val) == 0 {
		u = tree.Next(u)
	}
	return u
}

func (tree *AVLTree[TVal]) First() int {
	u := tree.Root
	for u >= 0 && tree.lefts[u] >= 0 {
		u = tree.lefts[u]
	}
	return u
}

func (tree *AVLTree[TVal]) Last() int {
	u := tree.Root
	for u >= 0 && tree.rights[u] >= 0 {
		u = tree.rights[u]
	}
	return u
}

func (tree *AVLTree[TVal]) balance(u int) {
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

func (tree *AVLTree[TVal]) fixChildPar(u int) {
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

func (tree *AVLTree[TVal]) fixHeight(u int) {
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

	if leftH > rightH {
		tree.heights[u] = leftH + 1
	} else {
		tree.heights[u] = rightH + 1
	}
}

func (tree *AVLTree[TVal]) rotateLeft(u int) {
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

func (tree *AVLTree[TVal]) rotateRight(u int) {
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

func (tree *AVLTree[TVal]) newNode(val TVal) int {
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

func (tree *AVLTree[TVal]) delNode(u int) {
	tree.dels = append(tree.dels, u)
	tree.pars[u] = -1
	tree.N--
}

func (tree *AVLTree[TVal]) print(u int) {
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

type KeyValue[TKey any, TValue any] struct {
	key   TKey
	value TValue
}

type TreeMap[TKey any, TValue any] struct {
	avlTree *AVLTree[*KeyValue[TKey, TValue]]
}

func NewTreeMap[TKey any, TValue any](cmp func(a, b TKey) int) *TreeMap[TKey, TValue] {

	mp := &TreeMap[TKey, TValue]{
		avlTree: NewAVLTree[*KeyValue[TKey, TValue]](func(a, b *KeyValue[TKey, TValue]) int {
			return cmp(a.key, b.key)
		}),
	}
	return mp
}

func (treemap *TreeMap[TKey, TValue]) Size() int {
	return treemap.avlTree.Size()
}

func (treemap *TreeMap[TKey, TValue]) Get(key TKey) (value TValue, err error) {
	kv := &KeyValue[TKey, TValue]{
		key: key,
	}

	u := treemap.avlTree.Get(kv)
	if u < 0 {
		err = fmt.Errorf("key not found")
	} else {
		value = treemap.avlTree.Vals[u].value
	}
	return
}

func (treemap *TreeMap[TKey, TValue]) Put(key TKey, value TValue) {
	kv := &KeyValue[TKey, TValue]{
		key:   key,
		value: value,
	}
	treemap.avlTree.Put(kv)
}

func (treemap *TreeMap[TKey, TValue]) Remove(key TKey) {
	kv := &KeyValue[TKey, TValue]{
		key: key,
	}
	treemap.avlTree.Remove(kv)
}

func (treemap *TreeMap[TKey, TValue]) Has(key TKey) bool {
	kv := &KeyValue[TKey, TValue]{
		key: key,
	}

	u := treemap.avlTree.Get(kv)
	return u >= 0
}

func (treemap *TreeMap[TKey, TValue]) Find(key TKey) int {
	kv := &KeyValue[TKey, TValue]{
		key: key,
	}

	u := treemap.avlTree.Get(kv)
	return u
}

func (treemap *TreeMap[TKey, TValue]) GetKeyValue(iterator int) (key TKey, value TValue, err error) {
	if iterator < 0 || iterator >= len(treemap.avlTree.Vals) {
		err = fmt.Errorf("iterator out of bound")
		return
	}

	kv := treemap.avlTree.Vals[iterator]
	key = kv.key
	value = kv.value
	return
}

func (treemap *TreeMap[TKey, TValue]) GetKeyValuePointer(iterator int) *KeyValue[TKey, TValue] {
	if iterator < 0 || iterator >= len(treemap.avlTree.Vals) {
		return nil
	}

	kv := treemap.avlTree.Vals[iterator]
	return kv
}

func (treemap *TreeMap[TKey, TValue]) Next(iterator int) int {
	return treemap.avlTree.Next(iterator)
}

func (treemap *TreeMap[TKey, TValue]) Prev(iterator int) int {
	return treemap.avlTree.Prev(iterator)
}

func (treemap *TreeMap[TKey, TValue]) LowerBound(key TKey) int {
	kv := &KeyValue[TKey, TValue]{
		key: key,
	}
	return treemap.avlTree.LowerBound(kv)
}

func (treemap *TreeMap[TKey, TValue]) UpperBound(key TKey) int {
	kv := &KeyValue[TKey, TValue]{
		key: key,
	}
	return treemap.avlTree.UpperBound(kv)
}

func (treemap *TreeMap[TKey, TValue]) First() int {
	return treemap.avlTree.First()
}

func (treemap *TreeMap[TKey, TValue]) Last() int {
	return treemap.avlTree.Last()
}

////////////////////////////////

type TreeSet[TKey any] struct {
	avlTree *AVLTree[TKey]
}

func NewTreeSet[TKey any](cmp func(a, b TKey) int) *TreeSet[TKey] {
	mp := &TreeSet[TKey]{
		avlTree: NewAVLTree[TKey](cmp),
	}
	return mp
}

func (treeset *TreeSet[TKey]) Size() int {
	return treeset.avlTree.Size()
}

func (treeset *TreeSet[TKey]) Put(key TKey) {
	treeset.avlTree.Put(key)
}

func (treeset *TreeSet[TKey]) Remove(key TKey) {
	treeset.avlTree.Remove(key)
}

func (treeset *TreeSet[TKey]) Has(key TKey) bool {
	u := treeset.avlTree.Get(key)
	return u >= 0
}

func (treeset *TreeSet[TKey]) Find(key TKey) int {
	u := treeset.avlTree.Get(key)
	return u
}

func (treeset *TreeSet[TKey]) GetKey(iterator int) (key TKey, err error) {
	if iterator < 0 || iterator >= len(treeset.avlTree.Vals) {
		err = fmt.Errorf("iterator out of bound")
		return
	}

	key = treeset.avlTree.Vals[iterator]
	return
}

func (treeset *TreeSet[TKey]) Next(iterator int) int {
	return treeset.avlTree.Next(iterator)
}

func (treeset *TreeSet[TKey]) Prev(iterator int) int {
	return treeset.avlTree.Prev(iterator)
}

func (treeset *TreeSet[TKey]) LowerBound(key TKey) int {
	return treeset.avlTree.LowerBound(key)
}

func (treeset *TreeSet[TKey]) UpperBound(key TKey) int {
	return treeset.avlTree.UpperBound(key)
}

func (treeset *TreeSet[TKey]) First() int {
	return treeset.avlTree.First()
}

func (treeset *TreeSet[TKey]) Last() int {
	return treeset.avlTree.Last()
}

/////////////////////////////////

type MultiTreeSet[TKey any] struct {
	*TreeMap[TKey, int]
	N int
}

func NewMultiTreeSet[TKey any](cmp func(a, b TKey) int) *MultiTreeSet[TKey] {
	mp := &MultiTreeSet[TKey]{}
	mp.TreeMap = NewTreeMap[TKey, int](cmp)
	return mp
}

func (mtreeset *MultiTreeSet[TKey]) Size() int {
	return mtreeset.N
}

func (mtreeset *MultiTreeSet[TKey]) Put(key TKey) {
	if !mtreeset.Has(key) {
		mtreeset.TreeMap.Put(key, 0)
	}
	mtreeset.TreeMap.GetKeyValuePointer(mtreeset.TreeMap.Find(key)).value++
	mtreeset.N++
}

func (mtreeset *MultiTreeSet[TKey]) Remove(key TKey) {
	it := mtreeset.TreeMap.Find(key)
	if it < 0 {
		return
	}

	kv := mtreeset.TreeMap.GetKeyValuePointer(it)
	kv.value--
	if kv.value <= 0 {
		mtreeset.TreeMap.Remove(key)
	}
	mtreeset.N--
}

func (mtreeset *MultiTreeSet[TKey]) Has(key TKey) bool {
	return mtreeset.TreeMap.Has(key)
}

func (mtreeset *MultiTreeSet[TKey]) Count(key TKey) int {
	it := mtreeset.TreeMap.Find(key)
	if it < 0 {
		return 0
	}

	kv := mtreeset.TreeMap.GetKeyValuePointer(it)
	return kv.value
}

func (mtreeset *MultiTreeSet[TKey]) Find(key TKey) int {
	return mtreeset.TreeMap.Find(key)
}

func (mtreeset *MultiTreeSet[TKey]) LowerBound(key TKey) int {
	return mtreeset.TreeMap.LowerBound(key)
}

func (mtreeset *MultiTreeSet[TKey]) UpperBound(key TKey) int {
	return mtreeset.TreeMap.UpperBound(key)
}

func (mtreeset *MultiTreeSet[TKey]) First() int {
	return mtreeset.TreeMap.First()
}

func (mtreeset *MultiTreeSet[TKey]) Last() int {
	return mtreeset.TreeMap.Last()
}

func (mtreeset *MultiTreeSet[TKey]) GetKey(iterator int) (key TKey, count int, err error) {
	return mtreeset.GetKeyValue(iterator)
}

//////////////////////////

func maxResult(nums []int, k int) int {
	mst := NewMultiTreeSet[int](func(a, b int) int { return a - b })
	n := len(nums)
	dp := make([]int, n)

	for i := n - 1; i >= 0; i-- {

		dp[i] = nums[i]
		if mst.Size() > 0 {
			mx, _, _ := mst.GetKey(mst.Last())

			dp[i] = nums[i] + mx
		}

		mst.Put(dp[i])

		if mst.Size() > k {
			mst.Remove(dp[i+k])
		}

	}

	return dp[0]
}

func main() {
	nums := []int{1, -1, -2, 4, -7, 3}
	k := 2

	fmt.Println(maxResult(nums, k))
}
