package main

import (
	"testing"
)

func TestDfs(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	acutal := preorderTraversal(&root)
	expected := []int{1, 2, 3, 4, 5, 6, 7}

	if !intSliceEqual(acutal, expected) {
		t.Errorf("preorderTraversal acutal: %v, expected: %v\n", acutal, expected)
	}

}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkIntSliceEqual(b *testing.B) {
	sa := []int{1, 2, 3, 4, 5}
	sb := []int{1, 2, 3, 4, 5, 6}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		intSliceEqual(sa, sb)
	}
}

func TestIntSliceEqual(t *testing.T) {
	sa := []int{}    // []int{} 空切片
	sb := []int(nil) // nil 转成[]int，还是nil

	actual := intSliceEqual(sa, sb)
	expected := false

	if actual != expected {
		t.Errorf("Actual: %v, expected: %v\n", actual, expected)
	}
}
