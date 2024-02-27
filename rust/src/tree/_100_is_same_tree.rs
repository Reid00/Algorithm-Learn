use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

/// 相同的树
pub fn is_same_tree(p: Option<Rc<RefCell<TreeNode>>>, q: Option<Rc<RefCell<TreeNode>>>) -> bool {
    match (p, q) {
        (Some(n1), Some(n2)) => {
            n1.borrow().val == n2.borrow().val
                && is_same_tree(n1.borrow().left.clone(), n2.borrow().left.clone())
                && is_same_tree(n1.borrow().right.clone(), n2.borrow().right.clone())
        }
        (None, None) => true,
        _ => false,
    }
}
