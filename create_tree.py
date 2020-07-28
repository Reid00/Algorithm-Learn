"""
树的基本操作，创建，广度优先，深度优先遍历
"""


class Node:
    """
    树的基本节点类
    """

    def __init__(self, elem, left=None, right=None):
        """节点的初始化,
        elem 为节点的元素
        left 为左孩子
        right 为右孩子
        """
        self.elem = elem
        self.left = left
        self.right = right


class BinTree:
    """
    二叉树的构建,用队列存储节点元素，方便后续操作
    """

    def __init__(self):
        """二叉树的初始化
        """
        self.root = None  # 初始化根节点为None

    def add(self, elem):
        """为树添加元素elem
        """
        node = Node(elem)

        if self.root is None:
            self.root = node
        else:
            queue = []
            queue.append(self.root)
            while queue:  # 对已有的节点进行层次遍历
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

    def breadth_traversal(self):
        """
        二叉树的广度优先遍历
        """
        queue = [self.root]
        while queue:
            cur_node = queue.pop(0)
            print(cur_node.elem, end=' ')
            if cur_node.left is not None:
                queue.append(cur_node.left)
            if cur_node.right is not None:
                queue.append(cur_node.right)

    # 二叉树的遍历
    # 前序遍历递归的方法
    def preorder_traversal(self, node):
        # 前序遍历 根，左，右
        if not node:
            return
        print(node.elem)
        self.preorder_traversal(node.left)
        self.preorder_traversal(node.right)

    #中序遍历递归的方法
    def inorder_traversal(self, node):
        # 中序遍历，左，根，右
        if node is None:
            return 'this is empty'
        self.inorder_traversal(node.left)
        print(node.elem)
        self.inorder_traversal(node.right)

    # 后序遍历递归的方法
    def postorder_traversal(self, root):
        # 后序遍历 左，右，根
        if not root:
            return
        self.preorder_traversal(root.left)
        self.preorder_traversal(root.right)
        print(root.elem)

    # 中序遍历的堆栈实现的代码如下
    def inorder_traversal_stack(self):
        # 中序遍历，左，根，右
        if not self.root:
            return
        stack = []  # 获取当前节点
        res = []    # 存储遍历出来的节点元素值
        cur_node = self.root
        while cur_node or stack:
            while cur_node:  # 当前节点不为None，将其添加到stack 中
                stack.append(cur_node)
                cur_node = cur_node.left  # 寻找当前节点的左子节点，直到当前节点无左子节点跳出内循环
            cur_node = stack.pop()  # 当前节点pop 出stack获取当前节点的地址值
            res.append(cur_node.elem)  # 将当前节点的数据添加到result 中
            cur_node = cur_node.right  # 寻找当前节点的右子节点
        print(res)

    def preorder_traversal_stack(self):
        """
        深度优先遍历之前序遍历用非递归的方式实现 -- 栈
        前序遍历: 根，左，右
        """
        if self.root is None:
            return 'this is empty'
        stack = []  # 获取节点
        res = []  # 最终遍历的结果
        cur_node = self.root
        while stack or cur_node:
            while cur_node:
                res.append(cur_node.elem)
                stack.append(cur_node)
                cur_node = cur_node.left
            cur_node = stack.pop()
            cur_node = cur_node.right
        print(res)

    def postorder_traversal_stack(self):
        """
        深度优先遍历之后序, 非递归的方式实现 -- 栈
        后序遍历: 左，右，根
        """
        if self.root is None:
            print('this is empty')
        stack1 = []
        stack2 = []
        res = []
        cur_node = self.root
        stack1.append(cur_node)
        while stack1:    # 找出后序遍历的逆序，存放在 stack2中
            cur_node = stack1.pop()
            if cur_node.left:
                stack1.append(cur_node.left)
            if cur_node.right:
                stack1.append(cur_node.right)
            stack2.append(cur_node)
        while stack2:   # 将 stack2中的元素出栈，即是后序遍历序列
            print(stack2.pop().elem)
        


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
    tree.breadth_traversal()
    tree.inorder_traversal_stack()
    tree.postorder_traversal(tree.root)
    print('====')
    tree.postorder_traversal_stack()
