
// API
pub fn to_lower_case(s: String) -> String {
    s.to_lowercase()
}



pub fn to_lower_case2(s: String) -> String {
    let mut result = String::new();

    for ch in s.chars() {
        if ch >= 'A' && ch <= 'Z' {
            result.push((ch as u8 - 'A' as u8 + 'a' as u8) as char);
        }else {
            result.push(ch);
        }
    }
    result
}


pub fn to_lower_case3(s: String) -> String {
    String::from_utf8(s.as_bytes().into_iter().map(
        |&i| if i < 91 && i > 64 { i + 32 } else { i }).collect()).unwrap()
}