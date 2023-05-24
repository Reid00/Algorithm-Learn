package stack

// 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
// 队列，先入先出，栈后入先出
// 队列实现栈 就是要把最后push 的元素放到队列的头部，就可以后入先出了
// 相当于把queue 当成一个环，push 一个元素后，把之前的元素从头push 到新的元素之后
type MyStack struct {
	queue1 []int // 每次push一个元素，就把queue1 的原先的元素从头push 到queue2中

	queue2 []int
}

func Constructor3() MyStack {
	return MyStack{}
}

func (ms *MyStack) Push(x int) {
	ms.queue2 = append(ms.queue2, x)
	// 把queue 中的元素，从头push 到queue2 中
	for _, v := range ms.queue1 {
		ms.queue2 = append(ms.queue2, v)
		ms.queue1 = ms.queue1[1:]
	}
	// for len(ms.queue1) > 0 {
	// 	first := ms.queue1[0]
	// 	ms.queue2 = append(ms.queue2, first)
	// 	ms.queue1 = ms.queue1[1:]
	// }

	// 交换queue1 和 queue2, 保持queue1 永远是data集合，queue2 是empty
	ms.queue1, ms.queue2 = ms.queue2, ms.queue1
}

func (ms *MyStack) Pop() int {
	first := ms.queue1[0]
	ms.queue1 = ms.queue1[1:]
	return first
}

func (ms *MyStack) Top() int {
	return ms.queue1[0]
}

func (ms *MyStack) Empty() bool {
	return len(ms.queue1) == 0
}
