


pub fn find_anagrams(s: String, p: String) -> Vec<i32> {
    // 滑动窗口
    let s: Vec<char> = s.chars().collect();
    let p: Vec<char> = p.chars().collect();

    if p.len() > s.len() {
        return vec![]
    }

    let mut diff = [0;26];
    let mut diff_cnt = 0;

    let mut ans: Vec<i32> = vec![];

    for (i, &ch) in p.iter().enumerate() {
        // 在 p 的长度范围内，s 对应各个位置的字符分布加一
        diff[s[i] as usize - 'a' as usize] += 1;
        // p 对应字符的减一，得到的结果即为差异
        diff[ch as usize - 'a' as usize] -= 1;
    }

    for &val in diff.iter() {
        if val != 0 {
            diff_cnt += 1;
        }
    }

    if diff_cnt == 0 {
        ans.push(0);
    }

    // 遍历s 后面的字符，依次更新diff 和 diff_cnt 
    for (i, &ch) in s.iter().enumerate() {
        if i == s.len() - p.len() {
            break;
        }

        // 遍历到i 意味着 i 所在字符被移出
        // 原先不相同，现在移出后相同
        if diff[ch as usize - 'a' as usize] == 1 {
            diff_cnt -= 1;
        }else if diff[ch as usize - 'a' as usize] == 0 {
            // 原先相同，现在不同
            diff_cnt += 1;
        }
        // diff window 中, ch 对应的字符减一
        diff[ch as usize - 'a' as usize] -= 1;

        // s 中对应的 i + p.len() 要移入窗口中        
        if diff[s[i+p.len()] as usize - 'a' as usize] == -1 {
            // 原先不同，差了一个，现在移入后相同
            diff_cnt -= 1;
        }else if diff[s[i + p.len()] as usize - 'a' as usize] == 0 {
            diff_cnt += 1;
        }
        
        diff[s[i + p.len()] as usize - 'a' as usize] += 1;

        if diff_cnt == 0 {
            ans.push(i as i32 + 1);
        }
    }

    ans
}