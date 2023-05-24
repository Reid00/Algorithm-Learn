
// O(n) 的空间复杂度, 需要一个额外的辅助栈，存储最小值
struct MinStack {
    stack: Vec<i32>,
    min_stack: Vec<i32>,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MinStack {

    fn new() -> Self {
        let min_stack = vec![i32::MAX];
        let stack = vec![];
        MinStack{
            stack,
            min_stack,
        }
    }
    
    fn push(&mut self, val: i32) {
        self.stack.push(val);
        let old_min = self.min_stack.last().unwrap();

        use std::cmp::min;
        self.min_stack.push(min(*old_min, val));
    }
    
    fn pop(&mut self) {
        self.stack.pop();
        self.min_stack.pop();
    }
    
    fn top(&self) -> i32 {
        *self.stack.last().unwrap()
    }
    
    fn get_min(&self) -> i32 {
        *self.min_stack.last().unwrap()
    }
}