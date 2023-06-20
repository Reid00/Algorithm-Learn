


// 因为二叉搜索树和中序遍历的性质，所以二叉搜索树的中序遍历是按照键增加的顺序进行的。
// 于是，我们可以通过中序遍历找到第 k 个最小元素
pub fn kth_smallest(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i32 {

    if root.is_none() {
        return 0
    }
    let mut k = k;

    // 中序遍历，迭代实现
    let mut stack = vec![];

    let mut node =root;

    while !stack.is_empty() || node.is_some() {

        // 中序遍历 左根右
        // 先遍历左子树
        while let Some(n) = node {
            node = n.borrow_mut().left.take();
            stack.push(n);
        }

        // 遍历根节点 此时node 是没有左子树的根节点
        let popped = stack.pop().unwrap();
        k -= 1;
        if k == 0 {
            return popped.borrow().val;
        }

        // 遍历右子树
        node = popped.borrow_mut().right.take();
    }
    0
}

// kth_smallest BFS + 排序
pub fn kth_smallest2(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i32 {

    if root.is_none() {
        return 0
    }

    let mut data = vec![];

    use std::collections::VecDeque;

    let mut queue = VecDeque::new();
    queue.push_back(root);

    while !queue.is_empty() {

        (0..queue.len()).for_each(|_| {

            let node = queue.pop_front().unwrap().unwrap();

            data.push(node.borrow().val);

            if node.borrow().left.is_some() {
                queue.push_back(node.borrow_mut().left.take());
            }

            if node.borrow().right.is_some() {
                queue.push_back(node.borrow_mut().right.take());
            }

        });

    }

    data.sort();
    data[(k-1) as usize]
}