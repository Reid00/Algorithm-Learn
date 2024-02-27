use super::ListNode;

/// K 个一组翻转链表
pub fn reverse_k_group(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
    let mut remain = head;
    let mut dummy = Box::new(ListNode::new(0));
    let mut tail = &mut dummy;

    while remain.is_some() {
        let (new_head, new_remain) = reverse_one(remain, k);

        remain = new_remain;
        tail.next = new_head;
        while tail.next.as_ref().is_some() {
            tail = tail.next.as_mut()?;
        }
    }

    dummy.next
}

/// 反转一次，返回反转后的head 和 remain
/// 如果最后一次不足以反转，remain 为None
fn reverse_one(
    mut head: Option<Box<ListNode>>,
    k: i32,
) -> (Option<Box<ListNode>>, Option<Box<ListNode>>) {
    let mut prev = head.as_mut();

    for _ in 0..k {
        // 因为从head 开始，k 次循环之后，prev 最后指向的是下一段k 的head
        // 如果prev 是None 说明长度不满k, 可以返回
        if prev.is_none() {
            return (head, None);
        }

        prev = prev.unwrap().next.as_mut();
    }

    // 反转之后上一段的tail，拥有下一段的所有权
    let mut cur = head;
    let mut dummy = ListNode::new(0);

    for _ in 0..k {
        if let Some(mut n) = cur {
            // 从下一个断开
            cur = n.next.take();
            // 重要
            n.next = dummy.next.take();

            dummy.next = Some(n);
        }
    }

    (dummy.next, cur)
}
