use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

// is_symmetric 对称二叉树 BFS 实现
pub fn is_symmetric(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
    if root.is_none() {
        return true;
    }

    use std::collections::VecDeque;
    let mut queue = VecDeque::new();

    // 把root 的左右子树放到queue 中，查看左右子树是否对称(相邻的左右子树应该对称)
    if let Some(root) = root {
        queue.push_back(root.borrow_mut().left.take());
        queue.push_back(root.borrow_mut().right.take());

        // queue 中的数据出队列，查看是否相同
        while let (Some(n1), Some(n2)) = (queue.pop_front(), queue.pop_front()) {
            match (n1, n2) {
                (None, None) => (), // 都是None  跳过
                (Some(x), Some(y)) => {
                    if x.borrow().val == y.borrow().val {
                        queue.push_back(x.borrow_mut().left.take());
                        queue.push_back(y.borrow_mut().right.take());
                        queue.push_back(x.borrow_mut().right.take());
                        queue.push_back(y.borrow_mut().left.take());
                    } else {
                        return false;
                    }
                }
                _ => return false,
            }
        }
    }

    true
}

// is_symmetric 对称二叉树 DFS 实现
pub fn is_symmetric2(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
    if root.is_none() {
        return true;
    }

    if let Some(n) = root {
        return check(&n.borrow().left, &n.borrow().right);
    }
    true
}

fn check(left: &Option<Rc<RefCell<TreeNode>>>, right: &Option<Rc<RefCell<TreeNode>>>) -> bool {
    if left.is_none() && right.is_none() {
        return true;
    }
    // left, right 其中一个为None 另一个不为None
    if left.is_none() || right.is_none() {
        return false;
    }

    if let (Some(l), Some(r)) = (left, right) {
        return l.borrow().val == r.borrow().val
            && check(&l.borrow().left, &r.borrow().right)
            && check(&l.borrow().right, &r.borrow().left);
    }
    true
}

// is_symmetric 对称二叉树 DFS 实现
pub fn is_symmetric3(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
    if root.is_none() {
        return true;
    }

    if let Some(root) = root {
        return check_val(root.borrow().left.clone(), root.borrow().right.clone());
    }
    false
}

fn check_val(left: Option<Rc<RefCell<TreeNode>>>, right: Option<Rc<RefCell<TreeNode>>>) -> bool {
    match (left, right) {
        (Some(l), Some(r)) => {
            l.borrow().val == r.borrow().val
                && check_val(l.borrow().left.clone(), l.borrow().right.clone())
                && check_val(l.borrow().right.clone(), l.borrow().left.clone())
        }

        (None, None) => true,
        _ => false,
    }
}
