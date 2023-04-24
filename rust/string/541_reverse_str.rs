

pub fn reverse_str(s: String, k: i32) -> String {

    let mut s = s.chars().collect::<Vec<char>>();


    for i in (0..s.len()).step_by(2 *k as usize) {
        // 剩余字串长度是否大于k
        let length = if s.len() - i > k as usize {
            k as usize
        }else {
            s.len() - i
        };

        for j in 0..length/2 {
            s.swap(i + j, i + length -j -1);
        }

    }

    s.iter().collect()
}   