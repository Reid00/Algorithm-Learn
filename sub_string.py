# -*- encoding: utf-8 -*-
'''
@File        :sub_string.py
@Time        :2020/12/21 11:54:49
@Author      :Reid
@Version     :1.0
@Desc        :最小覆盖子串 -- 滑动窗口
'''

import collections
import pandas as pd
import numpy as np
from matplotlib import pyplot as plt


class Solution:
    def right(self, s):
        """
        判断多括号的合法性
        """
        # 维护一个栈，用来存储左括号，如果碰到右括号，从栈中拿出最近的左括号比对，是否匹配
        stack = []
        mapping = {
            ')': '(',
            ']': '[',
            '}': '{'
        }

        for char in s:
            if char in ['{', '(', '[']:
                stack.append(char)
            else:
                # 判断右括号和最近的左括号是否匹配
                if stack and mapping[char] == stack[-1]:
                    stack.pop()
                else:
                    # 和最近的左括号不匹配
                    return False
        return not stack


def data_visualization():
    """
    数据可视化
    """
    print('data visualization...')
    df = pd.DataFrame(np.abs(np.random.randn(5, 3)), 
                index=['Mon', 'Tue', 'Wen', 'Thir', 'Fri'], 
                columns=['A', 'B', 'C'])
    df.plot()
    plt.show()

if __name__ == "__main__":
    so = Solution()
    s = "[]())"
    # t = "ABC"
    print(so.right(s))
    data_visualization()