


fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {

    // 快慢指针，快指针向前走，把不重复的元素放到慢指针前面
    // let mut slow = 1;

    // for fast in 1..nums.len() {
    //     if nums[fast] != nums[fast-1] {
    //         nums[slow] = nums[fast];
    //         slow +=1 ;
    //     }
    // }
    // slow as i32

    // 调用API 的方式
    nums.dedup();
    nums.len() as i32
}
