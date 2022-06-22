def binary_search(alist, item):
    # 二分查找，递归法
    n = len(alist)
    if n > 0:
        mid = n // 2
        if alist[mid] == item:
            return True, mid
        elif item < alist[mid]:
            return binary_search(alist[:mid], item)
        else:
            return binary_search(alist[mid + 1:], item)
    return False


def binary_search_2(alist, item):
    # 二分查找，非 递归方法
    n = len(alist)
    first = 0
    last = n - 1
    while first <= last:
        mid = (first + last) // 2
        if alist[mid] == item:
            return True, mid
        elif item < alist[mid]:
            last = mid - 1
        else:
            first = mid + 1
    return False


if __name__ == '__main__':
    li = [18, 26, 44, 54, 55, 71, 93]
    print(binary_search_2(li, 54))
    print(binary_search_2(li, 100))
