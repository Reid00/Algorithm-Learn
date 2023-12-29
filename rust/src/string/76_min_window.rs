/// 最小覆盖子串
pub fn min_window(s: String, t: String) -> String {
    // 滑动窗口解法
    // win 窗口中的字符, need 是需要的字符，都用 hashMap
    // 当win 满足条件时， 缩小窗口的左边界，l++

    use std::collections::HashMap;

    let (mut win, mut need) = (HashMap::new(), HashMap::new());

    for ch in t.chars() {
        need.entry(ch).and_modify(|x| *x += 1).or_insert(1);
    }

    // 滑动窗口的左右边界
    let (mut l, mut r) = (0, 0);
    // 窗口中已经 满足匹配条件 的字符
    let mut matched = 0;
    // 窗口满足条件时的长度
    let mut min_len = i32::MAX;
    // 窗口满足条件时，s 的索引
    let (mut start, mut end) = (0, 0);

    let s1 = s.chars().collect::<Vec<_>>();

    while r < s1.len() {
        let ch = s1[r];
        r += 1;

        // ch 在need 中 是需要的
        if let Some(cnt) = need.get(&ch) {
            let num = win.entry(ch).and_modify(|x| *x += 1).or_insert(1);

            // 字符 ch 个数满足后，matched ++
            if cnt == num {
                matched += 1;
            }
        }

        // 窗口中的元素满足条件之后
        while matched == need.len() {
            let ch_l = s1[l];

            if r - l < min_len as usize {
                min_len = (r - l) as i32;
                // 记录满足条件的s 的左右索引位置
                println!("win: {:?}, {}", l, r);
                (start, end) = (l, r);
            }

            // 如果 ch_l 在need 中, 判断matched 是否需要减少
            if let Some(&cnt) = need.get(&ch_l) {
                let num = win.get_mut(&ch_l).unwrap();
                if cnt == *num {
                    matched -= 1;
                }
                *num -= 1;
            }
            l += 1;
        }
    }

    if min_len == i32::MAX {
        return "".to_string();
    }
    s[start..end].to_string()
}
