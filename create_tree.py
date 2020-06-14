"""
树的基本操作，创建，广度优先，深度优先遍历
"""


class Node:
    """
    树的基本节点类
    """
    def __init__(self,elem,left=None,right = None):
        """节点的初始化,
        elem 为节点的元素
        left 为左孩子
        right 为右孩子
        """
        self.elem = elem
        self.left = left
        self.right =right

class BinTree:
    """
    二叉树的构建,用队列存储节点元素，方便后续操作
    """
    def __init__(self):
        """二叉树的初始化
        """
        self.root = None         #初始化根节点为None

    def add(self,elem):
        """为树添加元素elem
        """
        node = Node(elem)

        if self.root is None:
            self.root = node
        else:
            queue = []
            queue.append(self.root)
            while queue: #对已有的节点进行层次遍历
                cur_node = queue.pop(0)
                if cur_node.left is None:
                    cur_node.left = node
                    return
                else:  # 如果左子树都不为空，加入队列继续判断
                    queue.append(cur_node.left)
                if cur_node.right is None:
                    cur_node.right = node
                    return
                else:
                    # 如果右子树都不为空，加入队列继续判断
                    queue.append(cur_node.right)


    def breadth_travel(self):
        """
        二叉树的广度优先遍历
        """
        queue = [self.root]
        while queue:
            cur_node = queue.pop(0)
            print(cur_node.elem,end= ' ')
            if cur_node.left is not None:
                queue.append(cur_node.left)
            if cur_node.right is not None:
                queue.append(cur_node.right)
            
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
            
if __name__ == "__main__":
    tree = BinTree()
    tree.add(0)
    tree.add(1)
    tree.add(2)
    tree.add(3)
    tree.add(4)
    tree.add(5)
    tree.add(6)
    tree.add(7)
    tree.breadth_travel()