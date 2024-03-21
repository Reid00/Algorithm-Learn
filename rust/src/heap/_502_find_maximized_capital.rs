use std::collections::BinaryHeap;

/// IPO
/// 根据贪心的策略，我们要想总的资本最大，那么每一次都要在能启动的项目中选择利润最大的那个。
/// 因此我们可以使用一个大顶堆来维护当前所有能启动的项目，堆顶项目即为利润最大的那个。
pub fn find_maximized_capital(k: i32, w: i32, profits: Vec<i32>, capital: Vec<i32>) -> i32 {
    let n = profits.len();
    // 创建利润的大根堆，每次选择利润最大的
    let mut h = BinaryHeap::new();

    // 利润和资本的pair 对， 防止按照资本排序后找不到对应的利润
    // let mut pairs = vec![];
    // for item in profits.into_iter().zip(capital) {
    //     pairs.push(item);
    // }
    let mut pairs: Vec<(i32, i32)> = profits.into_iter().zip(capital).collect();

    // 对项目所需资本排序
    pairs.sort_by(|a, b| a.1.cmp(&b.1));

    let mut w = w;

    // 选择到哪个资本了
    let mut i = 0;
    // 选择次数k
    let mut k = k;
    while k > 0 {
        // 贪心: 把所有可以选择的项目放到大根堆中，然后选择堆顶(利润最大的项目)
        while i < n && w >= pairs[i].1 {
            h.push(pairs[i].0);
            i += 1;
        }
        // 没有项目可以选择
        if h.is_empty() {
            break;
        }

        k -= 1;
        w += h.pop().unwrap();
    }
    w
}
