

// has_path_sum 路径总和  某个路径上是否存在到叶子节点总合等于targetSum 的路径
// BFS  加入到queue中的时候，同时把到此点的路径和计算出来
pub fn has_path_sum(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> bool {

    if root.is_none() {
        return false
    }

    use std::collections::VecDeque;
    let mut queue = VecDeque::new();
    // queue 中添加 元组(root, path.sum)  当前节点和到此节点的路径和
    // let v = root.as_ref().unwrap().borrow().val;
    // 更加ruster 的写法
    let v = root.as_ref().map(|node| node.borrow().val).unwrap();
    queue.push_back((root, v));

    while !queue.is_empty() {

        
        for _ in 0..queue.len() {
            if let Some((node, path_sum)) = queue.pop_front(){
                if let Some(n) = node {
                    // 左右子树为None 才为叶子节点
                    if n.borrow().left.is_none() && n.borrow().right.is_none() && path_sum == target_sum {
                        return true
                    }
                    
                    println!("n and path_sum: {}, {}", n.borrow().val, path_sum);

                    if n.borrow().left.is_some() {
                        let left = n.borrow_mut().left.take();
                        let v = left.as_ref().unwrap().borrow().val;
                        queue.push_back((left, path_sum + v));
                    }

                    if n.borrow().right.is_some() {
                        let right = n.borrow_mut().right.take();
                        let v = right.as_ref().unwrap().borrow().val;
                        queue.push_back((right, path_sum + v));
                    }

                }
            }
        }
    }
    false
}


// has_path_sum 路径总和  某个路径上是否存在到叶子节点总合等于targetSum 的路径
// DFS
pub fn has_path_sum2(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> bool {

    has_sum(&root, target_sum)
  
}

fn has_sum(root: &Option<Rc<RefCell<TreeNode>>>, target: i32) -> bool {

    if root.is_none() {
        return false
    }

    if let Some(root) = root {
        
        if root.borrow().left.is_none() && root.borrow().right.is_none() && root.borrow().val == target {
            return true
        }
        
        // 剩余路径的和为target - root.val
        return has_sum(&root.borrow().left, target - root.borrow().val) || 
        has_sum(&root.borrow().right, target - root.borrow().val)
    }
    return false
}