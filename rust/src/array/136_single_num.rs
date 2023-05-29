

pub fn single_number(nums: Vec<i32>) -> i32 {

    let mut i = 0;

    for &v in nums.iter() {
        i ^= v;
    }
    i
}
