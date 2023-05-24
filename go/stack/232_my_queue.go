package stack

// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//  用栈模拟队列，以为着要用队列的特性，先入先出 来模拟栈后入先出
// 需要pop / peek 时， 将in 中的数据 写入out 中
type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor2() MyQueue {
	return MyQueue{}
}

func (m *MyQueue) Push(x int) {
	m.inStack = append(m.inStack, x)
}

func (m *MyQueue) in2out() {
	for len(m.inStack) > 0 {
		popElem := m.inStack[len(m.inStack)-1]
		m.outStack = append(m.outStack, popElem)
		m.inStack = m.inStack[:len(m.inStack)-1]
	}
}

func (m *MyQueue) Pop() int {
	if len(m.outStack) == 0 {
		m.in2out()
	}

	elem := m.outStack[len(m.outStack)-1]
	m.outStack = m.outStack[:len(m.outStack)-1]
	return elem
}

func (m *MyQueue) Peek() int {
	if len(m.outStack) == 0 {
		m.in2out()
	}

	return m.outStack[len(m.outStack)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.inStack) == 0 && len(m.outStack) == 0
}
