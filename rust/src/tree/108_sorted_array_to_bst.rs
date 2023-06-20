

// sorted_array_to_bst 中序遍历
pub fn sorted_array_to_bst(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {

    let nums = &nums[..];

    builder(nums, 0,  (nums.len() -1) as i32)
}

fn builder(nums: &[i32], left: i32, right: i32) -> Option<Rc<RefCell<TreeNode>>> {

    if left > right {
        return None
    }

    // BUG 1 + (1-0) >>1
    // let mid = left + (right - left) >> 1;
    //  FIX ABOVE
    let mid  = left + ((right - left)>>1);
    // or below
    // let mid = (right + left) >> 1;


    let root = Some(
        Rc::new(RefCell::new(
            TreeNode{
                val: nums[mid as usize],
                left: None,
                right: None,
            }
        ))
    );

    if let Some(node) = &root {
        // left 用i32 not usize 是为了防止mid = 0 , 0-1 越界
        node.borrow_mut().left = builder(nums, left, mid -1);
        node.borrow_mut().right = builder(nums, mid + 1, right); 
    };

    root
}