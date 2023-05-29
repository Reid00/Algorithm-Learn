

pub fn contains_duplicate(nums: Vec<i32>) -> bool {


    // hashmap or hashSet
    // use std::collections::HashMap;

    // let mut hmap = HashMap::new();
    // for &item in &nums{
    //     if let Some(_) = hmap.get(&item) {
    //         return false
    //     }
    //     hmap.insert(item, ());
    // }
    // false

    // hashSet
    use std::collections::HashSet;

    let mut hset = HashSet::new();

    for &item in nums.iter() {
        if hset.contains(&item) {
            return true
        }
        hset.insert(item);
    }
    false
}