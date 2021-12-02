package binaryTree

import (
	"reflect"
	"testing"
)

func TestBinaryTree_Get(t1 *testing.T) {
	t := &BinaryTree{}

	//插入
	test01 := []struct {
		name  string
		key   string
		value interface{}
	}{
		{"==>put", "a", "北京"},
		{"==>put", "b", "上海"},
		{"==>put", "c", "广州"},
		{"==>put", "d", "深圳"},
		{"==>put", "e", "武汉"},
		{"==>put", "f", "大连"},
		{"==>put", "g", "西藏"},
	}

	for _, tt := range test01 {
		t1.Run(tt.name, func(t1 *testing.T) {
			t.Put(tt.key, tt.value)
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
}
