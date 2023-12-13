pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut nums = nums;
    nums.sort();

    let mut ans = vec![];

    // 枚举第一个元素
    for first in 0..nums.len() {
        // 排除重复元素

        // 优化1:
        if nums[first] > 0 {
            break;
        }

        if first >= 1 && nums[first] == nums[first - 1] {
            continue;
        }

        // 优化2：
        if first + 2 < nums.len() && nums[first] + nums[first + 1] + nums[first + 2] > 0 {
            break;
        }

        let mut third = nums.len() - 1;

        for second in first + 1..nums.len() {
            // 排除重复元素
            if second > first + 1 && nums[second] == nums[second - 1] {
                continue;
            }

            // 优化3
            if nums[first] + nums[second] + nums[third] < 0 {
                // 小于0  直接增加second
                continue;
            }

            // NOTE: 用while 而不是if
            while second < third && nums[first] + nums[second] + nums[third] > 0 {
                third -= 1;
            }

            if second == third {
                continue;
            }

            if nums[first] + nums[second] + nums[third] == 0 {
                ans.push(vec![nums[first], nums[second], nums[third]]);
            }
        }
    }
    ans
}
