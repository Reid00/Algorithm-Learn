def bubble_sort(alist):
    # 冒泡排序
    for j in range(len(alist)):  # 控制走多少次
        for i in range(len(alist) - 1 - j):  # 控制走多少步
            if alist[i] > alist[i + 1]:
                alist[i], alist[i + 1] = alist[i + 1], alist[i]
    # return alist

if __name__ == '__main__':
    li = [54, 26, 93, 71, 18, 44, 55]
    print(li)
    bubble_sort(li)
    print(li)
