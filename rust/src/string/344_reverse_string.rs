
// API
pub fn reverse_string(s: &mut Vec<char>) {
    s.reverse();
}



pub fn reverse_string2(s: &mut Vec<char>) {
       
   let (mut i, mut j) = (0, s.len());

   while i < j {
        s.swap(i, j);
        i += 1;

        j -=1 ;
   }
}


pub fn reverse_string3(s: &mut Vec<char>) {
   
    let (mut i, mut j) = (0, s.len()-1);
 
    while i < j {
         //vec 中元素类型为char  实现了 Copy trait 可以使用下面的方式进行赋值
         // 否则会出现所有权移动
        (s[i], s[j]) = (s[j], s[i]);
        i += 1;
         j -= 1;
    }
 }

 pub fn reverse_string4(s: &mut Vec<char>) {

    let mut c;

    let n = s.len();

    for i in 0..n/2 {
        c = s[i];
        s[i] = s[n-1-i];
        s[n-1-i] = c;
    }
  
}