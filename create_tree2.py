# -*- encoding: utf-8 -*-
'''
@File        :create_tree2.py
@Time        :2020/12/07 16:22:44
@Author      :Reid
@Version     :1.0
@Desc        :二叉树的基本操作(create, traversal, 二叉树所有节点值加一, 判断两个二叉树是否完全相同等等)
'''

from collections import deque
from collections.abc import Iterable


class Node:
    """定义树的节点"""

    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class BinaryTree:
    """
    二叉树
    """

    def __init__(self):
        self.root = None

    def add(self, value):
        """
        给二叉树添加元素
        """
        node = Node(value)
        if self.root is None:
            self.root = node
        else:
            queue = deque([self.root])
            while queue:
                cur_node = queue.popleft()
                if cur_node.left is None:
                    cur_node.left = node
                    return 
                else:
                    queue.append(cur_node.left)
                if cur_node.right is None:
                    cur_node.right = node
                    return
                else:
                    queue.append(cur_node.right)

    def bfs(self):
        """
        广度优先遍历
        """
        queue = deque([self.root])
        visited = []
        while queue:
            cur_node = queue.popleft()
            visited.append(cur_node.val)
            if cur_node.left is not None:
                queue.append(cur_node.left)
            if cur_node.right is not None:
                queue.append(cur_node.right)
        return visited


def value_plus_one(node: Node):
    """
    把二叉树所有的节点值加一
    """
    if node is None:
        return
    node.val += 1
    value_plus_one(node.left)
    value_plus_one(node.right)


def is_same_tree(node1: Node, node2: Node):
    """
    判断两颗二叉树是否完全相同
    """
    # 两者都为None，相同
    if node1 is None and node2 is None:
        return True
    # 一个为空，另一个非空，不同  
    if node1 is None or node2 is None:
        return False  
    # 两个都不为空，但是根节点值不相同
    if not node1.val == node2.val:
        return False
    # 比较除了根节点以外的节点
    return is_same_tree(node1.left, node2.left) and is_same_tree(node1.right, node2.right)


if __name__ == "__main__":
    t = BinaryTree()
    for i in range(10):
        t.add(i)
    print(t.bfs())
    value_plus_one(t.root)
    # print(res)
    print(t.bfs())
    t1 = BinaryTree()
    t1.add(2)
    t1.add(3)
    t1.add(4)
    t1.add(5)
    t1.add(6)
    t1.add(7)
    print(is_same_tree(t.root, t1.root))




