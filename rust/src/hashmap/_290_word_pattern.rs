/// 单词规律
pub fn word_pattern(pattern: String, s: String) -> bool {
    // hashMap pattern -> s and s -> pattern

    use std::collections::HashMap;

    let (mut map1, mut map2) = (HashMap::new(), HashMap::new());

    if pattern.chars().collect::<Vec<_>>().len() != s.split(" ").collect::<Vec<_>>().len() {
        return false;
    }

    for (ch, word) in pattern.chars().zip(s.split(" ")) {
        if let Some(wd) = map1.insert(ch, word) {
            if wd != word {
                return false;
            }
        }

        if let Some(c) = map2.insert(word, ch) {
            if c != ch {
                return false;
            }
        }
    }
    true
}
