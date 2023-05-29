

pub fn intersect(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
    // HashMap
    use std::collections::HashMap;

    let mut map1 = HashMap::new();

    for &v in nums1.iter() {
        let val = map1.entry(v).or_insert(0);
        *val += 1;
    }

    let mut result = Vec::new();

    for &v in nums2.iter() {
        if let Some(val) = map1.get_mut(&v) {
            if *val > 0 {
                result.push(v);
            }
            *val -= 1;
        }
    }
    result
}

pub fn intersect2(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
    // 排序 + 双指针
    let mut nums1 = nums1;
    let mut nums2 = nums2;
    nums1.sort();
    nums2.sort();

    let (mut idx1, mut idx2) = (0, 0);

    let mut result = vec![];

    while idx1 < nums1.len() && idx2 < nums2.len() {
        if nums1[idx1] < nums2[idx2] {
            idx1 += 1;
        }else if nums2[idx2] < nums1[idx1] {
            idx2 += 1;
        }else {
            result.push(nums1[idx1]);
            idx1 += 1;
            idx2 += 1;
        }
    }

    result
}