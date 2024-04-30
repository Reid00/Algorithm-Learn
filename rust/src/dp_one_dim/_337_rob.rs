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
use std::cell::RefCell;
use std::rc::Rc;

/// 打家劫舍 III
/// 房屋邻接是树形，不是数组，所以需要dfs
pub fn rob(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    // 遍历每个节点的时候 都有两种选择，偷当前节点和不偷当前节点，所以用个长度为2的数组dp即可
    // 下标为0记录不偷该节点所得到的的最大金钱，下标为1记录偷该节点所得到的的最大金钱。
    // 偷当前节点的时候不能偷左右孩子
    // 遍历顺序, 用后序遍历
    // 通过递归左节点，得到左节点偷与不偷的金钱。	left = robTree(root.Left)
    // 	通过递归右节点，得到右节点偷与不偷的金钱。  right = robTree(root.Right)
    // 如果是偷当前节点，那么左右孩子就不能偷，val1 = cur->val + left[0] + right[0];
    // 如果不偷当前节点，那么左右孩子就可以偷，至于到底偷不偷一定是选一个最大的，
    // 所以：val2 = max(left[0], left[1]) + max(right[0], right[1]);

    let res = rob_tree(root);
    res.0.max(res.1)
}

pub fn rob_tree(root: Option<Rc<RefCell<TreeNode>>>) -> (i32, i32) {
    match root {
        None => (0, 0),
        Some(node) => {
            let left = rob_tree(node.borrow().left.clone());
            let right = rob_tree(node.borrow().right.clone());

            // 偷当前节点，不能偷左右孩子
            let v1 = node.borrow().val + left.0 + right.0;
            // 不偷当前节点，偷左右孩子
            let v2 = left.0.max(left.1) + right.0.max(right.1);

            (v2, v1)
        }
    }
}
