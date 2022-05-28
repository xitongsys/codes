package main

import "fmt"

type RBKeyType interface {
	Less(RBKeyType) bool
	Equal(RBKeyType) bool
}

type RBValueType interface {}


type RBTree struct {
	Root *RBNode
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (rbt *RBTree) Add(key RBKeyType, value RBValueType) {
	if rbt.Root == nil {
		rbt.Root = NewRBNode(key, value, 0)
		return
	}
	rbt.Root.Add(key, value)

	for rbt.Root.Parent != nil {
		rbt.Root = rbt.Root.Parent
	}
}

func (rbt *RBTree) Remove(key RBKeyType) {
	if rbt.Root == nil {
		return
	}

	if rbt.Root.Left == nil && rbt.Root.Right == nil && rbt.Root.Key.Equal(key) {
		rbt.Root = nil
		return
	}

	if node := rbt.Root.Find(key); node != nil {
		node.Remove()
	}
}

func (rbt *RBTree) LowerBound(key RBKeyType) *RBNode {
	if rbt.Root == nil {
		return nil
	}
	return rbt.Root.LowerBound(key)
}

func (rbt *RBTree) Check() int {
	if rbt.Root == nil {
		return 0
	}
	return rbt.Root.Check()
}


func (rbt *RBTree) String() string {
	if rbt.Root == nil {
		return "nil"
	}
	return rbt.Root.String()
}

func (rbt *RBTree) Find(key RBKeyType) *RBNode {
	if rbt.Root == nil {
		return nil
	}
	return rbt.Root.Find(key)
}


type RBNode struct {
	Left, Right, Parent *RBNode
	Color int //0: black, 1: red
	Key RBKeyType
	Value RBValueType
}

func NewRBNode(key RBKeyType, value RBValueType, color int) *RBNode {
	return &RBNode{
		Color: color,
		Key: key,
		Value: value,
	}
}


func (rbnode *RBNode) Add(key RBKeyType, value RBValueType) {
	if rbnode.Key.Equal(key) {
		rbnode.Value = value

	} else if rbnode.Key.Less(key) {
		if rbnode.Right == nil {
			rbnode.Right = &RBNode {
				Color: 1,
				Key: key,
				Value: value,
				Parent: rbnode,
			}
			rbnode.Right.Adjust()
		}else{
			rbnode.Right.Add(key, value)
		}

	}else{
		if rbnode.Left == nil {
			rbnode.Left = &RBNode {
				Color: 1,
				Key: key,
				Value: value,
				Parent: rbnode,
			}
			rbnode.Left.Adjust()

		}else{
			rbnode.Left.Add(key, value)
		}
	}
}

func (rbnode *RBNode) remove() {
	p := rbnode.Parent
	if p != nil {
		if p.Left == rbnode {
			p.Left = nil
		}else{
			p.Right = nil
		}
	}
}

func (rbnode *RBNode) Remove() {
	p := rbnode.Parent
	if n := rbnode.Next(); n != nil {
		rbnode.Key = n.Key
		rbnode.Value = n.Value
		n.Remove()

	}else{
		if rbnode.Color == 1 { // red leaf node 
			if p.Left == rbnode { 
				p.Left = nil

			}else{ 
				p.Right = nil
			}

		} else {
			if rbnode.Left != nil { // black with one left red child
				rbnode.Key = rbnode.Left.Key
				rbnode.Value = rbnode.Left.Value
				rbnode.Left = nil

			} else { // black leaf node
				s := p.Sibling()

				if p.Color == 1 { // parent red
					if s.Left == nil && s.Right == nil {
						p.Color = 0
						s.Color = 1
						rbnode.remove()

					} else if s.Left != nil && s.Right == nil {
						rbnode.remove()
						s.RotateRight()
						s.Parent.Color = 0
						s.Color = 1
						p.RotateLeft()

					} else if s.Left == nil && s.Right != nil {
						rbnode.remove()
						p.RotateLeft()

					} else {
						rbnode.remove()
						p.RotateLeft()
						p = p.Parent
						p.Color = 1
						p.Left.Color = 0
						p.Right.Color = 1
					}

				} else { // parent black
					if s.Color == 1 { // sibling red
						rbnode.remove()
						s.Color = 0
						s.Left.Color = 1
						p.RotateLeft()

					} else { // sibling black
						rbnode.remove()
						if s.Left == nil && s.Right != nil {
							s.Right.Color = 0
							p.RotateLeft()

						} else if s.Left != nil && s.Right == nil {
							s.Color = 1
							s.Left.Color = 0
							s.RotateRight()

						} else if s.Left != nil && s.Right != nil {
							s.Right.Color = 0
							p.RotateLeft()

						} else {
							s.Color = 1
						}
					}
				}
			}
		}
	}
}

func (rbnode *RBNode) LowerBound(key RBKeyType) *RBNode {
	if rbnode.Key.Equal(key) {
		return rbnode
	}

	if rbnode.Key.Less(key) {
		if rbnode.Right == nil {
			return nil
		}
		return rbnode.Right.LowerBound(key)
	}


	if rbnode.Left == nil {
		return rbnode
	}

	node := rbnode.Left.LowerBound(key)
	if node == nil {
		node = rbnode
	}

	return node
}

func (rbnode *RBNode) Find(key RBKeyType) *RBNode {
	node := rbnode.LowerBound(key)
	if node != nil && ! node.Key.Equal(key){
		node = nil
	}
	return node
}

func (rbnode *RBNode) Next() *RBNode {
	if rbnode.Right != nil {
		node := rbnode.Right
		for node.Left != nil {
			node = node.Left
		}
		return node

	} else if rbnode.Parent != nil {
		return rbnode.Parent
	}
	return nil
}

func (rbnode *RBNode) Prev() *RBNode {
	if rbnode.Left != nil {
		node := rbnode.Left;
		for node.Right != nil {
			node = node.Right
		}
		return node
	}
	return nil
}

func (rbnode *RBNode) Sibling() *RBNode {
	if rbnode.Parent == nil {
		return nil
	}

	if rbnode == rbnode.Parent.Left {
		return rbnode.Parent.Right
	}
	return rbnode.Parent.Left
}

func (rbnode *RBNode) Adjust() {
	if rbnode.Parent == nil {
		rbnode.Color = 0
		return
	}

	p, pp := rbnode.Parent, rbnode.Parent.Parent
	if p.Color == 0 {
		return
	}

	ps := p.Sibling()
	if ps == nil || ps.Color == 0 { // parent red, parent-sibling black
		if p == pp.Left && rbnode == p.Left { // LL
			p.Color = 0
			pp.Color = 1
			pp.RotateRight()

		} else if p == pp.Left && rbnode == p.Right { // LR -> LL
			p.RotateLeft()
			rbnode.Adjust()

		} else if p == pp.Right && rbnode == p.Right { // RR
			p.Color = 0
			pp.Color = 1
			pp.RotateLeft()

		} else if p == pp.Right && rbnode == p.Left { // RL -> RR
			p.RotateRight()
			rbnode.Adjust()
		}

	} else { // parent red, parent-sibling red
		p.Color = 0
		ps.Color = 0
		pp.Color = 1
		pp.Adjust()
	}
}

func (rbnode *RBNode) RotateLeft() {
	p, right := rbnode.Parent, rbnode.Right
	rbnode.Right = right.Left
	rbnode.Parent = right 
	right.Left = rbnode
	right.Parent = p
	if p != nil {
		if p.Left == rbnode {
			p.Left = right
		} else {
			p.Right = right
		}
	}
}

func (rbnode *RBNode) RotateRight() {
	p, left := rbnode.Parent, rbnode.Left
	rbnode.Left = left.Right
	rbnode.Parent = left
	left.Right = rbnode
	left.Parent = p
	if p != nil {
		if p.Left == rbnode {
			p.Left = left
		} else {
			p.Right = left
		}
	}
}

func (rbnode *RBNode) Check() int {
	hl, hr := 0, 0
	if rbnode.Left != nil {
		hl = rbnode.Left.Check()
		if rbnode.Color == 1 && rbnode.Left.Color == 1 {
			return -1
		}
	}

	if rbnode.Right != nil {
		hr = rbnode.Right.Check()
		if rbnode.Color == 1 && rbnode.Right.Color == 1 {
			return -1
		}
	}

	if hl != hr || hl < 0 || hr < 0 {
		return -1
	}

	h := hl
	if rbnode.Color == 0 {
		h++
	}

	return h
}

func (rbnode *RBNode) String() string {
	res := ""
	if rbnode.Left != nil {
		res = rbnode.Left.String()
	}
	res += " " + fmt.Sprint(rbnode.Key)
	if rbnode.Right != nil {
		res += rbnode.Right.String()
	}
	return res
}

