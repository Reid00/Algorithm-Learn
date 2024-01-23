/// 同构字符串
pub fn is_isomorphic(s: String, t: String) -> bool {
    // 构建两个hashmap 互相映射字符 s[i] -> t[i] && t[i]->s[i]
    use std::collections::HashMap;

    let (mut s2t, mut t2s) = (HashMap::new(), HashMap::new());

    for (a, b) in s.chars().zip(t.chars()) {
        if let Some(c) = s2t.insert(a, b) {
            if c != b {
                return false;
            }
        }

        if let Some(c) = t2s.insert(b, a) {
            if c != a {
                return false;
            }
        }
    }

    true
}

/// 同构字符串
pub fn is_isomorphic2(s: String, t: String) -> bool {
    // 构建两个hashmap 互相映射字符 s[i] -> t[i] && t[i]->s[i]
    use std::collections::HashMap;

    let (mut s2t, mut t2s) = (HashMap::new(), HashMap::new());

    let s: Vec<char> = s.chars().collect();
    let t: Vec<char> = t.chars().collect();

    for i in 0..s.len() {
        let (s_ch, t_ch) = (s[i], t[i]);

        if let Some(v) = s2t.get(&s_ch) {
            if *v != t_ch {
                return false;
            }
        }

        if let Some(v) = t2s.get(&t_ch) {
            if *v != s_ch {
                return false;
            }
        }

        s2t.insert(s_ch, t_ch);
        t2s.insert(t_ch, s_ch);
    }

    true
}
