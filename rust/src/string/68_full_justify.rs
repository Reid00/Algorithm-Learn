// full_justify 左右文本对齐
pub fn full_justify(words: Vec<String>, max_width: i32) -> Vec<String> {
    let mut ans = vec![];
    let max_width = max_width as usize;
    let (mut i, n) = (0, words.len());

    loop {
        // 当前行的开始的位置，在words 中的index
        let start = i;
        // 当前行字符的长度
        let mut word_len = 0;

        while i < n && word_len + words[i].len() + i - start <= max_width {
            word_len += words[i].len();
            i += 1;
        }

        // 如果是最后一行
        if i == n {
            // 单词左对齐，然后长度不够，用' '补上
            let s = words[start..].join(" ");
            let space = " ".repeat(max_width - s.len());
            ans.push(format!("{}{}", s, space));
            return ans;
        }

        // 不是最后一行
        let space = max_width - word_len;
        // 如果是只有一个单词，靠左对齐，后面有空格填充
        if i - start == 1 {
            let s = format!("{}{}", words[start], " ".repeat(space));
            ans.push(s);
        } else {
            let (avg, extra) = (space / (i - start - 1), space % (i - start - 1));
            let avg_space = " ".repeat(avg);
            // 前extra 个单词 用 avg_space + 1 来join
            let s1 = words[start..start + extra + 1].join(&format!("{}{}", avg_space, " "));
            let s2 = words[start + extra + 1..i].join(&avg_space);
            ans.push(s1 + &avg_space + &s2);
        }
    }
}
