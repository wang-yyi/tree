package redBlackTree

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestRedBlackTree_Get(t1 *testing.T) {
	t := &RedBlackTree{}

	//插入
	test01 := []struct {
		name  string
		key   int
		value interface{}
	}{
		{"==>add", 1, "北京"},
		{"==>add", 2, "上海"},
		{"==>add", 3, "广州"},
		{"==>add", 4, "深圳"},
		{"==>add", 5, "武汉"},
		{"==>add", 6, "大连"},
		{"==>add", 7, "西藏"},
		{"==>add", 8, "三亚"},
		{"==>add", 9, "沈阳"},
		{"==>add", 10, "长春"},
	}

	for _, tt := range test01 {
		t1.Run(tt.name, func(t1 *testing.T) {
			t.Add(tt.key, tt.value)
		})
	}

	//查询
	test02 := []struct {
		name string
		key  int
		want interface{}
	}{
		{"==>get", 2, "上海"},
		{"==>get", 4, "深圳"},
	}

	for _, tt := range test02 {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := t.Get(tt.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}

	//查看树结构（中序遍历）
	arr := make([]int, 0)
	t.Root.traverseInOrderKey(&arr)
	fmt.Println("==============中序遍历==================>>")
	fmt.Println(arr)
	fmt.Println("==============树的结构==================>>")
	b, _ := json.Marshal(t)
	fmt.Println(string(b))
}
