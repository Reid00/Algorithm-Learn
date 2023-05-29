

pub fn valid_palindrome(s: String) -> bool {
    // 非迭代器解法
    let s: Vec<char> = s.chars().collect::<Vec<char>>();

    let (mut l, mut r) = (0, s.len()-1);

    while l < r {
        if s[l] == s[r] {

            l += 1;
            r -= 1;

        }else {
            let (mut flag1, mut flag2) = (true, true);
            // 删除 r 所在元素
            let (mut i, mut j ) = (l, r-1);
            while i < j {
                if s[i] != s[j] {
                    flag1 = false;
                    break;
                }
                i += 1;
                j -= 1;
            }
            // 删除 l 所在元素
            let (mut i, mut j) = (l+1, r);
            while i < j {
                if s[i] != s[j] {
                    flag2 = false;
                    break;
                }
                i += 1;
                j -= 1;
            }

            return flag1 || flag2
        }
    }

    true
}