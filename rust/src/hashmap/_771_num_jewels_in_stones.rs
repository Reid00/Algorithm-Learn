pub fn num_jewels_in_stones(jewels: String, stones: String) -> i32 {
    // HashMap
    use std::collections::HashMap;

    let mut hmap = HashMap::new();

    for ch in stones.chars() {
        *hmap.entry(ch).or_insert(0) += 1;
    }

    let mut cnt = 0;

    for ch in jewels.chars() {
        if let Some(val) = hmap.get(&ch) {
            cnt += *val;
        }
    }
    cnt
}

pub fn num_jewels_in_stones2(jewels: String, stones: String) -> i32 {
    // HashSet + 迭代器
    use std::collections::HashSet;

    let hset = jewels.chars().collect::<HashSet<char>>();
    stones.chars().filter(|ch| hset.contains(ch)).count() as i32
}

pub fn num_jewels_in_stones3(jewels: String, stones: String) -> i32 {
    // 迭代器 + matches
    stones.matches(|c| jewels.contains(c)).count() as i32
}
