package hashmap

// isHappy 快乐数
func isHappy(n int) bool {
	// 对于happy num 如果是一个快乐数 最终平方和相加为1
	// if not, 会以某个数为节点 进行循环 => hashMap 检测循环点/ 快慢指针检测一个链表是否有环
	// HashMap
	hmap := make(map[int]bool)

	// hmap[n]默认false  代表n 不在循环里面，true 代表已经存在，进入环中
	for ; n != 1 && !hmap[n]; n, hmap[n] = step(n), true {
	}

	return n == 1
}

func step(n int) int {
	var sum int

	for n > 0 {
		// 个位数字的平方和
		sum += (n % 10) * (n % 10)

		n /= 10
	}
	return sum
}

func isHappy2(n int) bool {
	// 快慢指针
	// 快指针先走两步，最终如果有环 两者会在某处相遇
	slow, fast := n, step(n)

	for slow != fast && fast != 1 {
		slow, fast = step(slow), step(step(fast))
	}

	return fast == 1
}
