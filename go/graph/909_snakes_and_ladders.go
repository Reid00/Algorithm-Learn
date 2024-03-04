package graph

// snakesAndLadders 蛇梯棋
// 这题相当于有个棋盘，从左下角开始蛇形向上每个格子有个从1开始递增的序号，
// 你从序号1的格子开始，你每步可以从当前格子开始往后走1~6格，你只能按照序号的升序走，
// 不能往回走（如第5格走到第4格），走到最后一格就结束了。
// 每个格子内有传送门或啥都没有，
// -1代表啥都没有，大于0的数字代表传送门，传送门意味着你可以传送到当前数字的格子去
// （如当前是35，你就可以传送到序号为35的格子去），如果你传送过去后的格子内依然是大于0的数字，
// 你不能连续传送
// 现在要求最少需要几步可以走到终点。
func snakesAndLadders(board [][]int) int {

	c := convert(board)

	visited := make(map[int]bool)

	// BFS 从1 开始遍历
	q := []int{1}

	step, target := 0, len(c)-1

	for len(q) > 0 {

		for _, v := range q {
			q = q[1:]

			if visited[v] {
				continue
			}

			if v == target {
				return step
			}

			visited[v] = true

			for j := 1; j < 6; j++ { //每次可以前进1-6步
				if next := v + j; next <= target { //防止越界

					if val := c[next]; val == -1 { // 没有跳转，把当前位置next 入队列
						q = append(q, next)
					} else { // 有跳转，将跳转后的位置val 入队列
						q = append(q, val)
					}
				}
			}

		}

		step++
	}

	return -1
}

// convert 按照蛇形的顺序遍历棋盘，将其转换为可以顺序遍历的一维数组
func convert(grid [][]int) []int {

	n := len(grid)
	// grid 中格子从1开始
	res := make([]int, n*n+1)
	// res 下标和是否z型
	k, reverse := 1, false

	for i := n - 1; i >= 0; i-- {
		if !reverse {
			for j := 0; j < n; j++ {
				res[k] = grid[i][j]
				k++
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				res[k] = grid[i][j]
				k++
			}
		}
		reverse = !reverse
	}
	return res
}
