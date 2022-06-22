import numpy as np

def merge(left,right):
    lst=[]
    l=r=0
    while l<len(left) and r<len(right):
        if left[l]<right[r]:
            lst.append(left[l])
            l +=1
        else:
            lst.append(right[r])
            r +=1

    if l==len(left):
        for i in right[r:]:
            lst.append(i)
    else:
        for i in left[l:]:
            lst.append(i)
    # method two:
    # if l==len(left):
    #     arr=np.hstack((arr,right[r:]))
    # else:
    #     arr=np.hstack((arr,left[l:]))
    return lst

def merge_sort(lst):
    if len(lst)<=1:
        return lst
    middle= len(lst)//2
    left=merge_sort(lst[:middle])
    right=merge_sort(lst[middle:])
    return merge(left,right)

if __name__ == "__main__":
    lst=[1,9,2,100,50,6,3]
    res=merge_sort(lst)
    print(res)

