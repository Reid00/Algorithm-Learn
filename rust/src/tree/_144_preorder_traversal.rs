use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

// 递归实现
pub fn preorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
    // rust 不允许递归调用闭包, 所以fn preorder 定义在外面
    let ans = preorder(&root);

    ans
}

pub fn preorder(root: &Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
    if let Some(node) = root {
        // 根节点
        let mut r = vec![node.borrow().val];
        // 递归调用左孩子
        let mut r_left = preorder(&node.borrow().left);
        // 递归调用右孩子
        let mut r_right = preorder(&node.borrow().right);

        r.append(&mut r_left);
        r.append(&mut r_right);
        r
    } else {
        vec![]
    }
}

// 迭代实现
pub fn preorder_traversal2(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
    let mut ans = vec![];
    let mut stack = vec![];
    let mut node = root;

    while !stack.is_empty() || node.is_some() {
        // 根节点 及其左孩子压栈
        while let Some(n) = node {
            ans.push(n.borrow().val);
            node = n.borrow_mut().left.take();
            // 此处和上步不能调换，否则n 的所有权转移到stack 中
            stack.push(n);
        }

        // 出栈获取其右孩子
        if let Some(n) = stack.pop() {
            node = n.borrow_mut().right.take();
        }
    }

    ans
}
