

/// 赎金信
pub fn can_construct(ransom_note: String, magazine: String) -> bool {
    use std::collections::HashMap;

    if magazine.len() == 0 || magazine.len() < ransom_note.len() {
        return false;
    }

    let mut hmap = HashMap::new();

    for ch in magazine.chars() {
        hmap.entry(ch).and_modify(|v| *v += 1).or_insert(1);
    }

    for ch in ransom_note.chars() {
        if let Some(cnt) = hmap.get_mut(&ch) {
            *cnt -= 1;
            if *cnt < 0 {
                return false;
            }
        } else {
            return false;
        }
    }

    true
}
