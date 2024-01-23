


pub fn get_least_numbers(arr: Vec<i32>, k: i32) -> Vec<i32> {

    if arr.len() == 0 || k == 0 {
        return vec![]
    }

    use std::collections::BinaryHeap;
    use std::cmp::Reverse;  // covert to MiniHeap

    // MaxHeap 不符合要求
    let heap = BinaryHeap::from(arr); 

    // convert to MiniHeap
    let mut heap: BinaryHeap<Reverse<i32>> = heap.into_iter().map(Reverse).collect();


    let mut ans = Vec::new();

    for _ in 0..k {
        let v = heap.pop().unwrap().0;
        ans.push(v);
    }
    ans
}

