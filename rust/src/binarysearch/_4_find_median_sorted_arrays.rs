/// 寻找两个正序数组的中位数
pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
    let (mut nums1, mut nums2) = (nums1, nums2);

    let l = nums1.len() + nums2.len();
    // 偶数
    if l & 1 == 0 {
        let e1 = get_k_nums(&mut nums1, &mut nums2, l / 2);
        let e2 = get_k_nums(&mut nums1, &mut nums2, l / 2 - 1);
        return (e1 as f64 + e2 as f64) / 2.0;
    } else {
        return get_k_nums(&mut nums1, &mut nums2, l / 2) as f64;
    }
}

/// 返回两个有序数组合并后第k小的元素， 计数从0开始
fn get_k_nums(mut nums1: &[i32], mut nums2: &[i32], mut k: usize) -> i32 {
    // let mut nums1 = nums1.as_slice();
    // let mut nums2 = nums2.as_slice();

    loop {
        let (l1, l2) = (nums1.len(), nums2.len());
        let (mid1, mid2) = (l1 >> 1, l2 >> 1);

        if l1 == 0 {
            return nums2[k];
        }

        if l2 == 0 {
            return nums1[k];
        }

        if k == 0 {
            return nums1[0].min(nums2[0]);
        }

        // k 在两个数组的前半部分元素中
        if k <= mid1 + mid2 {
            // 抛弃后面的元素, 保留前面的元素
            if nums1[mid1] < nums2[mid2] {
                nums2 = &nums2[..mid2];
            } else {
                nums1 = &nums1[..mid1];
            }
        } else {
            // k 在后半部分元素
            // 抛弃前面的元素
            if nums1[mid1] < nums2[mid2] {
                nums1 = &nums1[mid1 + 1..];
                // update k ,因为前面少了(mid1+1)元素
                k -= mid1 + 1;
            } else {
                nums2 = &nums2[mid2 + 1..];
                k -= mid2 + 1;
            }
        }
    }
}
