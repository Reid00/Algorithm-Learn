use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

// build_tree 从前序与中序遍历序列构造二叉树
// 前序 跟左右  => 第一个是root
// 中序 左根右  => 根据上面的root确定其左边是左子树
pub fn build_tree(preorder: Vec<i32>, inorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
    builder(&preorder, &inorder)
}

fn builder(preorder: &[i32], inorder: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
    if preorder.len() == 0 || inorder.len() == 0 {
        return None;
    }

    let root_val = preorder[0];

    let root = Some(Rc::new(RefCell::new(TreeNode {
        val: root_val,
        left: None,
        right: None,
    })));

    // root 在中序遍历的下标
    let idx = inorder.iter().position(|&x| x == root_val).unwrap();

    if let Some(n) = &root {
        n.borrow_mut().left = builder(&preorder[1..idx + 1], &inorder[..idx]);

        n.borrow_mut().right = builder(&preorder[idx + 1..], &inorder[idx + 1..])
    }
    root
}

/// 迭代的方式实现
/// Error  参考Go 的105 后续有时间实现
#[allow(unreachable_code)]
pub fn build_tree2(preorder: Vec<i32>, inorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
    if preorder.is_empty() {
        return None;
    }

    let root_val = preorder[0];
    let root = Some(Rc::new(RefCell::new(TreeNode::new(root_val))));

    use std::collections::VecDeque;
    let mut stack = VecDeque::new();

    stack.push_back(root.clone());

    let mut inorder_idx = 0;
    // 参考Go105 后续有时间实现
    unimplemented!();
    root
}
