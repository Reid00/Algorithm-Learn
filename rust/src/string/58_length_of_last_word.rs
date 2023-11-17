pub fn length_of_last_word(s: String) -> i32 {
    // s.bytes().rev().skip_while(|&c| c== b' ')
    // .take_while(|&c|c != b' ').count() as i32

    s.into_bytes()
        .into_iter()
        .rev()
        .skip_while(|&c| c == b' ')
        .take_while(|&c| c != b' ')
        .count() as i32
}

fn length_of_last_word2(s: String) -> i32 {
    let words = s.trim().split_whitespace().rev().collect::<Vec<&str>>();

    words[0].len() as i32
}
