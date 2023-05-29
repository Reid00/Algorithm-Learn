

// next_greater_element 单调栈 + hashMap
pub fn next_greater_element(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
    use std::collections::HashMap;

    let mut hmap = HashMap::new();
    let mut result = vec![];
    let mut stack = vec![];

    // 从后往前遍历
    for v in nums2.iter().rev() {

        // 当v 大于栈顶元素， 说明栈顶的元素不是next greater elem
        while !stack.is_empty() && v >= stack.last().unwrap() {
            stack.pop();
        }

        if stack.is_empty() {
            hmap.insert(*v, -1);
        }else {
            hmap.insert(*v, *stack.last().unwrap());
        }

        stack.push(*v)
    }

    for item in nums1.into_iter() {
        result.push(*hmap.get(&item).unwrap());
    }
    result
}