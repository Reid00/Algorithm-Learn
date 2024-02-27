use rust_leetcode::tree;

// 下面error 因为main.rs 和 lib.rs 或者其他bin 下面的不同的二进制包中
// crate 代表不同的文件。换句话说，main.rs 中用crate 这个crate root 从main.rs 中去找，不会从lib.rs中
// use crate::tree;
use std::cell::RefCell;
use std::rc::Rc;

fn main() {
    let root = Some(Rc::new(RefCell::new(tree::TreeNode::new(2))));
    tree::inorder_traversal(root);
    println!("hello 2");
}
