use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

/// 从中序与后序遍历序列构造二叉树
pub fn build_tree(inorder: Vec<i32>, postorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
    build(&inorder, &postorder)
}

fn build(inorder: &[i32], postorder: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
    if inorder.is_empty() {
        return None;
    }

    let root_val = postorder.last()?;
    let left_size = inorder.iter().position(|x| x == root_val)?;

    let root = Some(Rc::new(RefCell::new(TreeNode::new(*root_val))));

    if let Some(r) = &root {
        r.borrow_mut().left = build(&inorder[..left_size], &postorder[..left_size]);
        r.borrow_mut().right = build(
            &inorder[left_size + 1..],
            &postorder[left_size..postorder.len() - 1],
        );
    }
    return root;
}
