pub fn h_index(citations: Vec<i32>) -> i32 {
    let (mut l, mut r) = (0, citations.len());

    while l < r {
        let mid = (l + r + 1) >> 1;

        let mut cnt = 0;

        for &v in citations.iter() {
            if v >= mid as i32 {
                cnt += 1;
            }
        }

        if cnt >= mid {
            l = mid;
        } else {
            r = mid - 1;
        }
    }
    l as i32
}
