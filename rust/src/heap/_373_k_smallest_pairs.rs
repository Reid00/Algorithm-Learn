use std::cmp::Reverse;
use std::collections::BinaryHeap;

/// 查找和最小的 K 对数字
/// 最小堆 + 多路归并
pub fn k_smallest_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> Vec<Vec<i32>> {
    // 1. 将nums1 前k个元素 和 nums2 的第一个元素的和 sum，及对应下标(i,j) 放入heap 中
    // 2. 出最小堆
    // 3. 然后加入nums1 i 和 nums2 的下一个元素 j+1 加入其中

    let mut heap = BinaryHeap::with_capacity(k as usize);
    let (mut nums1, mut nums2) = (nums1, nums2);
    let (mut n, mut m) = (nums1.len(), nums2.len());

    let flag = n > m;
    // 优化让较小的数组放到前面
    if flag {
        (n, m, nums1, nums2) = (m, n, nums2, nums1);
    }

    // k 较小的时候，可以用来优化
    if n > k as usize {
        n = k as usize;
    }

    let mut ans = vec![];

    // 入堆
    // k 有可能远大于n, 不能直接0..k
    for i in 0..n {
        // (sum, i, j)
        // 初始化的时候，将nums1 的前k 个元素 和 nums2[0] 组合入堆
        // 默认是最大堆， reverse 反转成最小堆
        heap.push((Reverse(nums1[i] + nums2[0]), i, 0));
    }

    while !heap.is_empty() && ans.len() < k as usize {
        // 堆顶元素
        let small = heap.pop().unwrap();
        let (x, y) = (small.1, small.2);
        // ans 要求第一个元素来自元素的nums1
        if flag {
            ans.push(vec![nums2[y], nums1[x]]);
        } else {
            ans.push(vec![nums1[x], nums2[y]]);
        }

        if y < m - 1 {
            heap.push((Reverse(nums1[x] + nums2[y + 1]), x, y + 1));
        }
    }

    ans
}
