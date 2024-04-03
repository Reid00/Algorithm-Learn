package dptwodim

import "testing"

func TestWeightBag(t *testing.T) {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagweight := 4

	actual := weightBag(weight, value, bagweight)
	expected := 35 // 第一个和第二个

	if actual != expected {
		t.Error("0-1 背包test cases 失败")
	}

}
func TestWeightBag2(t *testing.T) {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagweight := 4

	actual := weightBag2(weight, value, bagweight)
	expected := 35 // 第一个和第二个

	if actual != expected {
		t.Error("0-1 背包test cases 失败")
	}

}
