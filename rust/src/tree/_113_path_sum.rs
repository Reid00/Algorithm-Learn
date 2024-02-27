use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

/// BFS 每次都要保留路径的clone, 太耗内存
pub fn path_sum_(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> Vec<Vec<i32>> {
    let mut result = vec![];
    if root.is_none() {
        return result;
    }

    use std::collections::VecDeque;
    let mut queue = VecDeque::new();

    let v = root.as_ref().map(|x| x.borrow().val).unwrap();
    let mut path = vec![];
    path.push(v);
    // (当前节点，到当前节点的路径和, 路径)
    queue.push_back((root, v, path.clone()));

    while !queue.is_empty() {
        // 层次遍历，队列先进先出
        for _ in 0..queue.len() {
            // 出队列，node 和其所在path 和
            if let Some((n, path_sum, mut path)) = queue.pop_front() {
                if let Some(node) = n {
                    if path_sum == target_sum
                        && node.borrow().left.is_none()
                        && node.borrow().right.is_none()
                    {
                        result.push(path.clone());
                    }

                    if node.borrow().left.is_some() {
                        let left_v = node.borrow().left.as_ref().map(|x| x.borrow().val).unwrap();
                        path.push(left_v);
                        queue.push_back((
                            node.borrow().left.clone(),
                            left_v + path_sum,
                            path.clone(),
                        ));
                        // path 要弹出最后加入的元素，避免影响right
                        path.pop();
                    }

                    if node.borrow().right.is_some() {
                        let right_v = node
                            .borrow()
                            .right
                            .as_ref()
                            .map(|x| x.borrow().val)
                            .unwrap();
                        path.push(right_v);
                        queue.push_back((
                            node.borrow().right.clone(),
                            right_v + path_sum,
                            path.clone(),
                        ));
                        // path 要弹出最后加入的元素
                        path.pop();
                    }
                }
            }
        }
    }

    result
}
