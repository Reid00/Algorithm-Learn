/// 电话号码的字母组合
pub fn letter_combinations(digits: String) -> Vec<String> {
    use std::collections::HashMap;

    if digits.len() == 0 {
        return vec![];
    }

    let mut phone_map = HashMap::new();
    phone_map.insert(b'2', "abc");
    phone_map.insert(b'3', "def");
    phone_map.insert(b'4', "ghi");
    phone_map.insert(b'5', "jkl");
    phone_map.insert(b'6', "mno");
    phone_map.insert(b'7', "pqrs");
    phone_map.insert(b'8', "tuv");
    phone_map.insert(b'9', "wxyz");

    let mut res = vec![];
    let path = "";

    fn backtrack(
        idx: usize,
        digits: &str,
        path: &str,
        res: &mut Vec<String>,
        map: &HashMap<u8, &str>,
    ) {
        if idx == digits.len() {
            res.push(path.to_string());
            return;
        }

        let digit = digits.as_bytes().get(idx).unwrap();

        let letters = map.get(digit).unwrap();

        for item in letters.chars() {
            let p = format!("{}{}", path, item);
            backtrack(idx + 1, digits, &p, res, map)
        }
    }

    backtrack(0, &digits, path, &mut res, &phone_map);
    res
}
