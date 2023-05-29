


pub fn majority_element(nums: Vec<i32>) -> i32 {
    // 排序
    // let mut nums = nums;
    // nums.sort();
    // nums[nums.len()/2]

    // 摩尔投票法 -- 用Option 枚举
    // let mut curr_elem = None;
    // let mut curr_count = 0;

    // for n in nums {
    //     match curr_elem {
    //         // 如果找到多数元素，对比新元素，继续计数
    //         Some(last) => {
    //             if last == n {
    //                 curr_count += 1;
    //             }else {
    //                 curr_count -= 1;
    //                 if curr_count == 0 {
    //                     curr_elem = None;
    //                 }
    //             }
    //         },
    //         None => {
    //             curr_elem = Some(n);
    //             curr_count += 1; 
    //         }
    //     }
    // }

    // match curr_elem {
    //     Some(majority) => majority,
    //     None => panic!("no majoriy")
    // }

     // 摩尔投票法 -- 不用枚举
    //  let mut marjor = 0;
    //  let mut count = 0;
    //  for n in nums {
    //     if count == 0 {
    //         marjor = n;
    //     }

    //     if marjor == n {
    //         count += 1;
    //     }else {
    //         count -= 1;
    //     }
    //  }
    //  marjor

    // Hmap
    use std::collections::HashMap;
    let mut hmap = HashMap::new();
    let ret = 0;

    for &n in &nums {
        let count = hmap.entry(n).or_insert(0);
        *count += 1;
    }

    for (k, v) in hmap {
        if v > (nums.len() / 2){
            return k
        }
    }
    ret
}
