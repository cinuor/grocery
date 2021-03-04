# -*- coding: utf-8 -*-
#!/usr/bin/env python

def countSort(filename: str, newfilename: str):

    # prepare file
    min, max = 0, 120
    opened_files = list()
    of = open(filename, "r")
    nf = open(newfilename, "w", 1024)

    for i in range(min, max + 1):
        f = open("{}.txt".format(i), "w+", 1024)
        opened_files.append(f)

    # spread
    of.seek(0)
    for line in of:
        person = line.strip("\n").split(",")  # ["abcd", "35"]
        index = int(person[1])
        f = opened_files[index]
        f.write(line)

    for f in opened_files:
        f.seek(0)
        for line in f:
            nf.write(line)
        f.close()
    nf.close()


countSort("./person.txt", "sorted_person.txt")
