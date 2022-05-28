package main

type RBKeyType interface {
	Less(RBKeyType) bool
	Equal(RBKeyType) bool
}

type RBValueType interface {}


type RBTree struct {
	root *RBNode
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (rbt *RBTree) Add(key RBKeyType, value RBValueType) {
	if rbt.root == nil {
		rbt.root = NewRBNode(key, value, 0)
		return
	}
	rbt.root.Add(key, value)

}

func (rbt *RBTree) Remove(key RBKeyType) {
	if rbt.root == nil {
		return
	}

	if rbt.root.left == nil && rbt.root.right == nil && rbt.root.key.Equal(key) {
		rbt.root = nil
		return
	}

	if node := rbt.root.Find(key); node != nil {
		node.Remove()
	}
}

func (rbt *RBTree) LowerBound(key RBKeyType) *RBNode {
	if rbt.root == nil {
		return nil
	}
	return rbt.root.LowerBound(key)
}

func (rbt *RBTree) Find(key RBKeyType) *RBNode {
	if rbt.root == nil {
		return nil
	}
	return rbt.root.Find(key)
}


type RBNode struct {
	left, right, parent *RBNode
	color int //0: black, 1: red
	key RBKeyType
	value RBValueType
}

func NewRBNode(key RBKeyType, value RBValueType, color int) *RBNode {
	return &RBNode{
		color: color,
		key: key,
		value: value,
	}
}


func (rbnode *RBNode) Add(key RBKeyType, value RBValueType) {
	if rbnode.key.Equal(key) {
		rbnode.value = value

	} else if rbnode.key.Less(key) {
		if rbnode.right == nil {
			rbnode.right = &RBNode {
				color: 1,
				key: key,
				value: value,
				parent: rbnode,
			}
			rbnode.right.Adjust()
		}else{
			rbnode.right.Add(key, value)
		}

	}else{
		if rbnode.left == nil {
			rbnode.left = &RBNode {
				color: 1,
				key: key,
				value: value,
				parent: rbnode,
			}
			rbnode.left.Adjust()

		}else{
			rbnode.left.Add(key, value)
		}
	}
}

func (rbnode *RBNode) remove() {
	p := rbnode.parent
	if p != nil {
		if p.left == rbnode {
			p.left = nil
		}else{
			p.right = nil
		}
	}
}

func (rbnode *RBNode) Remove() {
	p := rbnode.parent
	if n := rbnode.Next(); n != nil {
		rbnode.key = n.key
		rbnode.value = n.value
		n.Remove()

	}else{
		if rbnode.color == 1 { // red leaf node 
			if p.left == rbnode { 
				p.left = nil

			}else{ 
				p.right = nil
			}

		} else {
			if rbnode.left != nil { // black with one left red child
				rbnode.key = rbnode.left.key
				rbnode.value = rbnode.left.value
				rbnode.left = nil

			} else { // black leaf node
				s := p.Sibling()

				if p.color == 1 { // parent red
					if s.left == nil && s.right == nil {
						p.color = 0
						s.color = 1
						rbnode.remove()

					} else if s.left != nil && s.right == nil {
						rbnode.remove()
						s.RotateRight()
						s.parent.color = 0
						s.color = 1
						p.RotateLeft()

					} else if s.left == nil && s.right != nil {
						rbnode.remove()
						p.RotateLeft()

					} else {
						rbnode.remove()
						p.RotateLeft()
						p = p.parent
						p.color = 1
						p.left.color = 0
						p.right.color = 1
					}

				} else { // parent black
					if s.color == 1 { // sibling red
						rbnode.remove()
						s.color = 0
						s.left.color = 1
						p.RotateLeft()

					} else { // sibling black
						rbnode.remove()
						if s.left == nil && s.right != nil {
							s.right.color = 0
							p.RotateLeft()

						} else if s.left != nil && s.right == nil {
							s.color = 1
							s.left.color = 0
							s.RotateRight()

						} else if s.left != nil && s.right != nil {
							s.right.color = 0
							p.RotateLeft()

						} else {
							s.color = 1
						}
					}
				}
			}
		}
	}
}

func (rbnode *RBNode) LowerBound(key RBKeyType) *RBNode {
	if rbnode.key.Equal(key) {
		return rbnode
	}

	if rbnode.key.Less(key) {
		if rbnode.right == nil {
			return nil
		}
		return rbnode.right.LowerBound(key)
	}


	if rbnode.left == nil {
		return rbnode
	}

	node := rbnode.left.LowerBound(key)
	if node == nil {
		node = rbnode
	}

	return node
}

func (rbnode *RBNode) Find(key RBKeyType) *RBNode {
	node := rbnode.LowerBound(key)
	if node != nil && ! node.key.Equal(key){
		node = nil
	}
	return node
}

func (rbnode *RBNode) Next() *RBNode {
	if rbnode.right != nil {
		node := rbnode.right
		for node.left != nil {
			node = node.left
		}
		return node

	} else if rbnode.parent != nil {
		return rbnode.parent
	}
	return nil
}

func (rbnode *RBNode) Prev() *RBNode {
	if rbnode.left != nil {
		node := rbnode.left;
		for node.right != nil {
			node = node.right
		}
		return node
	}
	return nil
}

func (rbnode *RBNode) Sibling() *RBNode {
	if rbnode.parent == nil {
		return nil
	}

	if rbnode == rbnode.parent.left {
		return rbnode.parent.right
	}
	return rbnode.parent.left
}


func (rbnode *RBNode) Adjust() {
	if rbnode.parent == nil {
		rbnode.color = 0
		return
	}

	p, pp := rbnode.parent, rbnode.parent.parent
	if p.color == 0 {
		return
	}

	ps := p.Sibling()
	if ps == nil || ps.color == 0 { // parent red, parent-sibling black
		if p == pp.left && rbnode == p.left { // LL
			p.color = 0
			pp.color = 1
			pp.RotateRight()

		} else if p == pp.left && rbnode == p.right { // LR -> LL
			p.RotateLeft()
			rbnode.Adjust()

		} else if p == pp.right && rbnode == p.right { // RR
			p.color = 0
			pp.color = 1
			pp.RotateLeft()

		} else if p == pp.right && rbnode == p.left { // RL -> RR
			p.RotateRight()
			rbnode.Adjust()
		}

	} else { // parent red, parent-sibling red
		p.color = 0
		ps.color = 0
		pp.color = 1
		pp.Adjust()
	}
}

func (rbnode *RBNode) RotateLeft() {
	p, right := rbnode.parent, rbnode.right
	rbnode.right = right.left
	rbnode.parent = right 
	right.left = rbnode
	right.parent = p
}

func (rbnode *RBNode) RotateRight() {
	p, left := rbnode.parent, rbnode.left
	rbnode.left = left.right
	rbnode.parent = left
	left.right = rbnode
	left.parent = p
}


