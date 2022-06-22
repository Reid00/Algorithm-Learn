def select_sort(alist):
    # 选择排序
    n = len(alist)

    for j in range(n - 1):
        min_index = j
        for i in range(j + 1, n):
            if alist[min_index] > alist[i]:
                min_index = i

        if j != min_index:
            alist[min_index], alist[j] = alist[j], alist[min_index]


if __name__ == '__main__':
    li = [54, 26, 93, 71, 18, 44, 55]
    print(li)
    select_sort(li)
    print(li)
