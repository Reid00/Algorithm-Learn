class Deque:
    # 双端队列
    def __init__(self):
        self.__list = []

    def add_front(self, item):
        # 从前面入队
        self.__list.insert(0, item)

    def add_rear(self, item):
        # 从后面入队
        self.__list.append(item)

    def pop_front(self):
        # 从前面出队
        return self.__list.pop(0)

    def pop_rear(self):
        # 从后面出队
        return self.__list.pop()

    def is_empty(self):
        # 判空
        return self.__list == []

    def size(self):
        # 返回大小
        return len(self.__list)


if __name__ == '__main__':
    q = Deque()
    q.add_front(0)
    q.add_front(1)
    q.add_rear(2)
    q.add_rear(3)
    print(q.pop_front())
    print(q.pop_front())
    print(q.pop_rear())
    print(q.pop_rear())