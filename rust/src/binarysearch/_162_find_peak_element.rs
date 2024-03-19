/// 寻找峰值  峰值元素是指其值严格大于左右相邻值的元素。
pub fn find_peak_element(nums: Vec<i32>) -> i32 {
    // 因为左右两侧都是负无穷，所以沿着上坡的方向一定能找到一个峰值
    let n = nums.len() as i64;
    let (mut l, mut r) = (0, n - 1);

    let get = |i: i64| {
        if i < 0 || i >= n {
            return i32::MIN;
        }
        return nums[i as usize];
    };

    while l <= r {
        let mid = l + ((r - l) >> 1);

        if get(mid) > get(mid - 1) && get(mid) > get(mid + 1) {
            return mid as i32;
        }

        // 上坡方向, 峰值在右侧
        if get(mid) < get(mid + 1) {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    l as i32
}
