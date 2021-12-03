package avlTree

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestAvlTree_Get(t1 *testing.T) {
	t := &AvlTree{}

	//插入
	test01 := []struct {
		name  string
		key   string
		value interface{}
	}{
		{"==>add", "a", "北京"},
		{"==>add", "b", "上海"},
		{"==>add", "c", "广州"},
		{"==>add", "d", "深圳"},
		{"==>add", "e", "武汉"},
		{"==>add", "f", "大连"},
		{"==>add", "g", "西藏"},
		{"==>add", "h", "三亚"},
	}

	for _, tt := range test01 {
		t1.Run(tt.name, func(t1 *testing.T) {
			t.Add(tt.key, tt.value)
		})
	}

	//查询
	test02 := []struct {
		name string
		key  string
		want interface{}
	}{
		{"==>get", "b", "上海"},
		{"==>get", "d", "深圳"},
	}

	for _, tt := range test02 {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := t.Get(tt.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}

	//查看树结构（中序遍历）
	arr := make([]string, 0)
	t.Root.traverseInOrderKey(&arr)
	fmt.Println("==============中序遍历==================>>")
	fmt.Println(arr)
	fmt.Println("==============树的结构==================>>")
	b, _ := json.Marshal(t)
	fmt.Println(string(b))
}
