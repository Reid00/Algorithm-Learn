



pub fn is_palindrome(s: String) -> bool {

    // 常规
    let s = s.to_lowercase().chars().collect::<Vec<char>>();

    let (mut l, mut r) = (0, s.len()-1);

    while l < r {

        while l < r && !is_al_num(s[l]) {
            l += 1;
        }
        while l < r && !is_al_num(s[r]) {
            r -= 1;
        }

        if l < r {
            if s[l] != s[r] {
                return  false;
            }
            l += 1;
            r -= 1;
        }
    }   
    true
}

pub fn is_al_num(s: char) -> bool{
    (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9')
}
// -----------------------------------------------------------------------------
// 迭代器的方法
pub fn is_palindrome(s: String) -> bool {

    // 过滤出ascii 字母字符和数字，并且转小写
    let mut chars = s.chars()
        .filter(|c|c.is_alphanumeric())
        .map(|c|c.to_ascii_lowercase());

    while let (Some(a), Some(b)) = (chars.next(), chars.next_back()) {
        if a != b {
            return false;
        }
    }

    true
}