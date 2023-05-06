

pub fn intersection(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
    // 集合
    use std::collections::HashSet;
    let mut set = HashSet::new();

    for &v in &nums1 {
        set.insert(v);
    }

    let mut result = vec![];
    let mut set2 = HashSet::new();

    for &v in nums2.iter() {
        if set.contains(&v) && !set2.contains(&v) {
            result.push(v);
        }
        set2.insert(v);
    }
    result
}

pub fn intersection2(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
    // 集合
    use std::collections::HashSet;
    // let set: HashSet<i32> = nums1.iter().cloned().collect();
    let set: HashSet<i32> = nums1.into_iter().collect();

    let mut result_set = HashSet::new();

    for &v in nums2.iter() {
        if set.contains(&v) {
            result_set.insert(v);
        }
    }

    result_set.into_iter().collect()
}

pub fn intersection3(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
 
    use std::collections::HashSet;
    nums1.into_iter().collect::<HashSet<i32>>()
    .intersection(&nums2.into_iter().collect::<HashSet<i32>>())
    .map(|&x|x)
    .collect()

}