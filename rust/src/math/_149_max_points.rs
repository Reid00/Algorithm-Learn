/// 直线上最多的点数
pub fn max_points(points: Vec<Vec<i32>>) -> i32 {
    let mut res = 0;

    for i in 0..points.len() {
        use std::collections::HashMap;
        let mut hmap = HashMap::new();

        for j in 0..points.len() {
            if i != j {
                hmap.entry(line_scop(&points[i], &points[j]))
                    .and_modify(|x| *x += 1)
                    .or_insert(1);
            }
        }

        // 遍历完之后，比较哪个值大
        for (_, v) in hmap {
            if v > res {
                res = v
            }
        }
    }

    res + 1
}

fn line_scop(a: &[i32], b: &[i32]) -> (i32, i32) {
    let dx = a[0] - b[0];
    let dy = a[1] - b[1];

    if dx == 0 {
        return (0, 1); // 垂直线，表示斜率为正无穷大
    }

    let gcd_val = gcd(dx.abs(), dy.abs());
    ((dx / gcd_val), dy / gcd_val)
}

fn gcd(mut a: i32, mut b: i32) -> i32 {
    while b != 0 {
        let temp = b;
        b = a % b;
        a = temp;
    }
    a
}
