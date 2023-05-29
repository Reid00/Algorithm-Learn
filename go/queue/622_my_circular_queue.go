package queue

// 设计循环队列
// 你的实现应该支持如下操作：
// MyCircularQueue(k): 构造器，设置队列长度为 k 。
// Front: 从队首获取元素。如果队列为空，返回 -1 。
// Rear: 获取队尾元素。如果队列为空，返回 -1 。
// enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
// deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
// isEmpty(): 检查循环队列是否为空。
// isFull(): 检查循环队列是否已满。
type MyCircularQueue struct {
	front    int // 队首
	rear     int // 队尾元素的下一个元素index(即需要enqueue 的元素下标)
	elements []int
}

// Constructor2 长度为k, 但是为了方便实现IsEmpty 和IsFull，elements 长度为k+1 作为哨兵
// IsEmpty => front == rear
// IsFull => (rear + 1) / len = front
func Constructor2(k int) MyCircularQueue {
	ele := make([]int, k+1)
	return MyCircularQueue{
		elements: ele,
	}
}

// EnQueue 向循环队列插入一个元素，需要判断是否已经满了
func (mq *MyCircularQueue) EnQueue(value int) bool {
	if mq.IsFull() {
		return false
	}

	mq.elements[mq.rear] = value
	// 取模为了防止mq.rear + 1 超过了容量大小
	mq.rear = (mq.rear + 1) % len(mq.elements)
	return true
}

// DeQueue 先判断是否为空
func (mq *MyCircularQueue) DeQueue() bool {
	if mq.IsEmpty() {
		return false
	}
	mq.front = (mq.front + 1) % len(mq.elements)
	return true
}

// Front 取出队首的元素
func (mq *MyCircularQueue) Front() int {
	if mq.IsEmpty() {
		return -1
	}
	return mq.elements[mq.front]
}

// Rear 取出队尾元素
func (mq *MyCircularQueue) Rear() int {
	if mq.IsEmpty() {
		return -1
	}
	// BUG 逻辑上长度为k+1, rear == front 时应该为空, 但是有可能满了之后，
	// 进行 dequeue, front from 0 => 1, 这时enqueue, rear = 0
	// mq.rear-1  有可能会是-1 防止index out of bound
	// return mq.elements[mq.rear-1]
	return mq.elements[(mq.rear-1+len(mq.elements))%len(mq.elements)]
}

// IsEmpty 判空
func (mq *MyCircularQueue) IsEmpty() bool {
	return mq.front == mq.rear
}

// IsFull 已满
func (mq *MyCircularQueue) IsFull() bool {
	return (mq.rear+1)%len(mq.elements) == mq.front
}
