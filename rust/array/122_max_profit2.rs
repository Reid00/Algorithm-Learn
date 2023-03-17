

pub fn max_profit2(prices: Vec<i32>) -> i32 {

    let mut profit = 0;
    // 非iter 写法
    // for i in 1..prices.len() {
    //     if prices[i] > prices[i-1] {
    //         profit += prices[i] - prices[i-1];
    //     }
    // }
    // profit
    

    // iteration 写法
    // prices.windows(2).map(|x| (x[1]-x[0]).max(0)).sum()
    prices.windows(2).map(|x| x[1]-x[0]).filter(|&x| x> 0).sum()
}