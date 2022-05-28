package main

import (
	"fmt"
)

type Comparator func(a, b interface{}) int

type color bool
const (
	black, red color = true, false
)

type RBTree struct {
	Root       *RBNode
	size       int
	Comparator Comparator
}

type RBNode struct {
	Key    interface{}
	Value  interface{}
	color  color
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
}

func NewRBTree(comparator Comparator) *RBTree {
	return &RBTree{Comparator: comparator}
}

func (tree *RBTree) Put(key interface{}, value interface{}) {
	var insertedRBNode *RBNode
	if tree.Root == nil {
		tree.Comparator(key, key)
		tree.Root = &RBNode{Key: key, Value: value, color: red}
		insertedRBNode = tree.Root
	} else {
		node := tree.Root
		loop := true
		for loop {
			compare := tree.Comparator(key, node.Key)
			switch {
			case compare == 0:
				node.Key = key
				node.Value = value
				return
			case compare < 0:
				if node.Left == nil {
					node.Left = &RBNode{Key: key, Value: value, color: red}
					insertedRBNode = node.Left
					loop = false
				} else {
					node = node.Left
				}
			case compare > 0:
				if node.Right == nil {
					node.Right = &RBNode{Key: key, Value: value, color: red}
					insertedRBNode = node.Right
					loop = false
				} else {
					node = node.Right
				}
			}
		}
		insertedRBNode.Parent = node
	}
	tree.insertCase1(insertedRBNode)
	tree.size++
}

func (tree *RBTree) Begin() *RBNode {
	if tree.Root == nil {
		return nil
	}

	node := tree.Root 
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (tree *RBTree) Last() *RBNode {
	if tree.Root == nil {
		return nil
	}

	node := tree.Root 
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (tree *RBTree) Get(key interface{}) (value interface{}, found bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.Value, true
	}
	return nil, false
}

func (tree *RBTree) GetRBNode(key interface{}) *RBNode {
	return tree.lookup(key)
}

func (tree *RBTree) Remove(key interface{}) {
	var child *RBNode
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumRBNode()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.color == black {
			node.color = nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceRBNode(node, child)
		if node.Parent == nil && child != nil {
			child.color = black
		}
	}
	tree.size--
}

func (tree *RBTree) Empty() bool {
	return tree.size == 0
}

func (tree *RBTree) Size() int {
	return tree.size
}

func (node *RBNode) Size() int {
	if node == nil {
		return 0
	}
	size := 1
	if node.Left != nil {
		size += node.Left.Size()
	}
	if node.Right != nil {
		size += node.Right.Size()
	}
	return size
}

func (tree *RBTree) Keys() []interface{} {
	keys := []interface{}{}
	var dfs func(node *RBNode)
	dfs = func(node *RBNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		keys = append(keys, node.Key)
		dfs(node.Right)
	}
	
	dfs(tree.Root)
	return keys
}

func (tree *RBTree) Values() []interface{} {
	values := []interface{}{}
	var dfs func(node *RBNode)
	dfs = func(node *RBNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		values = append(values, node.Value)
		dfs(node.Right)
	}
	
	dfs(tree.Root)
	return values
}

func (tree *RBTree) Left() *RBNode {
	var parent *RBNode
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Left
	}
	return parent
}

func (tree *RBTree) Right() *RBNode {
	var parent *RBNode
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Right
	}
	return parent
}

func (tree *RBTree) Floor(key interface{}) (floor *RBNode, found bool) {
	found = false
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			node = node.Left
		case compare > 0:
			floor, found = node, true
			node = node.Right
		}
	}
	if found {
		return floor, true
	}
	return nil, false
}

func (tree *RBTree) Ceiling(key interface{}) (ceiling *RBNode, found bool) {
	found = false
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			ceiling, found = node, true
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	if found {
		return ceiling, true
	}
	return nil, false
}

func (tree *RBTree) Clear() {
	tree.Root = nil
	tree.size = 0
}

func (tree *RBTree) String() string {
	str := "RedBlackRBTree\n"
	if !tree.Empty() {
		output(tree.Root, "", true, &str)
	}
	return str
}

func (node *RBNode) String() string {
	return fmt.Sprintf("%v", node.Key)
}

func output(node *RBNode, prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str)
	}
}

func (tree *RBTree) lookup(key interface{}) *RBNode {
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	return nil
}

func (node *RBNode) grandparent() *RBNode {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

func (node *RBNode) uncle() *RBNode {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *RBNode) sibling() *RBNode {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (tree *RBTree) rotateLeft(node *RBNode) {
	right := node.Right
	tree.replaceRBNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *RBTree) rotateRight(node *RBNode) {
	left := node.Left
	tree.replaceRBNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (tree *RBTree) replaceRBNode(old *RBNode, new *RBNode) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (tree *RBTree) insertCase1(node *RBNode) {
	if node.Parent == nil {
		node.color = black
	} else {
		tree.insertCase2(node)
	}
}

func (tree *RBTree) insertCase2(node *RBNode) {
	if nodeColor(node.Parent) == black {
		return
	}
	tree.insertCase3(node)
}

func (tree *RBTree) insertCase3(node *RBNode) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.Parent.color = black
		uncle.color = black
		node.grandparent().color = red
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *RBTree) insertCase4(node *RBNode) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	tree.insertCase5(node)
}

func (tree *RBTree) insertCase5(node *RBNode) {
	node.Parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		tree.rotateLeft(grandparent)
	}
}

func (node *RBNode) maximumRBNode() *RBNode {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (tree *RBTree) deleteCase1(node *RBNode) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *RBTree) deleteCase2(node *RBNode) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.Parent.color = red
		sibling.color = black
		if node == node.Parent.Left {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *RBTree) deleteCase3(node *RBNode) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		tree.deleteCase1(node.Parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *RBTree) deleteCase4(node *RBNode) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		node.Parent.color = black
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *RBTree) deleteCase5(node *RBNode) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == red &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		sibling.Left.color = black
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Right) == red &&
		nodeColor(sibling.Left) == black {
		sibling.color = red
		sibling.Right.color = black
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *RBTree) deleteCase6(node *RBNode) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.Parent)
	node.Parent.color = black
	if node == node.Parent.Left && nodeColor(sibling.Right) == red {
		sibling.Right.color = black
		tree.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == red {
		sibling.Left.color = black
		tree.rotateRight(node.Parent)
	}
}

func nodeColor(node *RBNode) color {
	if node == nil {
		return black
	}
	return node.color
}


