
struct MyCircularDeque {
    front: usize,
    rear: usize,
    elements: Vec<i32>,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyCircularDeque {

    fn new(k: i32) -> Self {
        let elements = vec![0; (k+1) as usize];
        MyCircularDeque {front: 0, rear: 0, elements}
    }
    
    fn insert_front(&mut self, value: i32) -> bool {
        if self.is_full() {
            return false
        }

        let n = self.elements.len();
        self.elements[(self.front -1 +n ) % n] = value;

        self.front = (self.front -1 + n) % n;
        true
    }
    
    fn insert_last(&mut self, value: i32) -> bool {
        if self.is_full() {
            return false
        }

        self.elements[self.rear] = value;

        self.rear = (self.rear + 1) % self.elements.len();

        true
    }
    
    fn delete_front(&mut self) -> bool {
        if self.is_empty() {
            return false
        }

        self.front = (self.front + 1) % self.elements.len();

        true
    }
    
    fn delete_last(&mut self) -> bool {
        if self.is_empty() {
            return false
        }

        self.rear = (self.rear -1 + self.elements.len()) % self.elements.len();

        true
    }
    
    fn get_front(&self) -> i32 {
        if self.is_empty() {
            return -1
        }

        self.elements[self.front]
    }
    
    fn get_rear(&self) -> i32 {
        if self.is_empty() {
            return -1
        }
        self.elements[(self.rear-1+self.elements.len()) % self.elements.len()]
    }
    
    fn is_empty(&self) -> bool {
        self.front == self.rear
    }
    
    fn is_full(&self) -> bool {
        self.front == (self.rear + 1) % self.elements.len()
    }
}