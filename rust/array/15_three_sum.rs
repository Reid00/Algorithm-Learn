
pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut result: Vec<Vec<i32>> = Vec::new();

    if nums.len() < 3 {
        return result
    }

    let mut nums = nums;
    nums.sort();

    let n = nums.len();

    // 固定第一个元素
    for first in 0..n {
        // 避免三元组重复
        if first > 0 && nums[first] == nums[first -1 ] {
            continue;
        }
        
        let target = -1 * nums[first];
        let mut third = n -1;
        
        // 遍历第二个元素
        for second in first+1..n {
            // 避免三元组重复
            if second > first + 1 && nums[second] == nums[second -1] {
                continue;
            }
            // NOTE 此处时循环判断，不是if
            while second < third && nums[second] + nums[third] > target {
                third -= 1;
            }

            if second == third {
                break;
            }

            if nums[second] + nums[third] == target {
                result.push(vec![nums[first], nums[second], nums[third]]);
            }
        }
    }
    result

}