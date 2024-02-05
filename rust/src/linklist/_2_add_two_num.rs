use super::ListNode;

/// 哨兵节点用Option模式
pub fn add_two_numbers(
    l1: Option<Box<ListNode>>,
    l2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    let mut l1 = l1;
    let mut l2 = l2;

    // 哨兵节点, return dummpy.next
    let mut dummy = None;
    let mut cur = &mut dummy;

    let mut carry = 0;

    while l1.is_some() || l2.is_some() || carry != 0 {
        let (mut add1, mut add2) = (0, 0);

        if let Some(v) = l1 {
            add1 = v.val;
            l1 = v.next;
        }

        if let Some(v) = l2 {
            add2 = v.val;
            l2 = v.next;
        }

        let sum = add1 + add2 + carry;
        let val = sum % 10;
        carry = sum / 10;

        cur = &mut cur.insert(Box::new(ListNode::new(val))).next;
    }

    dummy
}

pub fn add_two_numbers2(
    l1: Option<Box<ListNode>>,
    l2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    let mut l1 = l1;
    let mut l2 = l2;

    // 哨兵节点, return dummpy.next
    let mut dummy = Box::new(ListNode::new(0));

    let mut cur = &mut dummy;

    let mut carry = 0;

    while l1.is_some() || l2.is_some() || carry != 0 {
        let (mut add1, mut add2) = (0, 0);

        if let Some(v) = l1 {
            add1 = v.val;
            l1 = v.next;
        }

        if let Some(v) = l2 {
            add2 = v.val;
            l2 = v.next;
        }

        let sum = add1 + add2 + carry;
        let val = sum % 10;
        carry = sum / 10;

        // 第一次cur.next 就是 == dummy.next
        cur.next = Some(Box::new(ListNode::new(val)));
        cur = cur.next.as_mut().unwrap();
    }

    dummy.next
}
