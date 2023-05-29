pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
    
    // hashmap
    use std::collections::HashMap;

    let hmap = HashMap::new();

    for (i, &v) in nums.iter().enumerate() {
        if let Some(&val) = hmap.get(&v){
            if i - val <= k as usize {
                return true
            }
        }

        hmap.insert(v, i);
    }
    false
}



pub fn contains_nearby_duplicate2(nums: Vec<i32>, k: i32) -> bool {
    // 滑动窗口

    use std::collections::HashSet;

    let mut hset = HashSet::new();
    let k = k as usize;

    for (i, &v) in nums.iter().enumerate() {
        if i > k {
            hset.remove(&nums[i-k-1]);
        }

        if hset.contains(&v){
            return true;
        }

        hset.insert(v);
    }
    false
    
}