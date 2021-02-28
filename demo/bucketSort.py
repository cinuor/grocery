# -*- coding: utf-8 -*-
#!/usr/bin/env python

def find_range(of, filename):
    min = 10000000000
    max = -1
    #  with open(filename, "r", 1024) as f:
    for line in of:
        num = int(line.strip('\n'))
        if num < min:
            min = num
        if num > max:
            max = num
    return min, max


def buketSort(cap, filename, new_filename):
    opened_files = list()
    of = open(filename, "r", 1024)
    nf = open(new_filename, "w", 1024)
    min, max = find_range(of, filename)

    # 新建小文件
    n = (max - min) // cap + 1
    for i in range(0, n):
        name = "{0}-{1}.txt".format(min + i * cap, min + (i + 1) * cap)
        f = open(name, "w+", 128)
        opened_files.append(f)

    # 把大文件中的数字分发到小文件中
    of.seek(0)
    for line in of:
        num = int(line.strip("\n"))
        f = opened_files[num // cap]
        f.write(str(num) + '\n')

    for f in opened_files:
        f.flush()
        #  f.close()

    # 把小文件排序，并把排序后的数字写入到新文件中
    for f in opened_files:
        f.seek(0)
        nums = [int(x) for x in f.readlines()]
        nums.sort()
        for num in nums:
            nf.write(str(num) + '\n')
        f.close()

    nf.flush()
    # 不再需要旧文件，关闭旧文件
    of.close()
    # 关闭新文件
    nf.close()


buketSort(1000000, "./numbers.txt", "./sorted_numbers.txt")
