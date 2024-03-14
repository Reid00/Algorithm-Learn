use super::ListNode;

/// 排序链表
pub fn sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    // 分治： 拆分 + 归并排序

    let mut head = head;

    if head.is_none() || head.as_ref().unwrap().next.is_none() {
        return head;
    }

    let mut len = 0;
    let mut pre = &head;
    while let Some(node) = pre {
        len += 1;
        pre = &node.next;
    }

    let mut pre = &mut head;

    for _ in 0..len / 2 {
        if let Some(ref mut node) = pre {
            pre = &mut node.next;
        }
    }

    let next = pre.take();
    let l1 = sort_list(head);
    let l2 = sort_list(next);

    merge(l1, l2)
}

fn merge(mut l: Option<Box<ListNode>>, mut r: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    match (l, r) {
        (None, Some(n)) | (Some(n), None) => Some(n),
        (Some(mut n1), Some(mut n2)) => {
            if n1.val < n2.val {
                n1.next = merge(n1.next.take(), Some(n2));
                Some(n1)
            } else {
                n2.next = merge(Some(n1), n2.next.take());
                Some(n2)
            }
        }
        (None, None) => None,
    }
}
