



// invertTree 翻转二叉树  递归
pub fn invert_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
    match root {
        None => None,
        Some(root) => {
            let left = invert_tree(root.borrow_mut().left.take());
            let right = invert_tree(root.borrow_mut().right.take());

            root.borrow_mut().left = right;
            root.borrow_mut().right = left;
            Some(root)
        }
    }

}