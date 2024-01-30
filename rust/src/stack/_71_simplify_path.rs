/// 简化路径
pub fn simplify_path(path: String) -> String {
    // 栈
    // 用/ 做分隔符拆开之后
    // 1. 正常目录、文件名
    // 2. . 目录 （skip）
    // 3. ..目录  (pop)
    // 4. 空格，有两个// 的情况拆分之后 (skip)

    let mut stack = vec![];

    for s in path.split('/') {
        if s == "" || s == "." {
            continue;
        } else if s == ".." {
            stack.pop();
        } else {
            stack.push(s)
        }
    }

    format!("/{}", stack.join("/"))
}

/// 简化路径
pub fn simplify_path_2(path: String) -> String {
    // 栈
    // 用/ 做分隔符拆开之后
    // 1. 正常目录、文件名
    // 2. . 目录 （skip）
    // 3. ..目录  (pop)
    // 4. 空格，有两个// 的情况拆分之后 (skip)

    let mut stack = vec![];

    for s in path.split('/') {
        match s {
            "." | "" => continue,
            ".." => {
                stack.pop();
            }
            _ => stack.push(s),
        };
    }

    format!("/{}", stack.join("/"))
}
