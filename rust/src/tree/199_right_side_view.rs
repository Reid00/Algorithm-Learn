

// BFS
pub fn right_side_view(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {

    if root.is_none() {
        return vec![];
    }

    let mut ans = vec![];

    use std::collections::VecDeque;

    let mut queue = VecDeque::new();
    queue.push_back(root);

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
        // println!("level: {:?}", level);
        ans.push(*level.last().unwrap());
    }

    ans

}