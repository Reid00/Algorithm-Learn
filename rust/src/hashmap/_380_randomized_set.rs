

use std::collections::HashMap;
use rand::Rng;
struct RandomizedSet {
    data: Vec<i32>,
    index: HashMap<i32, usize>,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl RandomizedSet {

    fn new() -> Self {
        let data = vec![];
        let index = HashMap::new();

        RandomizedSet{
            data,
            index,
        }
    }
    
    fn insert(&mut self, val: i32) -> bool {
        if self.index.contains_key(&val) {
            return false
        }

        // data length
        let length = self.data.len();
        // store index
        self.index.insert(val, length);
        // store data
        self.data.push(val);

        true
    }
    
    fn remove(&mut self, val: i32) -> bool {
        if !self.index.contains_key(&val) {
            return false
        }

        // 和最后一个元素交换后再删除
        let last = self.data.pop().unwrap();
        let idx = self.index.remove(&val).unwrap();
        // FIX 最后一个元素不需要交换更新，直接删除即可
        if idx != self.data.len() {
            // index update
            self.index.insert(last, idx);
            // 数据udpate
            self.data[idx] = last;
        }

        true
    }
    
    fn get_random(&self) -> i32 {

        let mut rng = rand::thread_rng();
        let idx = rng.gen_range(0..self.data.len());

        self.data[idx]
    }
}