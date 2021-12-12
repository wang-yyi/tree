package main

import (
	"fmt"
	"math/rand"
	"time"
	"tree/avlTree"
	"tree/binaryTree"
	"tree/redBlackTree"
)

const count = 1000000

var key int

func main() {
	rand.Seed(time.Now().Unix())
	nums := make([]int, 0)
	//制造二叉树极端情况（都在右边插入）
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	//随机取值
	for i := 0; i < count; i++ {
		k := rand.Intn(count)
		nums = append(nums, k)
	}

	key = nums[len(nums)-1]

	fmt.Println("============start==========>>")
	binary := binaryTree.BinaryTree{}
	binaryAdd(nums, &binary)
	binaryGet(key, &binary)
	fmt.Println()
	avl := avlTree.AvlTree{}
	avlAdd(nums, &avl)
	avlGet(key, &avl)
	fmt.Println()
	red := redBlackTree.RedBlackTree{}
	redBlackAdd(nums, &red)
	redBlackGet(key, &red)
}

func binaryAdd(arr []int, t *binaryTree.BinaryTree) {
	defer Cost("binaryAdd")()
	for _, s := range arr {
		t.Add(s, s)
	}
}

func binaryGet(key int, t *binaryTree.BinaryTree) {
	defer Cost("binaryGet")()
	fmt.Println("结果: ", t.Get(key))
}

func avlAdd(arr []int, t *avlTree.AvlTree) {
	defer Cost("avlAdd")()
	for _, s := range arr {
		t.Add(s, s)
	}
}

func avlGet(key int, t *avlTree.AvlTree) {
	defer Cost("avlGet")()
	fmt.Println("结果: ", t.Get(key))
}

func redBlackAdd(arr []int, t *redBlackTree.RedBlackTree) {
	defer Cost("redBlackAdd")()
	for _, s := range arr {
		t.Add(s, s)
	}
}

func redBlackGet(key int, t *redBlackTree.RedBlackTree) {
	defer Cost("redBlackGet")()
	fmt.Println("结果: ", t.Get(key))
}

func Cost(name string) func() {
	start := time.Now()
	return func() {
		costTime := time.Since(start)
		fmt.Println(fmt.Sprintf("%s: %v", name, costTime))
	}
}
