use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

pub fn sorted_array_to_bst(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
    helper(&nums, 0, nums.len() as i32 - 1)
}

fn helper(nums: &Vec<i32>, left: i32, right: i32) -> Option<Rc<RefCell<TreeNode>>> {
    if left > right {
        return None;
    }

    let mid = left + ((right - left) >> 1);

    let root = Rc::new(RefCell::new(TreeNode::new(mid)));

    root.borrow_mut().left = helper(nums, left, mid - 1);
    root.borrow_mut().right = helper(nums, mid + 1, right);

    Some(root)
}
