pub fn merge(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
    let mut intervals = intervals;
    // 排序
    intervals.sort();
    let mut ans = vec![];

    let (mut l, mut r) = (intervals[0][0], intervals[0][1]);

    intervals.iter().skip(1).for_each(|x| {
        // 后面的元素大于前面的元素 不需要合并
        if x[0] > r {
            ans.push(vec![l, r]);
            l = x[0]
        }
        r = r.max(x[1]);
    });
    ans.push(vec![l, r]);
    ans
}
