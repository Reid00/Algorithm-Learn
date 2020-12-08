# -*- encoding: utf-8 -*-
'''
@File        :BST.py
@Time        :2020/12/08 10:51:25
@Author      :Reid
@Version     :1.0
@Desc        :二叉搜索树的基础操作，判断BST 的合法性，增删查.(删，查略微复杂)
'''

# 二叉搜索树（Binary Search Tree，简称 BST）是一种很常用的的二叉树。
# 它的定义是：一个二叉树中，任意节点的值要大于等于左子树所有节点的值，且要小于等于右边子树的所有节点的值。
from collections import deque
from create_tree2 import BinaryTree
from create_tree2 import Node



def valid_BST(root, small, large):
    """
    :param root, small, large: 根节点，二叉树中的最小值，二叉树中的最大值
    """
    if root is None:
        return True
    if small >= root.val or large <= root.val:
        return False
    return valid_BST(root.left, small, root.val) and valid_BST(root.right, root.val, large)


def is_valid_BST(root):
    return valid_BST(root, -2**32, 2**32 - 1)


def is_in_bst(root: Node, target: int):
    """
    target 是否在二叉搜索树中
    """
    # 如果不是二叉搜索树遍历所有树
    if root is None:
        return False
    if root.val == target:
        return True
    return is_in_bst(root.left, target) or is_in_bst(root.right, target)
    # 如果是二叉搜索树遍历所有树, 所有左子树的值都小于根节点，右子树的值都大于根节点
    # if root is None:
    #     return False
    # if root.val == target:
    #     return True
    # if root.val < target:
    #     return is_in_bst(root.right, target)
    # if root.val > target:
    #     return is_in_bst(root.left, target)


def insert_into_bsf(root: Node, value: int):
    """
    在二叉搜索树中插入一个元素, 一旦涉及到“改”的操作，最终返回的都要是Node类型
    """
    if root is None:
        return Node(value)
    if root.val == value:
        # BST 一般不是插入相同的元素
        return 'exists the same element'
    if root.val < value:  # 做右子树进行插入
        root.right = insert_into_bsf(root.right, value)
    if root.val > value:
        root.left = insert_into_bsf(root.left, value)
    return root


def delete_node(root: Node, value: int):
    """
    在BST 中删除一个元素
    """
    if root is None:
        return
    if root.val == value:
        # 执行删除操作
        # 1. 如果左右孩子都为None 直接删除，该节点不存在了，返回None
        if root.left is None and root.right is None:
            return
        # 2. 如果只存右孩子， 删除后只剩下右孩子
        if root.left is None:
            return root.right
        if root.right is None:
            return root.left
        # 3. 如果存在左右孩子 1) 用左孩子的最大值替换该节点，然后删除左孩子的最大值 2) 用右孩子的最小值替换该节点，然后删除右孩子的最小值
        if root.left and root.right:
            min_node = get_min_node(root.right)
            root.val = min_node.val
            root.right = delete_node(root.right, min_node.val)

    if root.val < value:
        root.right = delete_node(root.right, value)
    if root.val > value:
        root.left = delete_node(root.left, value)
    return root


def get_min_node(root: Node):
    """
    获取BST中的最小节点。既是最左侧的节点
    """
    while root.left:
        root = root.left
    return root


def BFS(root: Node):
    """
    广度优先遍历
    """
    if root is None:
        return
    visited = []
    queue = deque([root])
    while queue:
        cur_node = queue.popleft()
        visited.append(cur_node.val)
        if cur_node.left is not None:
            queue.append(cur_node.left)
        if cur_node.right is not None:
            queue.append(cur_node.right)
    return visited


if __name__ == "__main__":
    b_tree = BinaryTree()
    for i in range(4, 20, 3):
        b_tree.add(i)
    print(b_tree.bfs())
    print(b_tree.root.val)
    print('is valid BST: ', is_valid_BST(b_tree.root))
    print('is In BST: ', is_in_bst(b_tree.root, 20))
    print('insert into BST: ', BFS(insert_into_bsf(b_tree.root, 2)))
    print('get min node: ',get_min_node(b_tree.root).val)
    print('delete a node: ', BFS(delete_node(b_tree.root, 13)))