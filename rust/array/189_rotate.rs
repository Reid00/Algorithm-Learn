

// API 组合调用
pub fn rotate(nums: &mut Vec<i32>, k: i32) {
    let k = k % nums.len() as i32;
    for _i in 0..k {
        if let Some(last) = nums.pop() {
            nums.insert(0, last);
        }
    }
}

// rotate2 array reverse
pub fn rotate2(nums: &mut Vec<i32>, k: i32) {
    // 数组翻转
    let n = nums.len();
    let k = k as usize % n;

    nums.reverse();
    nums[..k].reverse();
    nums[k..n].reverse();
}

// rotate3 内置rotate 方法
pub fn rotate3(nums: &mut Vec<i32>, k: i32) {

    let k = k as usize % nums.len();

    nums.rotate_right(k);

}

