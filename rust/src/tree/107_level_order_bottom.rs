

// level_order_bottom 从底部层次遍历，102 的结果反转
pub fn level_order_bottom(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {

    if root.is_none() {
        return vec![];
    }

    let mut ans = vec![];

    use std::collections::VecDeque;
    let mut queue = VecDeque::new();
    queue.push_back(root);


    while !queue.is_empty() {

        let mut level = vec![];

        (0..queue.len()).for_each(|_| {

            let head = queue.pop_front().unwrap().unwrap();
            level.push(head.borrow().val);

            if head.borrow().left.is_some() {
                queue.push_back(head.borrow_mut().left.take());
            }

            if head.borrow().right.is_some() {
                queue.push_back(head.borrow_mut().right.take());
            }
        });

        ans.push(level);

    }

    ans.reverse();
    ans
}