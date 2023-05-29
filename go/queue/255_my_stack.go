package queue

// 注释可以参考 stack 中的255题

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (ms *MyStack) Push(x int) {
	// queue 先入先出，要想实现stack 后入先出
	// 就需要把新进入的元素（即后入的元素）添加到queue 的头部
	// 加入queue2 的首部
	ms.queue2 = append(ms.queue2, x)
	// queue1 的元素从头到尾出队列
	for _, v := range ms.queue1 {
		ms.queue2 = append(ms.queue2, v)
		ms.queue1 = ms.queue1[1:]
	}
	// ms.queue2 = append(ms.queue2, ms.queue1...)

	// 交换queue1 和 queue2，保持永远从queue取数据
	ms.queue1, ms.queue2 = ms.queue2, ms.queue1
}

func (ms *MyStack) Pop() int {
	v := ms.queue1[0]
	ms.queue1 = ms.queue1[1:]
	return v
}

func (ms *MyStack) Top() int {
	return ms.queue1[0]
}

func (ms *MyStack) Empty() bool {
	return len(ms.queue1) == 0
}
