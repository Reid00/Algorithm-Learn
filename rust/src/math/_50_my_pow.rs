// Pow(x, n)
pub fn my_pow(x: f64, n: i32) -> f64 {
    // 快速幂法
    if n >= 0 {
        return quick_mu(x, n);
    }

    1.0 / quick_mu(x, -n)
}

pub fn quick_mu(x: f64, n: i32) -> f64 {
    if n == 0 {
        return 1.0;
    }

    let y = quick_mu(x, n / 2);

    if n % 2 == 0 {
        return y * y;
    }
    y * y * x
}
