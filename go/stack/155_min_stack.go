package stack

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// Method 2: 不用辅助栈，额外的空间O(1)
// 思路: stack 中存放 当前push value 和 上次 lastMinVal 之间的差值 diff, 即是: diff = value - lastMinVal
// if diff < 0 => value 更小，他应该是当前这整个数组中最小的值，即minVal = value
// 否则，minValue 是数组中最小的值
// 这样Pop 的值如果是在栈顶，新的minVal = minVal - diff
type MinStack struct {
	data   []int
	minVal int
}

func Constructor() MinStack {
	// minVal 初始时长多少无所谓，后续push 元素的时候会判断初始情况
	return MinStack{
		data:   make([]int, 0),
		minVal: -1,
	}
}

func (m *MinStack) Push(val int) {
	if len(m.data) == 0 {
		// 初始时 存放原始值的差值是0
		m.data = append(m.data, 0)
		m.minVal = val
	} else {
		diff := val - m.minVal
		if diff < 0 {
			m.minVal = val
		}
		m.data = append(m.data, diff)
	}
}

func (m *MinStack) Pop() {
	if len(m.data) == 0 {
		return
	}

	diff := m.data[len(m.data)-1]
	// 栈顶是最小值元素，pop 之后需要更新 minVal
	if diff < 0 {
		m.minVal = m.minVal - diff
	}
	// 删除元素
	m.data = m.data[:len(m.data)-1]
}

func (m *MinStack) Top() int {
	if len(m.data) == 0 {
		return -1
	}
	diff := m.data[len(m.data)-1]
	// 栈顶即是最小值，value = minVal
	if diff < 0 {
		return m.minVal
	}
	return m.minVal + diff
}

func (m *MinStack) GetMin() int {
	return m.minVal
}

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// Method 1: 辅助栈，额外的空间O(n)
// type MinStack struct {
// 	stack    []int
// 	minStack []int
// }

// func Constructor() MinStack {
// 	stack := make([]int, 0)
// 	minStack := make([]int, 0)
// 	minStack = append(minStack, math.MaxInt64)
// 	return MinStack{
// 		stack:    stack,
// 		minStack: minStack,
// 	}
// }

// func (m *MinStack) Push(val int) {
// 	m.stack = append(m.stack, val)
// 	// FIX 防止index out of bound, 在Constructor中初始化一个最大值
// 	top := m.minStack[len(m.minStack)-1]
// 	// push 之后，minStack 都保存了stack 对应位置时的最小值
// 	m.minStack = append(m.minStack, min(top, val))
// }

// func (m *MinStack) Pop() {
// 	m.stack = m.stack[:len(m.stack)-1]
// 	m.minStack = m.minStack[:len(m.minStack)-1]
// }

// func (m *MinStack) Top() int {
// 	top := m.stack[len(m.stack)-1]
// 	return top
// }

// func (m *MinStack) GetMin() int {
// 	return m.minStack[len(m.minStack)-1]
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}

// 	return y
// }
