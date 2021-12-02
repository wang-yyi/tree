package binaryTree

type BinaryTree struct {
	root *Node
}

type Node struct {
	Key   string
	Value interface{}
	left  *Node
	right *Node
}

func (t *BinaryTree) Put(key string, value interface{}) {
	n := &Node{Key: key, Value: value}
	if t.root == nil {
		t.root = n
	} else {
		t.putRecursive(t.root, n)
	}
}

func (t *BinaryTree) putRecursive(node, newNode *Node) {
	if newNode.Key < node.Key {
		if node.left != nil {
			t.putRecursive(node.left, newNode)
		} else {
			node.left = newNode
		}
	} else if newNode.Key > node.Key {
		if node.right != nil {
			t.putRecursive(node.right, newNode)
		} else {
			node.right = newNode
		}
	} else {
		node.Value = newNode.Value
	}
}

func (t *BinaryTree) Get(key string) interface{} {
	return t.getRecursive(t.root, key)
}

func (t *BinaryTree) getRecursive(node *Node, key string) interface{} {
	if node == nil {
		return nil
	}

	if key < node.Key {
		return t.getRecursive(node.left, key)
	}

	if key > node.Key {
		return t.getRecursive(node.right, key)
	}

	return node.Value
}
