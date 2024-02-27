use super::ListNode;

/// 移除指定元素
pub fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
    let mut dummy = Box::new(ListNode::new(-1));
    dummy.next = head;

    let mut cur = &mut dummy;

    // take() 断开链表，获取其所有权
    while let Some(node) = cur.next.take() {
        if node.val == val {
            cur.next = node.next;
        } else {
            cur.next = Some(node);
            cur = cur.next.as_mut()?;
        }
    }

    dummy.next
}
