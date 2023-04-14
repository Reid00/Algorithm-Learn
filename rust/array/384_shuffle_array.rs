


struct Solution {
    data: Vec<i32>,
    ori: Vec<i32>,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Solution {

    fn new(nums: Vec<i32>) -> Self {
        let nums = nums;
        let ori = nums.clone();

        Solution {data: nums, ori: ori }
    }
    
    fn reset(&mut self) -> Vec<i32> {
        self.data = self.ori.clone();
        self.data.clone()
    }
    
    fn shuffle(&mut self) -> Vec<i32> {
        // 洗牌算法
        let n = self.data.len();
        for i in 0..n {

            let j = i + rand::thread_rng().gen_range(0..n-i);
            self.data.swap(i, j);
        }

        self.data.clone()
    }
}