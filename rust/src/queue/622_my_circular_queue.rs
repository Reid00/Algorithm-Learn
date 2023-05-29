


struct MyCircularQueue {
    front: usize,
    rear: usize,
    elements: Vec<i32>,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyCircularQueue {

    fn new(k: i32) -> Self {
        let elem = vec![0;(k+1) as usize];
        MyCircularQueue{
            front: 0,
            rear: 0,
            elements: elem,
        }
    }
    
    fn en_queue(&mut self, value: i32) -> bool {
        if self.is_full() {
            return false
        }

        // rear 是下一个要插入的元素的index
        self.elements[self.rear] = value;
        self.rear = (self.rear + 1) % self.elements.len();

        true
    }
    
    fn de_queue(&mut self) -> bool {
        if self.is_empty() {
            return false
        }

        // 从队首删除一个元素，相当于front + 1
        self.front = (self.front + 1) % self.elements.len();

        true
    }
    
    fn front(&self) -> i32 {
        if self.is_empty() {
            return -1
        }

        self.elements[self.front]
    }
    
    fn rear(&self) -> i32 {
        if self.is_empty() {
            return -1
        }

        self.elements[(self.rear -1 + self.elements.len()) % self.elements.len()]
    }
    
    fn is_empty(&self) -> bool {
        self.front == self.rear
    }
    
    fn is_full(&self) -> bool {
        self.front == (self.rear + 1) % self.elements.len()
    }
}