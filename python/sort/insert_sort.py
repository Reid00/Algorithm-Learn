def insert_sort(alist):
    # 插入排序
    n = len(alist)
    for j in range(n - 1):
        for i in range(j, 0, -1):
            if alist[i] < alist[i - 1]:
                alist[i], alist[i - 1] = alist[i - 1], alist[i]
            else:
                break

if __name__ == '__main__':
    li = [54, 26, 93, 71, 18, 44, 55]
    print(li)
    insert_sort(li)
    print(li)

