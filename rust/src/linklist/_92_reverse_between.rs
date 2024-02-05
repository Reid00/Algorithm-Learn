use super::ListNode;

/// 反转链表 II
pub fn reverse_between(
    head: Option<Box<ListNode>>,
    left: i32,
    right: i32,
) -> Option<Box<ListNode>> {
    if head.is_none() {
        return None;
    }

    if left == right {
        return head;
    }
    // left 2, right  4
    // A -> B -> C -> D -> E
    // 1. A -> C -> B -> D - > E
    // 2. A -> D -> C -> B - > E

    // 穿针引线法  一次遍历
    let mut dummy = Some(Box::new(ListNode::new(0)));
    dummy.as_mut().unwrap().next = head;

    let mut prev = &mut dummy;

    for _ in 0..left - 1 {
        prev = &mut prev.as_mut().unwrap().next;
    }

    let mut cur = prev.as_mut().unwrap().next.take();

    for _ in 0..right - left {
        let mut next = &mut cur.as_mut().unwrap().next.take();

        cur.as_mut().unwrap().next = next.as_mut().unwrap().next.take();
        // 此处将next 插入到prev 后面，但是prev 在 上面cur 处被切断，导致链表 以cur 为中心分成了两个部分
        next.as_mut().unwrap().next = prev.as_mut().unwrap().next.take();
        prev.as_mut().unwrap().next = next.take();
    }
    // 上述遍历完形成两个链表  A->D->C + B->E
    // 将prev 移到C
    while prev.as_mut().unwrap().next.is_some() {
        prev = &mut prev.as_mut()?.next;
    }
    // 重新连接上
    prev.as_mut().unwrap().next = cur.take();

    dummy.unwrap().next
}
