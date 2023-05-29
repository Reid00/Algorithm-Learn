

pub fn max_profit(prices: Vec<i32>) -> i32 {

    let mut min_price = std::i32::MAX;
    let mut max_profit = 0;

    for &v in &prices{
        if v < min_price {
            min_price = v;
        }else if v - min_price > max_profit {
            max_profit = v - min_price;
        }
    }
    max_profit
}

// max_profit2 迭代器用法
pub fn max_profit2(prices: Vec<i32>) -> i32 {

    let mut min_price = std::i32::MAX;
    let mut max_profit = 0;

    prices.iter().skip(0).for_each(|&v|{
        min_price = min_price.min(v);
        max_profit = max_profit.max(v-min_price);
    });

    max_profit
   
}