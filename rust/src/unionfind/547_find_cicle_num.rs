

impl Solution {
    pub fn find_circle_num(is_connected: Vec<Vec<i32>>) -> i32 {

        let n = is_connected.len();
        let mut u_find = UnionFind::new(n);

        // 第 i 行代表 i 和每个元素的联通情况，1 代表连通
        for (i, row) in is_connected.iter().enumerate() {
            // 因为i== j 的时候是代表自己，不需要连通图
            // i 行 会和每个元素判断是否连通，j 行的时候不能重复判断，所以j > i或者j = i +1
            for (j, item) in row.iter().enumerate() {

                if j > i && *item == 1 {
                    u_find.union(i, j)
                }
            }
        }
        
        u_find.count()

    }
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


    fn find(&self, x: usize) -> usize {

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

    fn connect(&self, x: usize, y:usize) -> bool {
        let root_x = self.find(x);
        let root_y = self.find(y);
        root_x == root_y
    }

    fn count(&self) -> i32 {
        self.count as i32
    }

}