use::std::collections::HashMap;



pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {

    let mut hmap: HashMap<_, usize> = HashMap::with_capacity(nums.len());

   // 取出值具体匹配 2
   for i in 0..nums.len(){
        let val = nums[i];
        if let Some(k) = hmap.get(&(target - val)) {
            if *k != i {
                return vec![i as i32, *k as i32]
            }
        }

        hmap.insert(val, i);
   }

   panic!("not found")


    // 取出值具体匹配 1 leetcode 不知道为什么编译不通过
    // for (i, item) in nums.iter().enumerate(){
    //     if let Some(k) = hmap.get(&(target - *item)) {
    //         // k 能匹配到意味着有值
    //         if *k != i {
    //             return vec![i as i32, *k as i32]
    //         }
    //     }

    //     hmap.insert(item, i);
    // }


    // contains 判断 fail 不能输出idex
    // for (i, item) in nums.iter().enumerate() {
    //     if hmap.contains_key(&(target - *item)) {
    //         return vec![*item, target-*item]
    //     }else {
    //         hmap.insert(item, i);
    //     }
        
    // }

}