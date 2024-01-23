
// 自建最小堆
struct KthLargest {
    // 最小堆的数据
    data: Vec<i32>,
    // 最小堆的最大长度，默认为k
    size: i32,
    // 堆中已有的数据个数
    count: i32,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl KthLargest {

    fn new(k: i32, nums: Vec<i32>) -> Self {
        let mut kth = KthLargest{
            size: k,
            count: 0,
            data: vec![],
        };

        for &i in nums.iter() {
            kth.add(i);
        }
        kth
    }
    
    fn add(&mut self, val: i32) -> i32 {

        if self.count < self.size -1  {
            self.data.push(val);
            self.count += 1;
        }else if self.count == self.size -1 {
            // 加入数据后 容量刚满，进行小根堆的堆化
            self.data.push(val);
            self.count += 1;

            // 从第一个非叶子节点进行堆化
            for i in (0..self.data.len()/2 -1).rev() {
                self.heapfy(i as i32, self.data.len() as i32);
            }
        }else {
            // 当容量已经满了, val 大于堆顶的元素, 替换堆顶元素，再进行堆化
            if self.data[0] < val {
                self.data[0] = val;
                self.heapfy(0, self.data.len() as i32);
            }
        }

        self.data[0]
    }

    fn heapfy(&mut self, idx: i32, length: i32){

        let mut idx = idx as usize;
        let length = length as usize;

        loop {
            // 左孩子
            let l = 2 * idx + 1;
            let r = 2 * idx + 2;

            // 小根堆，父节点的值比左右孩子的小
            let mut smallest = idx;

            if l < length && self.data[smallest] > self.data[l] {
                smallest = l;
            }

            if r < length && self.data[smallest] > self.data[r] {
                smallest = r;
            }

            // 如果smallest = idx 根节点就是最小值，符合小根堆的特征，退出循环
            if smallest == idx {
                break;
            }

            // 使idx 所在的子树符堆特征
            self.data.swap(idx,smallest);

            // idx 所在子树满足堆特征后，交换后的smallest 也就是l/r 也需要满足
            idx = smallest;
        }

    }
}



// 官方自带最小堆
use std::collections::BinaryHeap;
use std::cmp::Reverse;
// Reverse(i) 保证按照升序放入的大堆根，输出是个小根堆的顺序

struct KthLargest2 {
   data: BinaryHeap<Reverse<i32>>,
}

impl KthLargest2 {
    fn new(k: i32, nums: Vec<i32>) -> Self {

        let mut kth = Self{
            data: BinaryHeap::with_capacity(k as usize),
        };

        nums.iter().for_each(|&x|kth.push(x));

        kth
    }

    fn add(&mut self, val: i32) -> i32 {
        self.push(val);
        self.peek()
    }

    fn push(&mut self, n: i32) {
        
        if self.data.len() == self.data.capacity() {
            // 容量已满，n 是否大于栈顶元素
            // 正常小根堆，如果新元素大于堆顶元素，则删除堆顶元素，加入新元素，然后堆化
            // Reverse 之后相反过来
            if *self.data.peek().unwrap() > Reverse(n) {
                self.data.pop();
            }else {
                return;
            }
        }
        // 容量未满，直接push
        self.data.push(Reverse(n));
    }
    
    fn peek(&self) -> i32 {
        self.data.peek().unwrap().0
    }
}