
// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//  用栈模拟队列，以为着要用队列的特性，先入先出 来模拟栈后入先出
// 需要pop / peek 时， 将in 中的数据 写入out 中
struct MyQueue {
    in_stack: Vec<i32>,
    out_stack: Vec<i32>,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyQueue {

    fn new() -> Self {
        MyQueue { in_stack: vec![], out_stack: vec![] }
    }
    
    fn push(&mut self, x: i32) {
        self.in_stack.push(x);
    }
    

    fn in2out(&mut self) {
        while let Some(v) = self.in_stack.pop() {
            self.out_stack.push(v);
        }
    }

    fn pop(&mut self) -> i32 {
        if self.out_stack.is_empty() {
            self.in2out();
        }

        self.out_stack.pop().unwrap()
    }
    
    fn peek(&mut self) -> i32 {
        if self.out_stack.is_empty() {
            self.in2out();
        }

        *self.out_stack.last().unwrap()
    }
    
    fn empty(&self) -> bool {
        self.in_stack.is_empty() && self.out_stack.is_empty()
    }
}