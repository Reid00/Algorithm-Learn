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

async def main():
    print('hello')
    await asyncio.sleep(1)
    print('world')


def gen():
    x = yield 1
    print('x is:', x)
    y = yield 2
    print('y is:', y)



if __name__ == "__main__":
    # asyncio.run(main())
    g = gen()
    # res = g.send(None)
    res = next(g)
    print('第一次yield 的返回值:', res)
    res = g.send('测试')
    print('第二次yield 的返回值:', res)
    g.send('y is replaced')
