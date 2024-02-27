use std::cell::RefCell;
use std::rc::Rc;

mod _100_is_same_tree;
mod _101_is_symmetric;
mod _102_level_order;
mod _104_max_depth;
mod _105_build_tree;
mod _106_build_tree;
mod _112_has_path_sum;
mod _113_path_sum;
mod _114_flatten;
mod _144_preorder_traversal;
mod _145_postorder_traversal;
mod _226_invert_tree;
mod _94_inorder_traversal;

pub use _94_inorder_traversal::inorder_traversal;

// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}
