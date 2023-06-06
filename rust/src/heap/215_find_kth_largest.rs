



pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
    // 长度为k的小根堆
    use std::collections::BinaryHeap;
    use std::cmp::Reverse;

    let mut heap = BinaryHeap::with_capacity(k as usize);

    for item in nums {

        if heap.len() == heap.capacity() {
            // Reverse 反转比较器，Reverse(i) > Reverse(j) means i < j
            // 小根堆 堆顶小于item 的时候 可以加入，Reverse 反转这个比较
            if *heap.peek().unwrap() > Reverse(item) {
                heap.pop();
            }else {
                continue;
            }
        }

        heap.push(Reverse(item));
    }

    heap.peek().unwrap().0
}
