package binaryTree

// BinaryTree 二叉树
type BinaryTree struct {
	Root *Node
}

type Node struct {
	Key   int
	Value interface{}
	Left  *Node
	Right *Node
}

// Add  添加数据
func (t *BinaryTree) Add(key int, value interface{}) {
	n := &Node{Key: key, Value: value}
	if t.Root == nil {
		t.Root = n
	} else {
		t.addRecursive(t.Root, n)
	}
}

func (t *BinaryTree) addRecursive(node, newNode *Node) {
	if newNode.Key < node.Key {
		if node.Left != nil {
			t.addRecursive(node.Left, newNode)
		} else {
			node.Left = newNode
		}
	} else if newNode.Key > node.Key {
		if node.Right != nil {
			t.addRecursive(node.Right, newNode)
		} else {
			node.Right = newNode
		}
	} else {
		node.Value = newNode.Value
	}
}

// Get 查询数据
func (t *BinaryTree) Get(key int) interface{} {
	return t.getRecursive(t.Root, key)
}

func (t *BinaryTree) getRecursive(node *Node, key int) interface{} {
	if node == nil {
		return nil
	}

	if key < node.Key {
		return t.getRecursive(node.Left, key)
	}

	if key > node.Key {
		return t.getRecursive(node.Right, key)
	}

	return node.Value
}

// traverseInOrderKey 中序遍历所有key
func (n *Node) traverseInOrderKey(resp *[]int) {
	if n == nil {
		return
	}

	n.Left.traverseInOrderKey(resp)
	*resp = append(*resp, n.Key)
	n.Right.traverseInOrderKey(resp)
}
