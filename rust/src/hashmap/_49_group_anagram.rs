
// groupAnagrams 将字母异位词分组，相同的放到一组
pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
    // 排序 + hashMap
    use std::collections::HashMap;

    let mut hmap: HashMap<Vec<char>, Vec<String>> = HashMap::new();

    let mut result: Vec<Vec<String>> = vec![];

    for str in strs.into_iter() {   
        let mut key: Vec<char> = str.chars().collect();
        key.sort();
        
        let v = hmap.entry(key).or_insert(vec![]);
        v.push(str);
    }

    for v in hmap.values() {
        result.push(v.clone());
    }

    result
}

// group_anagrams 计数的方式来实现
// 用hashMap 记录每个字母的个数，然后用hashmap 做key -> hashmap 不能做key
// 因为只有26个小写字母，用[i32;5] 数组做key
pub fn group_anagrams2(strs: Vec<String>) -> Vec<Vec<String>> {
    // 计数的方式
    use std::collections::HashMap;

    let mut hmap = HashMap::new();

    strs.into_iter().for_each(|x|{
        let mut cnt = [0;26];

        x.chars().for_each(|ch| cnt[ch as usize - 'a' as usize] += 1);
        hmap.entry(cnt).or_insert(vec![]).push(x);
    });

    hmap.values().cloned().collect()
}