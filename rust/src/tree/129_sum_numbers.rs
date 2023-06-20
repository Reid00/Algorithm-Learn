


// sum_numbers 求根节点到叶节点数字之和 
// DFS
pub fn sum_numbers(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    add_prev(&root, 0)
}

fn add_prev(root: &Option<Rc<RefCell<TreeNode>>>, prev: i32) -> i32 {

    if root.is_none() {
        return 0
    }
    
    if let Some(n) = root {
        let tmp = n.borrow().val + prev * 10;

        // 到叶子节点，必须保证n 的左右子树为none
        if n.borrow().left.is_none() && n.borrow().right.is_none() {
            return tmp
        }

        return add_prev(&n.borrow().left, tmp) + 
        add_prev(&n.borrow().right, tmp)
    }

    0
}

// sum_numbers 求根节点到叶节点数字之和 
// BFS 每次加入queue 的数据是node 和 父节点的路径和 组成的元组
pub fn sum_numbers(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {

    if root.is_none() {
        return 0;
    }

    let mut result = 0;

    use std::collections::VecDeque;
    let mut queue = VecDeque::new();
    queue.push_back((root, 0));


    while !queue.is_empty() {


        (0..queue.len()).for_each(|_|{

            let head = queue.pop_front().unwrap();

            if let (Some(node), prev_sum) = head {
                let tmp = prev_sum * 10 + node.borrow().val;

                // 当node 是叶子节点时，将结果add to result
                if node.borrow().left.is_none() && node.borrow().right.is_none() {
                    result += tmp;
                }

                if node.borrow().left.is_some() {
                    queue.push_back((node.borrow_mut().left.take(), tmp));
                }

                if node.borrow().right.is_some() {
                    queue.push_back((node.borrow_mut().right.take(), tmp));
                }

            }


        })
        
    }

    result
}