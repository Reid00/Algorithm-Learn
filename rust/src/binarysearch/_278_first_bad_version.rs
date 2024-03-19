



pub fn first_bad_version(n: i32) -> i32 {
	let (mut left, mut right) = (1, n);

    let isBadVersion = |x| -> bool {
        x == 4
    };

    while left < right {
        let mid = left + ((right-left)>>1) ;

        if isBadVersion(mid) {
            right = mid;
        }else {
            left = mid + 1
        }
    }

    left
}