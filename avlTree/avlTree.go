package avlTree

import (
	"errors"
)

// AvlTree 平衡二叉树
type AvlTree struct {
	Root *Node
	Size int
}

type Node struct {
	Key    string
	Value  interface{}
	Left   *Node
	Right  *Node
	Height int
}

// Add 添加/更新节点
func (t *AvlTree) Add(key string, val interface{}) {
	isAdd, nd := t.Root.add(key, val)
	t.Size += isAdd
	t.Root = nd
}

func (n *Node) add(key string, val interface{}) (int, *Node) {
	if n == nil {
		return 1, &Node{key, val, nil, nil, 0}
	}

	isAdd := 0
	if key < n.Key {
		isAdd, n.Left = n.Left.add(key, val)
	} else if key > n.Key {
		isAdd, n.Right = n.Right.add(key, val)
	} else {
		n.Value = val
	}

	// 更新节点高度和维护平衡
	return isAdd, n.updateHeightAndBalance(isAdd)
}

// Remove 移除节点
func (t *AvlTree) Remove(key string) error {
	if t.Root == nil {
		return errors.New("failed to remove,avlTree is empty")
	}

	var isRemove int
	isRemove, t.Root = t.Root.remove(key)
	t.Size -= isRemove
	return nil
}

// 删除nd为根节点中key节点,返回更新了高度和维持平衡的新nd节点
// 返回值 1 表明有节点被删除,0 表明没有节点被删除
func (n *Node) remove(key string) (int, *Node) {
	// 找不到key对应node,返回0,nil
	if n == nil {
		return 0, nil
	}

	var retNode *Node
	var isRemove int
	switch {
	case key < n.Key:
		isRemove, n.Left = n.Left.remove(key)
		retNode = n
	case key > n.Key:
		isRemove, n.Right = n.Right.remove(key)
		retNode = n
	default:
		if n.Left != nil && n.Right != nil {
			// 待删除节点左右子树均不为空的情况
			// 找到比待删除节点大的最小节点,即右子树的最小节点
			retNode = n.Right.getMinNode()
			_, retNode.Right = n.Right.remove(key)
			retNode.Left = n.Left
		} else if n.Left != nil {
			retNode = n.Left
		} else {
			retNode = n.Right
		}
		isRemove = 1
	}

	if retNode == nil {
		return isRemove, retNode
	}

	retNode = retNode.updateHeightAndBalance(isRemove)
	return isRemove, retNode
}

// Get 获取节点数据
func (t *AvlTree) Get(key string) interface{} {
	return t.Root.get(key)
}

func (n *Node) get(key string) interface{} {
	if n == nil {
		return nil
	}

	if n.Key == key {
		return n.Value
	}

	if key < n.Key {
		return n.Left.get(key)
	}

	return n.Right.get(key)
}

// traverseInOrderKey 中序遍历所有key
func (n *Node) traverseInOrderKey(resp *[]string) {
	if n == nil {
		return
	}

	n.Left.traverseInOrderKey(resp)
	*resp = append(*resp, n.Key)
	n.Right.traverseInOrderKey(resp)
}

// 找出以n为根节点中最小值的节点
func (n *Node) getMinNode() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.getMinNode()
}

// 更新节点高度和维护平衡
func (n *Node) updateHeightAndBalance(isChange int) *Node {
	// 0说明无改变,不必更新
	if isChange == 0 {
		return n
	}

	// 更新高度
	n.Height = 1 + max(n.Left.getHeight(), n.Right.getHeight())

	// 平衡维护
	node := new(Node)
	if n.getBalanceFactor() > 1 {
		// 左左LL
		if n.Left.getBalanceFactor() >= 0 {
			node = n.rightRotate()
		} else { // 左右LR
			n.Left = n.Left.leftRotate()
			node = n.rightRotate()
		}
		return node
	}

	if n.getBalanceFactor() < -1 {
		// 右右RR
		if n.Right.getBalanceFactor() <= 0 {
			node = n.leftRotate()
		} else { // 右左RL
			node.Right = n.Right.rightRotate()
			node = n.leftRotate()
		}
		return node
	}

	return n
}

// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//        n                              x
//       / \                           /   \
//      x   T4     向右旋转 (n)        y     n
//     / \       - - - - - - - ->    / \   / \
//    y   T3                       T1  T2 T3 T4
//   / \
// T1   T2
func (n *Node) rightRotate() *Node {
	x := n.Left
	n.Left = x.Right
	x.Right = n

	n.Height = 1 + max(n.Left.getHeight(), n.Right.getHeight())
	x.Height = 1 + max(x.Left.getHeight(), x.Right.getHeight())
	return x
}

// 对节点n进行向左旋转操作，返回旋转后新的根节点x
//    n                             x
//  /  \                          /   \
// T1   x      向左旋转 (n)       n     y
//     / \   - - - - - - - ->   / \   / \
//   T2  y                     T1 T2 T3 T4
//      / \
//     T3 T4
func (n *Node) leftRotate() *Node {
	x := n.Right
	n.Right = x.Left
	x.Left = n

	n.Height = 1 + max(n.Left.getHeight(), n.Right.getHeight())
	x.Height = 1 + max(x.Left.getHeight(), x.Right.getHeight())
	return x
}

// 获取高度
func (n *Node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

// 获取的平衡因子
func (n *Node) getBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.Left.getHeight() - n.Right.getHeight()
}

func max(int1, int2 int) int {
	if int1 >= int2 {
		return int1
	}
	return int2
}
