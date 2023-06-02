
// num_islands 岛屿数量
pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
    // 用并查集 将1 与四周的岛屿相连接， 0 与dummy 相连接
    let m = grid.len();
    let n= grid[0].len();

    let dummy = m * n ;
    let mut uf = UnionFind::new(dummy + 1);

    // d 常用来表达向上下左右四个方向搜索
    let d: Vec<Vec<i32>> = vec![vec![0,1], vec![0,-1], vec![1,0], vec![-1,0]];
    // 遍历二维数组，转成一维的并查集 坐标变换 (x,y) -> (x*n+y) n是列数
    for i in 0..m {
        for j in 0..n {

            if grid[i][j] == '1' {
                // 与四周的1 进行连通
                for k in 0..4 {

                    // 避免i,j 处于四周的情况进行 +1 -1 会越界
                    // BUG match pattern 不能有变量
                    // match (i,j) {
                    // (0, _) if d[k][0] == -1 => continue,
                    // (_, 0) if d[k][1] == -1 => continue,
                    // (m-1, _) if d[k][0] == 1 => continue,
                    // (_, n-1) if d[k][1] == 1 => continue,
                    // _ =>   (),                     
                    // }
                    if i == 0 && d[k][0] == -1 {
                        continue;
                    }else if j==0 && d[k][1] == -1{
                        continue;
                    }else if i==m-1 && d[k][0] == 1 {
                        continue;
                    }else if j == n-1 && d[k][1] == 1 {
                        continue;
                    }

                    let x = (i as i32 + d[k][0]) as usize;
                    let y = (j as i32 + d[k][1]) as usize;
                    if grid[x][y] == '1' {
                        uf.union(i*n+j, x*n+y);
                    }
                }
            }else {
                uf.union(dummy, i*n+j);
            }

        }
    }

    uf.count() -1
}


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