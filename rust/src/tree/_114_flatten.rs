use super::TreeNode;
use std::cell::RefCell;
use std::rc::Rc;

pub fn flatten(root: &mut Option<Rc<RefCell<TreeNode>>>) {
    // O(1)
    // 1. 寻找当前节点左子树最右边的节点作为前驱节点(predecessor)
    // 2. 讲当前节点的右子树作为predecessor 的右子树
    // 3. 讲当前节点的左子树(next)更新为右子树，左子树置空
    // 4. 循环cur = cur.Left

    let mut cur = root.as_ref().cloned();

    while let Some(c) = cur {
        // NOTE 此处要先声明一个可变的node， 不能后面一直用c.borrow_mut() 会产生太多可变引用，导致运行时出错
        let mut cur_node = c.borrow_mut();

        // 左子树存在，寻找左子树最右侧的节点
        if let Some(left) = cur_node.left.take() {
            let next = left.clone();
            let mut pre = left.clone();

            let mut pre_right = left.borrow().right.clone();

            while let Some(inner_pre) = pre_right {
                pre_right = inner_pre.borrow().right.clone();
                pre = inner_pre;
            }

            pre.borrow_mut().right = cur_node.right.clone();
            // pre.borrow_mut().right = cur_node.right.take();

            // left take 会直接置为None
            // cur_node.left = None;
            // 当前左子树更新为右子树
            cur_node.right = Some(next);
        }

        cur = cur_node.right.clone();
    }
}
