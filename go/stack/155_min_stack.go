package stack

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// Method 2: 不用辅助栈，额外的空间O(1)
// 思路是stack 中存放当前push 和 min_val 之间的差值diff
// pop 的时候需要 diff + min_val
type MinStack struct {
	stack  []int
	minVal int
}

func Constructor() MinStack {
	stack := make([]int, 0)
	minVal := -1
	return MinStack{
		stack:  stack,
		minVal: minVal,
	}
}

func (m *MinStack) Push(val int) {
	if len(m.stack) == 0 {
		// 如果开始栈中为空, diff = 0
		m.stack = append(m.stack, 0)
		m.minVal = val
	} else {
		diff := val - m.minVal
		if diff < 0 {
			m.minVal = val
		}
		m.stack = append(m.stack, diff)
	}

}

func (m *MinStack) Pop() {
	if len(m.stack) == 0 {
		return
	}

	diff := m.stack[len(m.stack)-1]
	// 说明存储的diff 是最顶上的数据，pop 出来之后需要更新下minVal
	if diff < 0 {
		m.minVal = m.minVal - diff
	} else {
		// diff 大于0, 说明minVal 和 top 元素无关
	}
	m.stack = m.stack[:len(m.stack)-1]

}

func (m *MinStack) Top() int {
	if len(m.stack) == 0 {
		return -1
	}

	diff := m.stack[len(m.stack)-1]
	if diff < 0 {
		// 之前push 进来的栈顶元素是 minVal
		return m.minVal
	} else {
		return diff + m.minVal
	}

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
