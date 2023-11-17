// O(n)的空间复杂度
pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
    let length = nums.len();

    let mut result = vec![1; length];

    // l[i] 代表i左侧所有元素的乘积, 不包含i
    let mut l = vec![1; length];

    // r[i] 代表i 右侧所有元素的乘积, 从右到左填充,不包含i
    let mut r: Vec<i32> = vec![1; length];

    // 填充l array 的值
    for i in 1..length {
        l[i] = nums[i - 1] * l[i - 1];
    }
    print!("{:?}", l);

    // 填充r array 的值
    for i in (0..length - 1).rev() {
        r[i] = nums[i + 1] * r[i + 1];
    }

    for i in 0..length {
        result[i] = l[i] * r[i];
    }
    result
}

// O(1)的空间复杂度
pub fn product_except_self2(nums: Vec<i32>) -> Vec<i32> {
    let length = nums.len();

    let mut result = vec![1; length];

    // 想要O(1) 的空间复杂度 可以用result[i] 暂时保存i 右侧所有元素的乘积
    for i in (0..length - 1).rev() {
        // result[i] 此时表示i右侧所有元素的乘积
        result[i] = result[i + 1] * nums[i + 1]
    }

    // i 左侧所有元素的乘积在遍历的过程中 用一个变量逐步去更新
    let mut l = 1;

    for i in 0..length {
        // 等号左侧 result[i] 表示除i外的其他元素的乘积
        // 等号右侧 result[i] 表示i右侧的所有元素的乘积
        // l 表示 i左侧所有元素的乘积，初始为1，result[0] 则是 1 * r[0]
        result[i] = l * result[i];
        l = l * nums[i];
    }
    result
}
