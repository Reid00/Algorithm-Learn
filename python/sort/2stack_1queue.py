# 两个栈实现堆

class MyQueue:
    def __init__(self, val):
        self.stackin = []
        self.stackout = []
        self.val = val

    def push(self):
        self.stackin.append(self.val)

    def pop(self):
        if self.stackout:
            return self.stackout.pop()
        while self.stackin:
            self.stackout.append(self.stackin.pop())

        return self.stackout.pop() if self.stackout else u'队列为空'
