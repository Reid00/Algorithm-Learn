

pub fn first_uniq_char(s: String) -> i32 {

    use std::collections::HashMap;

    let mut hmap = HashMap::new();

    for ch in s.chars() {
        let v = hmap.entry(ch).or_insert(0);
        *v += 1;
    }

    for (i, ch) in s.chars().enumerate() {
        if let Some(c) = hmap.get(&ch) {
            if *c == 1 {
                return i as i32
            }
        }
    }
    return -1
}


pub fn first_uniq_char2(s: String) -> i32 {

    use std::collections::HashMap;

    let mut hmap = HashMap::new();

    s.chars().for_each(|c| *hmap.entry(c).or_insert(0) += 1);

    for (i, ch) in s.chars().enumerate() {
        if hmap[&ch] == 1 {
            return i as i32
        }
    }
    -1
}