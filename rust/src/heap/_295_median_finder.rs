use std::cmp::Reverse;
use std::collections::BinaryHeap;

/// 数据流的中位数
/// 中位数将数据分为两个部分，左边的元素比中位数小，右边的元素比其大
/// 用两个堆q_min 来存储左边的元素， q_max 存储右边的元素
/// q_min 用大根堆，能直接拿到其最大值，q_max 用小根堆能直接拿到其最小值
/// 如果是奇数，q_min 比q_max 多一个元素(且最多只能多一个元素)，那么q_min 堆顶元素就是中位数
/// 如果是偶数，q_min 的堆顶和q_max 的堆顶是中位数
/// 需要注意的是: add_num 的时候，需要平衡左右堆的元素数量。if num < q_min[0] 较小，入左侧
/// 此刻需要判断元素数量是不是相差超过1，if yes, q_min 堆顶移到 q_max 中
/// 如果q_max 的元素数量大于q_min, 则需要平衡
struct MedianFinder {
    q_min: BinaryHeap<i32>,          // 大根堆
    q_max: BinaryHeap<Reverse<i32>>, // 小根堆
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MedianFinder {
    fn new() -> Self {
        Self {
            q_min: BinaryHeap::new(),
            q_max: BinaryHeap::new(),
        }
    }

    // 添加元素，注意平衡左右堆的元素数量
    fn add_num(&mut self, num: i32) {
        // 如果是空， 先加入q_min 中  || 小于左侧元素的最大值
        if self.q_min.is_empty() || num < *self.q_min.peek().unwrap() {
            self.q_min.push(num);

            // 元素相差大于1
            if self.q_min.len() - self.q_max.len() > 1 {
                if let Some(big) = self.q_min.pop() {
                    self.q_max.push(Reverse(big));
                }
            }
        } else {
            self.q_max.push(Reverse(num));

            if self.q_max.len() > self.q_min.len() {
                if let Some(small) = self.q_max.pop() {
                    self.q_min.push(small.0);
                }
            }
        }
    }

    fn find_median(&self) -> f64 {
        // 奇数
        if self.q_min.len() > self.q_max.len() {
            let peek = self.q_min.peek().unwrap();
            return *peek as f64;
        }
        // 偶数
        if let (Some(ele1), Some(ele2)) = (self.q_min.peek(), self.q_max.peek()) {
            return (*ele1 as f64 + ele2.0 as f64) / 2.0;
        }
        0.0
    }
}
