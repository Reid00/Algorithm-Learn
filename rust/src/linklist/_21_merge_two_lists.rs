use super::ListNode;

// match 方法
/// 合并两个有序链表
pub fn merge_two_lists(
    list1: Option<Box<ListNode>>,
    list2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    // 哨兵节点
    let mut dummy = None;
    let mut cur = &mut dummy;

    let mut l1 = list1;
    let mut l2 = list2;

    *cur = loop {

        match (l1, l2) {
            (Some(mut a), Some(mut b)) => {
                if a.val <= b.val {
                    l1 = a.next.take();
                    l2 = Some(b);            
                    cur = &mut cur.insert(a).next;
                } else {
                    l1 = Some(a);
                    l2 = b.next.take();
                    cur = &mut cur.insert(b).next;
                }
            }
            (x, y) => break x.or(y),
        }

    };

    dummy
}

/// 合并两个有序链表
pub fn merge_two_lists2(
    list1: Option<Box<ListNode>>,
    list2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    // 哨兵节点
    let mut dummy = ListNode::new(0);
    let mut cur = &mut dummy;

    let mut l1 = list1;
    let mut l2 = list2;

    while l1.is_some() && l2.is_some() {
        let (v1, v2) = (l1.as_ref().unwrap().val, l2.as_ref().unwrap().val);

        // 获取l1 or l2 的head 引用
        let l = if v1 <= v2 { &mut l1 } else { &mut l2 };

        // 取走 l 的head， 置none
        cur.next = l.take();
        cur = cur.next.as_mut().unwrap();
        // 更新l
        *l = cur.next.take();
    }

    // 优化
    cur.next = l1.or(l2);
    // if l1.is_some() {
    //     cur.next = l1;
    // }

    // if l2.is_some() {
    //     cur.next = l2;
    // }

    dummy.next
}
