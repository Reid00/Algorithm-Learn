
pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
    use std::collections::{BinaryHeap, HashMap};
    use std::cmp::Reverse;

    let mut hmap = HashMap::new();

    let mut mini_heap = BinaryHeap::with_capacity(k as usize);

    // 用hashMap 统计次数
    nums.iter().for_each(|&x|{
        let cnt = hmap.entry(x).or_insert(0);
        *cnt += 1;
    });

    // 将{num: 次数} 用小堆根存储起来
    for (key, val) in hmap {

        if mini_heap.len() == mini_heap.capacity() {
            // Reverse 反转比较器 Reverse(i) > Reverse(j) means i < j
            // 查看次数是否需要加入堆中, 小根堆topk 频繁，堆顶小于freq val 应该加入堆中
            // reverse 反转这个情况
            if *mini_heap.peek().unwrap() > Reverse((val, key)) {
                mini_heap.pop();
            }else{
                continue;
            }
        }

        mini_heap.push(Reverse((val, key)));
    }

    let mut ans = vec![];

    while let Some(cnt) = mini_heap.pop() {
        ans.push(cnt.0.0);
    }
    ans

}