pub fn majority_element_1(nums: Vec<i32>) -> Vec<i32> {

    let (mut element1, mut vote1) = (0, 0);
    let (mut element2, mut vote2) = (0, 0);

    for &v in nums.iter() {
        if vote1 > 0 && v == element1 {
            vote1 += 1;
        }else if vote2 > 0 && v == element2 {
            vote2 += 1;
        }else if vote1 == 0 {
            element1 = v;
            vote1 += 1;
        }else if vote2 == 0 {
            element2 = v;
            vote2 += 1;
        }else {
            vote1 -= 1;
            vote2 -= 1;
        }
    }

    let (mut cnt1, mut cnt2) = (0, 0);

    for &v in nums.iter() {
        if vote1 > 0 && v == element1 {
            cnt1 += 1;
        }
        if vote2 > 0 && v == element2 {
            cnt2 +=1;
        }
    }

    let mut result = vec![];

    if cnt1 > nums.len() / 3 {
        result.push(element1);
    }
    if cnt2 > nums.len() / 3 {
        result.push(element2);
    }

    result
}

pub fn majority_element_2(nums: Vec<i32>) -> Vec<i32> {

    // rusty 的写法

    let (mut element1, mut vote1) = (0, 0);
    let (mut element2, mut vote2) = (0, 0);


    for &v in nums.iter() {
        match (vote1 > 0, v == element1, vote2 > 0, v == element2) {
            (true, true, _, _) => vote1 +=1,
            (_, _, true, true) => vote2 +=1,
            (false, _, _, _) => {element1 = v; vote1 += 1},
            (_,_,false, _)  => {element2 = v; vote2 += 1},
            _ => {vote1 -= 1; vote2 -= 1}
        }
    }

    // let mut ans = Vec::new();
    let mut ans = vec![];

    if vote1 > 0 && nums.iter().filter(|&x| *x == element1).count() > nums.len() / 3 {
        ans.push(element1);
    }

    if vote2 > 0 && nums.iter().filter(|&&x| x == element2).count() > nums.len() / 3 {
        ans.push(element2);
    }

    ans
   
}