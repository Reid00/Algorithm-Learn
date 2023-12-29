///  串联所有单词的子串
pub fn find_substring(s: String, words: Vec<String>) -> Vec<i32> {
    let mut ans = vec![];

    let n = words.len();

    if n == 0 {
        return ans;
    }

    let k = words[0].len();
    let slen = s.len();

    if slen < k * n {
        return ans;
    }

    use std::collections::HashMap;
    let mut w_map: HashMap<&str, i32> = HashMap::new();

    // 构建words 的map 用于比较
    for w in words.iter() {
        w_map.entry(w).and_modify(|v| *v += 1).or_insert(1);
    }

    // let s:Vec<char> = s.chars().collect();
    for i in 0..k {
        // 窗口中有 words 的单词数量
        let mut count = 0;

        // 记录窗口中的数量
        let mut cnt_map: HashMap<&str, i32> = HashMap::new();

        let (mut l, mut r) = (i, i);

        while r <= slen - k {
            let word = &s[r..r + k];

            // word 在 w_map 中存在
            if w_map.contains_key(word) {
                let v = w_map.get(word).unwrap();

                while let Some(&cnt) = cnt_map.get(word) {
                    // 如果窗口 中存在，并且大于cnt 不满足条件，右移 窗口左指针
                    if cnt >= *v {
                        let w = &s[l..l + k];
                        cnt_map.entry(w).and_modify(|v| *v -= 1);

                        if count > 0 {
                            count -= 1;
                        }
                        l += k;
                    } else {
                        break;
                    }
                }

                cnt_map.entry(&word).and_modify(|v| *v += 1).or_insert(1);
                count += 1;
            } else {
                // word 在 w_map 中不存在
                while l < r {
                    let w = &s[l..l + k];
                    cnt_map.entry(w).and_modify(|v| *v -= 1);

                    if count > 0 {
                        count -= 1;
                    }

                    l += k;
                }
                l += k;
            }

            r += k;

            if count == n {
                ans.push(l as i32);
            }
        }
    }

    ans
}
