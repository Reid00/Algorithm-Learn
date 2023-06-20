


// postorder_traversal 后序遍历的迭代实现 (左右根)
pub fn postorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {

    // 先访问左子树，再右子树，都访问完毕后即为root

    let mut ans = vec![];

    let mut stack = vec![];

    let mut node = root;

    let mut prev: Option<Rc<RefCell<TreeNode>>> = None;

    while !stack.is_empty() || node.is_some() {

        // 先访问左孩子
        while let Some(n) = node {
          node = n.borrow_mut().left.take();
          stack.push(n);
        }

        // 左孩子访问完之后，访问右孩子
        if let Some(n) = stack.pop() {
            if n.borrow().right.is_none() || n.borrow().right == prev {
              ans.push(n.borrow().val);
              prev = Some(n);
              
            }else {
              node = n.borrow_mut().right.take();
              stack.push(n);
            }
        }

    }


    ans
}



// postorder_traversal 后序遍历的递归实现 (左右根)
pub fn postorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
  
    traversal(&root)
  
  }
  
fn traversal(root: &Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
  
    if let Some(n) =  root {
  
          let mut left = traversal(&n.borrow().left);
          let mut right = traversal(&n.borrow().right);
          let mut r = vec![n.borrow().val];
          
          left.append(&mut right);
          left.append(&mut r);
          left
    }else {
      vec![]
    }
  
}