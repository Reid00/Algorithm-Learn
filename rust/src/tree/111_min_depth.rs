

// min_depth BFS
pub fn min_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {

    if root.is_none() {
        return 0
    }

    let mut depth = 1;
    use std::collections::VecDeque;
    let mut queue = VecDeque::new();
    queue.push_back(root);

    while !queue.is_empty() {

        for _ in 0..queue.len() {

            let node = queue.pop_front().unwrap().unwrap();

            // 如果node 是叶子节点，返回
            if node.borrow().left.is_none() && node.borrow().right.is_none() {
                return depth;
            }
    
            if node.borrow().left.is_some() {
                queue.push_back(node.borrow_mut().left.take());
            }
    
            if node.borrow().right.is_some(){
                queue.push_back(node.borrow_mut().right.take());
            }
        }

        // 遍历完一层，depth + 1
        depth += 1;
    }

    depth
}


// min_depth DFS
pub fn min_depth2(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {

    dfs(&root)
}

fn dfs(root: &Option<Rc<RefCell<TreeNode>>>) -> i32 {

    if root.is_none() {
        return 0
    }

    let mut depth = i32::MAX;

    if let Some(root) = root {
        if root.borrow().left.is_none() && root.borrow().right.is_none() {
            return 1
        }

        if root.borrow().left.is_some() {
            depth = depth.min(dfs(&root.borrow().left));
        }

        if root.borrow().right.is_some() {
            depth = depth.min(dfs(&root.borrow().right));
        }

        depth + 1
    }else{
        0
    }


}