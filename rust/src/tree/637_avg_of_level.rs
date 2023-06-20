

// average_of_levels 给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。
// BFS 每层的结果，然后计算值
pub fn average_of_levels(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<f64> {

    if root.is_none() {
        return vec![]
    }

    let mut ans = vec![];

    use std::collections::VecDeque;
    let mut queue = VecDeque::new();
    queue.push_back(root);

    while !queue.is_empty() {

        let mut level = 0_f64;
        let cnt = queue.len();

        (0..cnt).for_each(|_| {

            let head = queue.pop_front().unwrap().unwrap();
            level += head.borrow().val as f64;

            if head.borrow().left.is_some() {
                queue.push_back(head.borrow_mut().left.take());
            }

            if head.borrow().right.is_some() {
                queue.push_back(head.borrow_mut().right.take());
            }
        });

        let avg = level / cnt as f64;
        ans.push(avg);
    }

    ans
}