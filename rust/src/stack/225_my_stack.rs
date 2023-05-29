
// 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

struct MyStack {
    queue1: Vec<i32>,
    queue2: Vec<i32>
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyStack {

    fn new() -> Self {
        MyStack { queue1: vec![], queue2: vec![] }
    }
    
    fn push(&mut self, x: i32) {
        self.queue2.push(x);

        // BUG 直接消耗掉queue1 的元素
        // for item in self.queue1.into_iter() {
        //     self.queue2.push(item);
        // }
        // while let Some(item) = self.queue1.drain(..).next() {
        //     self.queue2.push(item)
        // }

        while !self.queue1.is_empty() {
            let first = self.queue1.remove(0);
            self.queue2.push(first);
        }

        std::mem::swap(&mut self.queue1, &mut self.queue2);
    }
    
    fn pop(&mut self) -> i32 {
        self.queue1.remove(0)
    }
    
    fn top(&self) -> i32 {
        *self.queue1.first().unwrap()
    }
    
    fn empty(&self) -> bool {
        self.queue1.is_empty()
    }
}