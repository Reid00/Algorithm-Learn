use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

// 迭代实现
pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
    let mut ans = vec![];
    let mut stack = vec![];

    let mut node = root;

    while !stack.is_empty() || node.is_some() {
        // 中序遍历，将左孩子压栈
        while let Some(n) = node {
            // 取出左孩子
            node = n.borrow_mut().left.take();
            // 把当前节点压栈，后续出栈 取右孩子
            stack.push(n);
        }

        // 当node 为None时，意味着stack 中最后一个元素为最左的孩子
        if let Some(n) = stack.pop() {
            ans.push(n.borrow().val);
            node = n.borrow_mut().right.take();
        }
    }

    ans
}

// inorder_traversal 递归实现
pub fn inorder_traversal2(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
    let ans = inorder(&root);

    ans
}

fn inorder(root: &Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
    if let Some(node) = root {
        let mut left = inorder(&node.borrow().left);

        let mut r = vec![node.borrow().val];

        let mut right = inorder(&node.borrow().right);

        left.append(&mut r);
        left.append(&mut right);
        left
    } else {
        vec![]
    }
}
