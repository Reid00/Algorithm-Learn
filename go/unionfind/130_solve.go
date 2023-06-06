package unionfind

// 但这个问题也可以用 Union-Find 算法解决，虽然实现复杂一些，甚至效率也略低，
// 但这是使用 Union-Find 算法的通用思想，值得一学。
// 你可以把那些不需要被替换的 O 看成一个拥有独门绝技的门派，
// 它们有一个共同「祖师爷」叫 dummy，这些 O 和 dummy 互相连通，而那些需要被替换的 O 与 dummy 不连通。

// 首先要解决的是，根据我们的实现，Union-Find 底层用的是一维数组，
// 构造函数需要传入这个数组的大小，而题目给的是一个二维棋盘。

// 这个很简单，二维坐标 (x,y) 可以转换成 x * n + y 这个数（m 是棋盘的行数，n 是棋盘的列数）
// 敲黑板，这是将二维坐标映射到一维的常用技巧。

// 其次，我们之前描述的「祖师爷」是虚构的，需要给他老人家留个位置。
// 索引 [0.. m*n-1] 都是棋盘内坐标的一维映射，那就让这个虚拟的 dummy 节点占据索引 m * n 好了。
func solve(board [][]byte) {
	// 1. 创建并查集
	m := len(board)
	n := len(board[0])
	// m行 n 列，转化成一维数组长度 下标 m * n -1
	// 虚拟出父节点 dummy 和dummy 相连，意味着是有连通的，不会被替换
	dummy := m * n //dummy 的坐标

	uf := NewUnionFind(dummy + 1)

	// 2. 把board 的数据放到并查集中
	// 先把首列和尾列的O 和dummy 相连，因为四边的O 不会被替换
	for i := 0; i < m; i++ {
		// 首列
		if board[i][0] == 'O' {
			// 二维转一维 (x,y) => (x*n + y) n是列数
			uf.Union(dummy, i*n)
		}
		// 尾列
		if board[i][n-1] == 'O' {
			uf.Union(dummy, i*n+n-1)
		}
	}

	// 把首行和尾行的O 和dummy 相连
	for i := 0; i < n; i++ {
		// 首行
		if board[0][i] == 'O' {
			uf.Union(dummy, i)
		}
		// 尾行
		if board[m-1][i] == 'O' {
			uf.Union(dummy, (m-1)*n+i)
		}
	}

	// 方向数组 diretion 是上下左右搜索的常用手法
	diretion := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	// 3. 进行Union
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			// 从坐标(1,1) 开始遍历
			if board[i][j] == 'O' {
				// 将此 O 与上下左右的 O 连通
				for k := 0; k < 4; k++ {
					// x 左右移动
					x := i + diretion[k][0]
					// y 上下移动
					y := j + diretion[k][1]
					if board[x][y] == 'O' {
						// (x,y) 和 (i,j) 相连
						uf.Union(x*n+y, i*n+j)
					}
				}
			}
		}
	}
	// 4. 再遍历一次board 把不与dummy 相连的O替换成x
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' && !uf.Connected(dummy, i*n+j) {
				board[i][j] = 'X'
			}

		}
	}

}

// UnionFind 并查集 用一维数组表示连通分量
type UnionFind struct {
	// parent[x] 代表 x 的父节点是parent[x]
	parent []int
	// 连通分量的数量
	count int
}

// NewUnionFind 创建长度为n 的并查集
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	// 初始时，所有的父节点指向自己， 连通分量的数量为n
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
		count:  n,
	}
}

// Find 找到节点x 的父节点
func (u *UnionFind) Find(x int) int {

	if u.parent[x] != x {
		// 路径压缩，把x 到根节点parent[x] 之间的所有节点，都指向parent[x]
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

// Union 把节点x 和 y 连接起来
func (u *UnionFind) Union(x, y int) {
	rootX := u.Find(x)
	rootY := u.Find(y)
	if rootX == rootY {
		return
	}
	// 把x的根节点指向y的根节点，或者反过来也可以
	u.parent[rootX] = rootY
	u.count--
}

// Connected 判断x, y 是否连通，即是否有相同的父节点
func (u *UnionFind) Connected(x, y int) bool {
	rootX := u.Find(x)
	rootY := u.Find(y)

	return rootX == rootY
}

// Count 返回连通分量的个数
func (u *UnionFind) Count() int {
	return u.count
}
