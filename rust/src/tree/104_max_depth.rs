

// BFS
pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if root.is_none() {
        return 0
    }


    let mut ans = 0;

    use std::collections::VecDeque;
    let mut queue = VecDeque::new();

    queue.push_back(root);

    while !queue.is_empty() {

        (0..queue.len()).for_each(|_| {

            let head = queue.pop_front().unwrap().unwrap();
            
            if head.borrow().left.is_some() {
                queue.push_back(head.borrow_mut().left.take());
            }
            if head.borrow().right.is_some() {
                queue.push_back(head.borrow_mut().right.take());
            }
            // if let Some(n) = head.borrow_mut().left.take() {
            //     queue.push_back(Some(n));
            // };

            // if let Some(n) = head.borrow_mut().right.take() {
            //     queue.push_back(Some(n));
            // };
        });

        ans +=1;
    }

    ans
}

// DFS
pub fn max_depth2w(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if root.is_none() {
        return 0
    }

    dfs(&root)

}

fn dfs(root: &Option<Rc<RefCell<TreeNode>>>) -> i32 {

    if let Some(node) = root {
        
        let left_depth = dfs(&node.borrow().left);
        let right_depth = dfs(&node.borrow().right);

        left_depth.max(right_depth) + 1
    }else {
        0
    }

}