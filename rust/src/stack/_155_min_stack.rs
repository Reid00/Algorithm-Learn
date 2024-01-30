// O(1) 的空间复杂度, data 存储 diff = val - last_min_val
// 即: Top value = diff + last_min_val
// if diff < 0 => min_val = val
// i32 会越界，用i64
struct MinStack {
    data: Vec<i64>,
    min: i64,
}
/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MinStack {
    fn new() -> Self {
        Self {
            data: vec![],
            min: 0,
        }
    }

    fn push(&mut self, val: i32) {
        // 初始化时，min = 第一次的值，data 存储0
        if self.data.len() == 0 {
            self.data.push(0);
            self.min = val as i64
        } else {
            // 否则
            let diff = val as i64 - self.min;
            if diff < 0 {
                self.min = val as i64
            }
            self.data.push(diff)
        }
    }

    fn pop(&mut self) {
        if self.data.len() == 0 {
            return;
        }

        let diff = self.data.pop().unwrap();
        // 最后一个元素即是最小值，pop 后需要更新新的min_val
        if diff < 0 {
            self.min = self.min - diff
        }
    }

    fn top(&self) -> i32 {
        if self.data.len() == 0 {
            return -1;
        }

        let diff = self.data.last().unwrap();
        if *diff < 0 {
            return self.min as i32;
        } else {
            return (self.min + *diff) as i32;
        }
    }

    fn get_min(&self) -> i32 {
        self.min as i32
    }
}

// // O(n) 的空间复杂度, 需要一个额外的辅助栈，存储最小值
// struct MinStack {
//     stack: Vec<i32>,
//     min_stack: Vec<i32>,
// }

// /**
//  * `&self` means the method takes an immutable reference.
//  * If you need a mutable reference, change it to `&mut self` instead.
//  */
// impl MinStack {
//     fn new() -> Self {
//         let min_stack = vec![i32::MAX];
//         let stack = vec![];
//         MinStack { stack, min_stack }
//     }

//     fn push(&mut self, val: i32) {
//         self.stack.push(val);
//         let old_min = self.min_stack.last().unwrap();

//         use std::cmp::min;
//         self.min_stack.push(min(*old_min, val));
//     }

//     fn pop(&mut self) {
//         self.stack.pop();
//         self.min_stack.pop();
//     }

//     fn top(&self) -> i32 {
//         *self.stack.last().unwrap()
//     }

//     fn get_min(&self) -> i32 {
//         *self.min_stack.last().unwrap()
//     }
// }
