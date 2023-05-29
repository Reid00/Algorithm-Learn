// 单调队列, 用Vec超时
pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {

    let mut ans = vec![];

    let mut deque: Vec<usize> = vec![];

    for (i, &v) in nums.iter().enumerate() {

        // 当i 超过window 长度k后，剔除多余的元素
        // deque[0] 队列首部的索引
        if i >= k as usize && i - deque[0] >= k as usize  {
            deque.remove(0 as usize);
        }

        // 保证队列内递减
        while !deque.is_empty() && v >= nums[*deque.last().unwrap()] {
            deque.pop();
        }

        deque.push(i);

        if i as i32 >= k-1 {
            ans.push(nums[*deque.first().unwrap()]);
        }
    }

    ans
}


// 单调队列  用VecDeque
pub fn max_sliding_window2(nums: Vec<i32>, k: i32) -> Vec<i32> {
    use std::collections::VecDeque;

    let k = k as usize;

    let mut ans = vec![];

    let mut queue = VecDeque::with_capacity(k);

    fn push(q: &mut VecDeque<i32>, v:i32) {
        // 递增队列
        while let Some(&val) = q.front() {
            if v > val {
                q.pop_front();
            }else {
                break;
            }
        }
        q.push_front(v);
    }

    for i in 0..k-1 {
        push(&mut queue, nums[i]);
    }

    for i in k-1 ..nums.len() {
        println!("queue: {:?}", queue);
        push(&mut queue, nums[i]);

        ans.push(*queue.back().unwrap());
        // 计算完当前window 后，剔除边界元素
        // 如果队尾元素刚好是最大的元素，则从queue 的队尾删除
        // 否则在push 中删除
        if nums[i+1-k] == *queue.back().unwrap() {
            queue.pop_back();
        }
    }

    ans
}