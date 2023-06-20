


// zigzag_level_order 二叉树的锯齿形层序遍历
// 引入flag，偶数从左到右，奇数数从右到左
pub fn zigzag_level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
    if root.is_none() {
        return vec![];
    }

    let mut ans = vec![];

    use std::collections::VecDeque;

    let mut queue = VecDeque::new();
    queue.push_back(root);

    // 第一层默认false 不用把结果 reverse
    let mut flag = false;   

    while !queue.is_empty() {

        let mut level = vec![];

        for _ in 0..queue.len() {
            
            let head = queue.pop_front().unwrap().unwrap();
            level.push(head.borrow().val);

            if head.borrow().left.is_some() {
                queue.push_back(head.borrow_mut().left.take());
            }

            if head.borrow().right.is_some() {
                queue.push_back(head.borrow_mut().right.take());
            }
        }

        if flag {
            level.reverse();
        }

        flag = !flag;

        ans.push(level);

    }

    ans

}