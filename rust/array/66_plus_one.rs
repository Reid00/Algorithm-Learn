

pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {

    let mut result = digits;

    for i in (0..result.len()).rev() {

        println!("i {}", i);

        let sum = result[i] + 1;

        result[i] = sum % 10;
        // println!("res {:?}", result);
        let carry = sum / 10; 

        if carry == 0 {
            // 停止循环，把digit[:i] 赋值给result
            break
        }
        if i == 0 && carry > 0 {
            result.insert(0, carry);
        }

    }
    result
}