pub fn is_anagram(s: String, t: String) -> bool {
    // HashMap
    use std::collections::HashMap;

    let s: Vec<char> = s.chars().collect();
    let t: Vec<char> = t.chars().collect();

    let mut hmap_s = HashMap::new();
    let mut hamp_t = HashMap::new();

    for ch in s {
        *hmap_s.entry(ch).or_insert(0) += 1;
    }

    // println!("s{:?}", s);

    for ch in t {
        *hamp_t.entry(ch).or_insert(0) += 1;
    }

    hamp_t == hmap_s
}

pub fn is_anagram2(s: String, t: String) -> bool {
    // HashMap  用一个HashMap
    use std::collections::HashMap;

    let s: Vec<char> = s.chars().collect();
    let t: Vec<char> = t.chars().collect();

    let mut hmap = HashMap::new();

    for ch in s {
        *hmap.entry(ch).or_insert(0) += 1;
    }

    // println!("s{:?}", s);

    for ch in t {
        *hmap.entry(ch).or_insert(0) -= 1;
    }

    !hmap.iter().any(|(_, &v)| v != 0)
    // hmap.iter().all(|(_, x)| *x == 0)
}
