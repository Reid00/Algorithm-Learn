// convert N 字变换
pub fn convert(s: String, num_rows: i32) -> String {
    // 思路参考 Go 的实现
    if num_rows < 2 {
        return s;
    }

    let (mut i, mut flag) = (0, -1);

    let mut res: Vec<String> = vec!["".to_string(); num_rows as usize];

    for ch in s.chars() {
        res[i as usize].push(ch);

        if i == 0 || i == num_rows - 1 {
            flag = -flag;
        }
        i += flag;
    }

    res.join("")
}
