

class Node:
    # 定义树的节点
    def __init__(self,data=None,left=None,right=None):
        # 初始化类，data 为根节点，left 为做子树，right 为右子树
        self.data = data
        self.left = left
        self.right =right

class BinTree:
    def __init__(self):
        #初始化类
        self.root = None         #初始化根节点为None
        self.lst= []              #定义列表用于存储节点的地址

    def insert(self,data):
        # 定义insert 方法，向树结构中添加元素
        node = Node(data)           # 实例化树的节点
        if self.root is None:
            self.root=node
            self.lst.append(self.root)
        else:
            root_node = self.lst[0]     #将第一个元素设置为根节点
            if root_node.left is None:
                root_node.left = node
                self.lst.append(root_node.left)
            elif root_node.right is None:
                root_node.right = node
                self.lst.append(root_node.right)
                self.lst.pop(0)     # 弹出self.lst 第一个位置的元素
    
    # 二叉树的遍历
    # 前序遍历递归的方法
    def preorder_traversal(self,root):
        # 前序遍历 根，左，右
        if not root:
            return
        print(root.data)
        self.preorder_traversal(root.left)
        self.preorder_traversal(root.right)

    #中序遍历递归的方法
    def inorder_traversal(self,root):
        # 中序遍历，左，根，右
        if root is None:
            return 'this is empty'
        self.inorder_traversal(root.left)
        print(root.data)
        self.inorder_traversal(root.right)

    # 后序遍历递归的方法
    def postorder_traversal(self,root):
        # 后序遍历 左，右，根
        if not root:
            return
        self.preorder_traversal(root.left)
        self.preorder_traversal(root.right)
        print(root.data)

    # 中序遍历的堆栈实现的代码如下
    def inorder_traversal_stack(self,root):
        # 中序遍历，左，根，右
        if not root:
            return
        stack =[]
        res =[]
        node = root
        while node or stack:
            while node:                 #当前节点不为None，将其添加到stack 中
                stack.append(node)
                node = node.left        #寻找当前节点的左子节点，直到当前节点无左子节点跳出内循环
            node  = stack.pop()         #当前节点pop 出stack获取当前节点的地址值
            res.append(node.data)       #将当前节点的数据添加到result 中
            node = node.right           # 寻找当前节点的右子节点
        print(res)
            

def main():
    tree= BinTree()
    for i in range(10):
        tree.insert(i)
    tree.preorder_traversal(tree.root)

if __name__ == "__main__":
    main()