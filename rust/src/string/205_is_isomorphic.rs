
// is_isomorphic 是否是同构字符串
pub fn is_isomorphic(s: String, t: String) -> bool {
    use std::collections::HashMap;

    let mut s_hmap:HashMap<char, char> = HashMap::new();
    let mut t_hmap:HashMap<char, char> = HashMap::new();


    for (a, b) in s.chars().zip(t.chars()) {
        if let Some(c) = s_hmap.insert(a, b) {
            // s2t hmap 存在a-> b 的映射
            // 如果不等于c 则不对
            if c != b {
                return false
            }
        }

        if let Some(c) = t_hmap.insert(b, a) {
            if c != a {
                return false
            }
        }

    }

    true
}