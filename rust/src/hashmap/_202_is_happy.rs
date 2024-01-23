/// 快乐数
pub fn is_happy(n: i32) -> bool {
    // hashMap 原理详见Golang 的解法

    use std::collections::HashMap;

    let mut hmap = HashMap::new();

    let mut n = n;

    while n != 1 {
        // 如果hmap 中已经存在n 则会获得v 为true, 存在环，return false
        // if let Some(_) = hmap.insert(n, true) {
        //     return false;
        // }
        // or below
        if hmap.insert(n, true).is_some() {
            return false;
        }
        n = step(n);
    }
    n == 1
}

fn step(mut n: i32) -> i32 {
    let mut sum = 0;

    while n > 0 {
        sum += (n % 10) * (n % 10);
        n /= 10;
    }

    return sum;
}

fn is_happy2(n: i32) -> bool {
    // 用快慢指针检测环形链表

    let (mut slow, mut fast) = (n, step(n));

    while fast != 1 && fast != slow {
        slow = step(slow);
        fast = step(step(fast));
    }

    fast == 1
}
