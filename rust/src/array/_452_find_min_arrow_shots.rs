/// 用最少数量的箭引爆气球
pub fn find_min_arrow_shots(points: Vec<Vec<i32>>) -> i32 {
    if points.len() == 0 {
        return 0;
    }

    let mut points = points;

    // 安装points 元素中的x[1] 进行排序，它决定了箭能在的最大右侧位置
    points.sort_by_key(|x| x[1]);
    // points.sort_by(|a, b| a[1].cmp(&b[1]));

    let mut max_right = points[0][1];
    let mut res = 1;

    for point in points.iter() {
        if point[0] > max_right {
            max_right = point[1];
            res += 1;
        }
    }
    res
}
