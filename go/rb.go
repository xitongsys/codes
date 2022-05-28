
package rb

type RBKeyType interface {
	Less(RBKeyType) bool
	Equal(RBKeyType) bool
}

type RBValueType interface {}


type RBTree struct {
	root *RBNode
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
	rbt.root.Remove(key)
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

func (rbnode *RBNode) Remove(key RBKeyType) {
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
	if ps == nil || ps.color == 1 {
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


