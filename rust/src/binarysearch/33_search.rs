

// search 搜索旋转排序数组: 两部分组成，每个部分都是单调递增的。并且第一部分的最小值大于第二部分的最大值
// 思路: 二分查找，如果mid == target 直接返回
// 如果不相等，判断mid 处于哪一部分的数组，
// 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环. 
pub fn search(nums: Vec<i32>, target: i32) -> i32 {
    let (mut l, mut r) = (0 as i32, (nums.len()-1) as i32);
    
    while l <= r {

        let mid = l + ((r-l)>>1);

        if nums[mid as usize] == target {
            return mid
        }

        // mid 在右侧的递增序列中
        if nums[mid as usize] < nums[r as usize] {

            // 判断target 是否在递增序列中的右侧
            if target > nums[mid as usize] && target <= nums[r as usize] {
                l = mid + 1;
            }else {
                r = mid - 1;
            }

        }else{

            // 判断target 是否在左侧递增序列中的左侧递增部分
            if target >= nums[l as usize] && target < nums[mid as usize] {
                r = mid - 1 ;
            }else {
                l = mid + 1;
            }
        }

    }
    -1
}