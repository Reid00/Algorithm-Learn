/// 构建一个长度为k 的小根堆，比k大的元素进入小根堆中，k就是最小的元素
pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
    // 长度为k的小根堆
    use std::cmp::Reverse;
    use std::collections::BinaryHeap;

    let mut heap = BinaryHeap::with_capacity(k as usize);

    for item in nums {
        if heap.len() == heap.capacity() {
            // Reverse 反转比较器，Reverse(i) > Reverse(j) means i < j
            // 找到 第k 大的元素，长度为k 的小根堆，堆顶就是需要的k
            // 所以容量满了以后，比新元素item 比堆顶大了之后 可以入堆，即Reverse(item) < old
            if let Some(old) = heap.peek() {
                if Reverse(item) < *old {
                    // pop 然后入堆
                    heap.pop();
                } else {
                    // 不需要入堆
                    continue;
                }
            }
        }

        heap.push(Reverse(item));
    }

    heap.peek().unwrap().0
}
