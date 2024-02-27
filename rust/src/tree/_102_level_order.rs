use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

// level_order 层序遍历 BFS
pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
    // if root == None {
    //     return vec![]
    // }

    if root.is_none() {
        return vec![];
    }

    use std::collections::VecDeque;

    let mut ans = vec![];

    // 因为需要在头部进行删除元素，VecDeque 比Vec 更合适
    let mut queues = VecDeque::new();
    queues.push_back(root);

    while !queues.is_empty() {
        let mut level_res: Vec<i32> = vec![];

        for _ in 0..queues.len() {
            let head = queues.pop_front().unwrap().unwrap();

            level_res.push(head.borrow().val);

            if head.borrow().left.is_some() {
                queues.push_back(head.borrow_mut().left.take());
            }

            if head.borrow().right.is_some() {
                queues.push_back(head.borrow_mut().right.take());
            }
        }

        ans.push(level_res);
    }

    ans
}
