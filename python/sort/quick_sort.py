def quick_sort(alist):
    # 快速排序,递归方法实现
    left = []
    right = []
    pivot_list = []
    if len(alist) <= 1:
        return alist
    else:
        pivot = alist[0]
        for item in alist:
            if item < pivot:
                left.append(item)
            elif item > pivot:
                right.append(item)
            else:
                pivot_list.append(item)
        left = quick_sort(left)
        right = quick_sort(right)
        return left + pivot_list + right


if __name__ == "__main__":
    li = [4, 65, 2, -31, 0, 99, 83, 782, 1]
    print(li)
    print(quick_sort(li))
