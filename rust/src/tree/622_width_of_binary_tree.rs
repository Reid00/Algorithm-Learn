
// width_of_binary_tree 最大深度，BFS
pub fn width_of_binary_tree(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {

    if root.is_none() {
        return 0
    }

    let mut width = 1;

    use std::collections::VecDeque;
    let mut queue = VecDeque::new();

    // queue 中存储的是元组，idx of node 和 node
    // 最大宽度则是计算每一层的queue[last].idx - queue[0].idx  + 1
    queue.push_back((0, root));

    while !queue.is_empty() {

        let first = queue.front().unwrap();
        let last = queue.back().unwrap();
        width = width.max(last.0 - first.0);

        for _ in 0..queue.len() {

            let head = queue.pop_front().unwrap();
            
            let idx = head.0;
            let node = head.1.unwrap();

            if node.borrow().left.is_some() {
                queue.push_back((idx * 2 + 1, node.borrow_mut().left.take()));
            }

            if node.borrow().right.is_some() {
                queue.push_back((idx * 2 + 2, node.borrow_mut().right.take()));
            }
        }

    }
    width
}