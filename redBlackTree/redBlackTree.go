package redBlackTree

const (
	Red   = true
	Black = false
)

type RedBlackTree struct {
	Size int
	Root *Node
}

type Node struct {
	Key   int
	Value interface{}
	Color bool
	Left  *Node
	Right *Node
}

func (t *RedBlackTree) Add(key int, val interface{}) {
	isAdd, nd := t.Root.add(key, val)
	t.Size += isAdd
	t.Root = nd
	t.Root.Color = Black //根节点为黑色节点
}

// 递归写法:向树的root节点中插入key,val
// 返回1,代表加了节点
// 返回0,代表没有添加新节点,只更新key对应的value值
func (n *Node) add(key int, val interface{}) (int, *Node) {
	if n == nil { // 默认插入红色节点
		return 1, &Node{key, val, Red, nil, nil}
	}

	isAdd := 0
	if key < n.Key {
		isAdd, n.Left = n.Left.add(key, val)
	} else if key > n.Key {
		isAdd, n.Right = n.Right.add(key, val)
	} else {
		n.Value = val
	}

	return isAdd, n.updateRedBlackTree(isAdd)
}

// Get 查询数据
func (t *RedBlackTree) Get(key int) interface{} {
	return t.getRecursive(t.Root, key)
}

func (t *RedBlackTree) getRecursive(node *Node, key int) interface{} {
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

// 红黑树维护
func (n *Node) updateRedBlackTree(isChange int) *Node {
	// 0说明无新节点,不必维护
	if isChange == 0 {
		return n
	}

	node := n

	// 判断是否为情形2，是需要左旋转
	if node.Right.isRed() == Red && node.Left.isRed() != Red {
		node = node.leftRotate()
	}

	// 判断是否为情形3，是需要右旋转
	if node.Left.isRed() == Red && node.Left.Left.isRed() == Red {
		node = node.rightRotate()
	}

	// 判断是否为情形4，是需要颜色翻转
	if node.Left.isRed() == Red && node.Right.isRed() == Red {
		node.flipColors()
	}

	return node
}

func (n *Node) isRed() bool {
	if n == nil {
		return Black
	}
	return n.Color
}

//    n                      x
//  /   \     左旋转         /  \
// T1   x   --------->   Node   T3
//     / \              /   \
//    T2 T3            T1   T2
func (n *Node) leftRotate() *Node {
	// 左旋转
	retNode := n.Right
	n.Right = retNode.Left

	retNode.Left = n
	retNode.Color = n.Color
	n.Color = Red

	return retNode
}

//      n                    x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   Node
//  / \                       /  \
// y  T1                     T1  T2
func (n *Node) rightRotate() *Node {
	//右旋转
	retNode := n.Left
	n.Left = retNode.Right
	retNode.Right = n

	retNode.Color = n.Color
	n.Color = Red

	return retNode
}

// 颜色翻转
func (n *Node) flipColors() {
	n.Color = Red
	n.Left.Color = Black
	n.Right.Color = Black
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
