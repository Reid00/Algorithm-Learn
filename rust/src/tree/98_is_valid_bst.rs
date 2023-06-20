


// is_valid_bst 验证二叉树, 中序遍历
pub fn is_valid_bst(root: Option<Rc<RefCell<TreeNode>>>) -> bool {

    if root.is_none() {
        return true
    }
    // 中序遍历，迭代实现
    let mut stack = vec![];
    let mut node = root;
    let mut prev = i64::MIN;

    while !stack.is_empty() || node.is_some() {

        // 先遍历左子树
        while let Some(n) = node {
            node = n.borrow_mut().left.take();
            stack.push(n);
        }

        // 根节点，BST 的特征 根节点 >= 左子树，<= 右子树，否则不是BST
        if let Some(popped) = stack.pop(){
            if popped.borrow().val as i64 <= prev {
                return false;
            }
            prev = popped.borrow().val as i64;
            node = popped.borrow_mut().right.take();
        }
    }
    true
}


// is_valid_bst 验证二叉树， DFS
pub fn is_valid_bst2(root: Option<Rc<RefCell<TreeNode>>>) -> bool {

    helper(&root, i64::MIN, i64::MAX)

}

fn helper(root: &Option<Rc<RefCell<TreeNode>>>, lower: i64, upper: i64) -> bool {

    if root.is_none() {
        return true
    }

    // BST 的性质，根节点大于左子树，小于右子树
    if let Some(n) = root {
        if n.borrow().val as i64 <= lower || n.borrow().val as i64 >= upper {
            return false
        }

        let v = n.borrow().val as i64;
        return helper(&n.borrow().left, lower, v)
        && helper(&n.borrow().right, v, upper) 
    }
    
    true
}