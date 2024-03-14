use super::ListNode;

// 合并 K 个升序链表
pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
    if lists.is_empty() {
        return None;
    }
    // 归并排序
    merge(&lists, 0, lists.len() - 1)
}

fn merge(lists: &Vec<Option<Box<ListNode>>>, l: usize, r: usize) -> Option<Box<ListNode>> {
    if l == r {
        return lists[l].clone();
    }
    if l > r {
        return None;
    }

    let mid = l + ((r - l) >> 1);

    return merge_two_lists(merge(lists, l, mid), merge(lists, mid + 1, r));
}

fn merge_two_lists(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    match (l1, l2) {
        (Some(n1), None) | (None, Some(n1)) => Some(n1),
        (Some(mut n1), Some(mut n2)) => {
            if n1.val < n2.val {
                n1.next = merge_two_lists(n1.next.take(), Some(n2));
                Some(n1)
            } else {
                n2.next = merge_two_lists(Some(n1), n2.next.take());
                Some(n2)
            }
        }
        (None, None) => None,
    }
}
