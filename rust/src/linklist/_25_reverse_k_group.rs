use super::ListNode;

/// K 个一组翻转链表
pub fn reverse_k_group(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
    // 哨兵
    let mut dummy = ListNode::new(0);

    let mut prev = &mut dummy;

    let mut cur = head;

    return dummy.next;
}

fn reverse(
    mut head: Option<Box<ListNode>>,
    mut tail: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {

    let mut prev = tail.as_mut()?.next;

    while let Some(cur) = head {
        let next = cur.next;

        cur.next = next;
        prev = Some(cur);
        cur = next;
    }

    tail
}
