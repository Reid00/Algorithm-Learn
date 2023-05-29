package queue

// 设计循环双端队列
type MyCircularDeque struct {
	front, rear int
	elements    []int
}

func Constructor3(k int) MyCircularDeque {
	elem := make([]int, k+1)
	return MyCircularDeque{
		elements: elem,
	}
}

// InsertFront 队首插入元素， front = (front -1) % len(elem) 需要考虑front 为负数
func (dq *MyCircularDeque) InsertFront(value int) bool {
	if dq.IsFull() {
		return false
	}
	dq.elements[(dq.front-1+len(dq.elements))%len(dq.elements)] = value
	dq.front = (dq.front - 1 + len(dq.elements)) % len(dq.elements)
	return true
}

func (dq *MyCircularDeque) InsertLast(value int) bool {
	if dq.IsFull() {
		return false
	}

	dq.elements[dq.rear] = value
	dq.rear = (dq.rear + 1) % len(dq.elements)
	return true
}

func (dq *MyCircularDeque) DeleteFront() bool {
	if dq.IsEmpty() {
		return false
	}

	dq.front = (dq.front + 1) % len(dq.elements)
	return true
}

func (dq *MyCircularDeque) DeleteLast() bool {
	if dq.IsEmpty() {
		return false
	}

	dq.rear = (dq.rear - 1 + len(dq.elements)) % len(dq.elements)
	return true
}

func (dq *MyCircularDeque) GetFront() int {
	if dq.IsEmpty() {
		return -1
	}

	return dq.elements[dq.front]
}

func (dq *MyCircularDeque) GetRear() int {
	if dq.IsEmpty() {
		return -1
	}

	return dq.elements[(dq.rear-1+len(dq.elements))%len(dq.elements)]
}

func (dq *MyCircularDeque) IsEmpty() bool {
	return dq.front == dq.rear
}

func (dq *MyCircularDeque) IsFull() bool {
	return (dq.rear+1)%len(dq.elements) == dq.front
}
