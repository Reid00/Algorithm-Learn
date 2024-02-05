mod _206_reverse_list;
mod _21_merge_two_lists;
mod _25_reverse_k_group;
mod _2_add_two_num;
mod _92_reverse_between;

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}
