

pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
    // 左右指针
    let (mut l, mut r) = (0, numbers.len() - 1);

    while l < r {
        if numbers[l] + numbers[r] == target {
            return vec![(l + 1) as i32, (r + 1) as i32];
        }

        // 小于target，在不存在此时以l 为加数的组合了，重置尝试下个组合
        if numbers[l] + numbers[r] < target {
            l += 1;
            r = numbers.len() - 1;
        }

        if numbers[l] + numbers[r] > target {
            r -= 1;
        }
    }
    vec![-1, -1]
}
