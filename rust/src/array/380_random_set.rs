use std::collections::HashMap;
#[derive(Debug)]
struct RandomizedSet {
    // data
    data: Vec<i32>,
    // {data: index}
    map: HashMap<i32, usize>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl RandomizedSet {
    fn new() -> Self {
        Self {
            data: vec![],
            map: HashMap::new(),
        }
    }

    // insert O(1) 数组的insert时间复杂度是O(1)的
    fn insert(&mut self, val: i32) -> bool {
        // 如果已经存在，返回false
        if let Some(val) = self.map.get(&val) {
            return false;
        }

        // 如果不在，val insert array， data:index insert map
        let idx = self.data.len();
        self.map.insert(val, idx);
        self.data.push(val);
        true
    }

    // remove O(1) 需要每次O(1)的时间复杂度找到数据，所以需要HashMap
    // 但是数据在array 里面，想要O(1)删除，需要将要删除的数据index 和最后一个交换
    fn remove(&mut self, val: i32) -> bool {
        // 如果数据不存在， 返回false
        if !self.map.contains_key(&val) {
            return false;
        }

        // 0. Get del element indx
        let idx = self.map.remove(&val).unwrap();
        // 1. Get last element index
        let last = self.data.len() - 1;
        // 2. update last element in map
        self.map.insert(self.data[last], idx);
        // 3. 在data 中交换最后一个元素和需要被删除的元素
        self.data.swap(idx, last);

        // 4. 在map，data 中删除val 的值
        // map 中前面已经做了
        self.map.remove(&val); // 此处还需要删除一次是因为可能被删除的元素是最后一个元素，在前面又被插入了
        self.data.pop();

        // // 在map中删除元素，并得到idx
        // let idx = self.map.remove(&val).unwrap();
        // // 在data中找到最后一个元素的值, 同时删除了
        // let last = self.data.pop().unwrap();
        // // 在data中删除idx的值，即
        // // 1. 把last 赋给self.data[idx],
        // // 2. 删除sef.data[last_index],但是通过pop 方法已经删除了

        // // 此处需要注意的是，如果idx是最后的元素，之前的pop 已经删除了，不需要下面得操作
        // if idx != self.data.len() {
        //     self.data[idx] = last;
        //     self.map.insert(last, idx);
        // }

        true
    }

    // 随机获取数据，想要每次都随机 需要一个数组+ 一个随机数
    fn get_random(&self) -> i32 {
        use rand::Rng;
        let mut rander = rand::thread_rng();
        let idx = rander.gen_range(0..self.data.len());

        self.data[idx]
    }
}
