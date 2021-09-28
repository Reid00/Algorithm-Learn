# -*- encoding: utf-8 -*-
'''
@File        :iter_test.py
@Time        :2020/12/16 09:07:15
@Author      :Reid
@Version     :1.0
@Desc        :可迭代对象的测试
'''


import asyncio
import socket
from collections import Counter


async def main():
    print('hello')
    await asyncio.sleep(1)
    print('world')


def gen():
    x = yield 1
    print('x is:', x)
    y = yield 2
    print('y is:', y)


def single_number(nums: list):
    """找出只出现一次的数字
    """
    res = 0
    for i in nums:
        res ^= i
    return res


def reverse_string(s, left, right):
    """
    判断是否为回文s[left: right+1]
    """
    print(s[left: right+1])
    while left < right:
        if s[left] != s[right]:
            return False
        left += 1
        right -= 1

    return True


def quick_sort(lst: list, start, end):
    """快速排序

    Args:
        lst (List): [description]
    """
    if start >= end:
        return
    pivot = lst[start]
    low, high = start, end

    while low < high:
        while low < high and lst[high] >= pivot :
            high -= 1
        lst[low] = lst[high]

        while low < high and lst[low] < pivot:
            low += 1
        lst[high] = lst[low]

    lst[low] = pivot
    quick_sort(lst, start, low -1 )
    quick_sort(lst, low + 1, end)



if __name__ == "__main__":
    # a = 'abcdaa'
    # b = 'bcaada'
    # cnt_a, cnt_b = Counter(a), Counter(b)
    # print(single_number([4,1,2,1,2]))
    # print(8&(8-1))

    # print(reverse_string('easgsaeddhh', 0, 6))
    lst = [10,8,11,55,22,3]
    quick_sort(lst, 0, 5)
    print(lst)