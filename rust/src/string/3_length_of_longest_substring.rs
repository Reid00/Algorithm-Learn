/// 无重复字符的最长子串
pub fn length_of_longest_substring(s: String) -> i32 {
    // if s.len() == 0 {
    //     return 0
    // }
    if s.is_empty() {
        return 0;
    }

    let (mut slow, mut fast) = (0, 0);

    let mut length = i32::MIN;

    let s = s.chars().collect::<Vec<_>>();

    use std::collections::HashMap;
    let mut win = HashMap::new();

    while fast < s.len() {
        let ch = s[fast];
        fast += 1;
        *win.entry(ch).or_insert(0) += 1;

        while *win.get(&ch).unwrap() > 1 {
            let slow_ch = s[slow];
            win.entry(slow_ch).and_modify(|v| *v -= 1);
            slow += 1;
        }

        length = length.max((fast - slow) as i32);
    }
    length
}
