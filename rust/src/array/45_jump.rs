pub fn jump(nums: Vec<i32>) -> i32 {
    let mut steps = 0; // 需要跳跃的次数

    let n = nums.len();

    let mut farthest = 0; // 本次跳跃的最远位置
    let mut next_farthest = 0; // 下次跳跃的最远位置

    // 跳跃不需要遍历最后的位置下标n-1,因为最后的位置不需要跳跃
    for i in 0..n - 1 {
        next_farthest = next_farthest.max(i as i32 + nums[i]);

        // 本次跳远的所有位置都检查了
        if i as i32 == farthest {
            farthest = next_farthest; // 更新下一次跳跃的最远位置，准备下一次跳跃
            steps += 1;
        }
    }

    steps
}
