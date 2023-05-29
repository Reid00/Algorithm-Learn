

pub fn longest_common_prefix(strs: Vec<String>) -> String {
    // 纵向扫描
    if strs.len() == 0 {
        return "".to_string()
    }

    let pre = &strs[0];

    for i in 0..pre.len() {
        for j in 1..strs.len() {
            if i == strs[j].len() || strs[j].as_bytes()[i] != pre.as_bytes()[i] {
                return pre[..i].to_string()
            }
        }
    }

    pre.to_string()
}




pub fn longest_common_prefix2(strs: Vec<String>) -> String {
    // 取出vec 中最大和最小的字符串作比较，忽略中间的
    // 因为最小和最大的字符串差异最大，他们的公共前缀就是所有字符串的公共前缀

    strs.iter().max().unwrap().chars()
    .zip(strs.iter().min().unwrap().chars())
    .take_while(|x| x.0 == x.1)
    .map(|x|x.0)
    .collect()

}