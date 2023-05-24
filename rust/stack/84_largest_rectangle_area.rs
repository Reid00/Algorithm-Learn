


pub fn largest_rectangle_area(heights: Vec<i32>) -> i32 {
    let mut ans = 0;
    let mut stack = vec![];
    let n = heights.len();

    // 等于n 为了把压栈的数据可以全部出栈，因为i=n 时，height = 0 
    for i in 0..=n {

        let mut height = 0;
        if i < n {
            height = heights[i];
        }

        if stack.is_empty() {
            // 压栈宽和高
            stack.push(vec![1, height]);
            continue;
        }

        if let Some(v) = stack.last_mut() {
            // 比最后一个元素的高 大的元素都压栈
            if height == v[1] {
                v[0] += 1;
                continue;
            }
            if height > v[1] {
                stack.push(vec![1, height]);
                continue;
            }
        }

        let mut width = 0;
        // loop {
        //     if let Some(v) = stack.last() {
        //         if height < v[i] {
        //             // 出栈
        //             let top = stack.pop().unwrap();
        //             width += top[0];
        //             ans = ans.max(width * top[1]);
        //         }else{
        //             break;
        //         }
        //     }else{
        //         break;
        //     }
        // }
        while !stack.is_empty() {
            // 出栈比当前高度大的所有宽高组合
            if let Some(v) = stack.last() {
                if height < v[1] {
                    // 出栈
                    let top = stack.pop().unwrap();
                    width += top[0];
                    ans = ans.max(width * top[1]);
                }else {
                    // height > 高度退出循环
                    break;
                }
            }else {
                // last 为None 推出循环，实际上 !stack.is_empty 已经做了校验
                break;
            }
        }
        // 比height 高的出栈后，把height 对应的宽高压栈
        stack.push(vec![width + 1, height]);
    }

    ans
}