

pub fn reverse_words(s: String) -> String {

    // call API
    let s = s.trim();

    let mut words: Vec<&str> = s.split_whitespace().collect();

    words.reverse();

    words.join(" ").to_string()
}

pub fn reverse_words(s: String) -> String {

    // call API
    s.split_whitespace().rev().collect::<Vec<&str>>()
    .join(" ")
}


pub fn reverse_words(s: String) -> String {

    // 不使用API
    let mut s: Vec<char> = s.chars().collect();

    // 快慢指针
    let (mut slow, mut fast) = (0, 0);

    // 去除首部空格
    while fast < s.len() && s[fast] == ' ' {
        fast += 1;
    }

    // 去除中间位置的冗余空格
    while fast < s.len() {
        // fast 是usize， 当fast == 0 的时候，fast -1 会overflow
        if fast > 1 && s[fast] == s[fast-1] && s[fast] == ' ' {
            fast += 1;
            continue;
        }
        s[slow] = s[fast];
        fast += 1;
        slow += 1;
    }

    // 去除尾部位置的空格
    let t ;
    if slow -1 > 0 && s[slow-1] == ' ' {
        // s.truncate(slow -1);
        t = &mut s[..slow-1];
    }else {
        t = &mut s[..slow];
    }

    let s = t;
    // 反转整个字符串
    s.reverse();

    // 反转单词
    let mut i = 0;
    while i < s.len() {
        let mut j = i;
        // 以空格为单位进行单词切割
        while j < s.len() && s[j] != ' ' {
            j +=1;
        }
        println!("s: {:?}", &s[i..j]);
        s[i..j].reverse();

        i = j +1 ;
    }

    s.iter().collect()

}