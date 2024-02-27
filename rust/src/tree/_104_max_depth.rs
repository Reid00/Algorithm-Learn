use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

// BFS
pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if root.is_none() {
        return 0;
    }

    let mut depth = 0;

    use std::collections::VecDeque;

    let mut q = VecDeque::new();
    q.push_back(root);

    while !q.is_empty() {
        let n = q.len();

        for _ in 0..n {
            if let Some(Some(node)) = q.pop_front() {
                if node.borrow().left.is_some() {
                    q.push_back(node.borrow().left.clone());
                }

                if node.borrow().right.is_some() {
                    q.push_back(node.borrow().right.clone());
                }
            }
        }
        depth += 1;
    }
    depth
}

// DFS
pub fn max_depth2(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if root.is_none() {
        return 0;
    }

    dfs(root)
}

fn dfs(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if let Some(node) = root {
        let left_depth = dfs(node.borrow().left.clone());
        let right_depth = dfs(node.borrow().right.clone());

        left_depth.max(right_depth) + 1
    } else {
        0
    }
}

// DFS + match
pub fn max_depth3(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    match root {
        Some(n) => {
            let l_depth = max_depth3(n.borrow().left.clone());
            let r_depth = max_depth3(n.borrow().right.clone());
            l_depth.max(r_depth) + 1
        }
        _ => 0,
    }
}
