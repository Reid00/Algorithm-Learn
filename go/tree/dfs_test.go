package tree

import (
	"testing"

	"github.com/Reid00/Algorithm-Learn/sort"
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

func TestMergeSort(t *testing.T) {
	actual := sort.MergeSort([]int{9, 2, 5, 7, 6, 3})

	expected := []int{2, 3, 5, 6, 7, 9}
	if intSliceEqual(actual, expected) {
		t.Log("the result is the same")
	} else {
		t.Errorf("actual is %v, expeced is %v\n", actual, expected)
	}
}

func TestPartition(t *testing.T) {
	a := []int{9, 2, 5, 7, 6, 3}
	sort.QuickSort(a)
	expected := []int{2, 3, 5, 6, 7, 9}
	if intSliceEqual(a, expected) {
		t.Log("TestPartition result right.")
	} else {
		t.Errorf("TestPartition result wrong. actual: %d", a)
	}

}

func TestOpenLock(t *testing.T) {
	deadlock := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	actual := openLock(deadlock, target)
	expected := 6
	if actual == expected {
		t.Log("TestOpenLock right")
	} else {
		t.Errorf("TestOpenLock failed, actaul %v, expect %v\n", actual, expected)
	}
}

func TestLevelTraversal(t *testing.T) {
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

	acutal := levelTraversal(&root)
	t.Log("acutal: ", acutal)
}
