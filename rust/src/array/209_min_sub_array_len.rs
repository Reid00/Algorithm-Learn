pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
    // 快慢指针，弄个滑动窗口
    // 滑动窗口中sum 是两个左右指针夫妻共同的财产。好不容易共同财产大过target，记录一下两指针之间的距离，
    // 当财产满足时，左指针就会胡乱花销(做指针右移动)
    // 为了保证财产不缩水，右指针就会继续工作(右指针右移)， 直到左右指针离得最近的时候

    let (mut slow, mut fast) = (0, 0);

    let mut sum = 0;
    let mut ans = i32::MAX;

    while fast < nums.len() {
        sum += nums[fast];

        while sum >= target {
            ans = ans.min((fast - slow + 1) as i32);
            sum -= nums[slow];
            slow += 1;
        }
        fast += 1;
    }

    if ans == i32::MAX {
        return 0;
    }

    ans
}
