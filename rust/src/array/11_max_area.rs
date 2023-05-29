


pub fn max_area(height: Vec<i32>) -> i32 {

    let (mut  l, mut r) = (0 as usize, height.len()-1);
    let mut max_val = 0;

    while l < r {
        
        max_val = max(max_val, (min(height[l], height[r]) * (r - l) as i32));
        
        if height[l] < height[r] {
            l += 1 ;
        }else {
            r -= 1; 
        }

    }
    max_val
}