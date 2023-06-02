
pub fn solve(board: &mut Vec<Vec<char>>) {
    let m = board.len();
    let n = board[0].len();

    let dummy = m * n;
    
    let mut uf = UnionFind::new(dummy + 1);

    // 把board 四周的O 和dummy 相连，因为这种不会被替换
    // 首列和尾列
    for i in 0..m {
        if board[i][0] == 'O' {
            // 二维坐标(x,y) 转一维 (x*n, + y) n 是列数
            uf.union(dummy, i*n);
        }

        if board[i][n-1] == 'O' {
            uf.union(dummy, i*n + n-1);
        }
    }

    // 首行和尾行
    for i in 0..n {
        if board[0][i] == 'O' {
            uf.union(dummy,  i);
        }

        if board[m-1][i] == 'O' {
            uf.union(dummy, (m-1)*n + i);
        }
    }
    // d 是用来上下左右搜索的常用方法
    let d: Vec<Vec<i32>> = vec![vec![0,1], vec![0, -1], vec![-1,0], vec![1, 0]];
    // 遍历中间部分 遇到O 把其和四周的O 连接到一起
    for i in 1..m-1 {
        for j in 1..n-1  {
            if board[i][j] == 'O' {
                for k in 0..4 {
                    let x = (i as i32 + d[k][0]) as usize;
                    let y = (j as i32 + d[k][1]) as usize;
                    // 主要把 不是边界的O 但是和边界O 相连的过滤掉
                    if board[x][y] == 'O' {
                        uf.union(i*n+j, x*n+y);
                    }
                }
            }
        }
    }

    // 最后遍历一遍board 把不与dummy 相连的O 替换成X
    for i in 1..m {
        for j in 1..n {
            if board[i][j] == 'O' && !uf.connect(i*n+j, dummy) {
                board[i][j] = 'X';
            }
        }
    }

}

// --------------------------------------------------------------------

struct UnionFind {

    // parent[i] 表示节点i 的父节点是parent[i]
    parent: Vec<usize>,
    // 连通分量
    count: usize,
}

impl UnionFind {
    
    /* 构造函数，n 为图的节点总数 */
    fn new(n:usize) -> Self {
        // 初始时，每个节点的parent 是自己
        let mut parent = vec![0;n];
        for i in 0..n {
            parent[i] = i;
        }

        UnionFind {
            parent,
            count: n,
        }
    }


    fn find(&mut self, x: usize) -> usize {

        if self.parent[x] != x {
            // 找到x 的父节点，把x 到parent[x] 之间的所有节点压缩到parent[x] 之间
            self.parent[x] = self.find(self.parent[x]);
        }

        self.parent[x]
    }

    fn union(&mut self, from: usize, to: usize) {
        let from_root  = self.find(from);
        let to_root = self.find(to);

        if from_root == to_root {
            return;
        }
        // 找到from 和  to 的父节点，把 from 放到to 下面，或者反过来
        self.parent[from_root] = to_root;
        self.count -= 1;
    }

    fn connect(&mut self, x: usize, y:usize) -> bool {
        let root_x = self.find(x);
        let root_y = self.find(y);
        root_x == root_y
    }

    fn count(&self) -> i32 {
        self.count as i32
    }

}