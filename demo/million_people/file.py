# -*- coding: utf-8 -*-
#!/usr/bin/env python

import random

CHARSET = "abcdefghijklmnopqrstuvwxyz"


def make_file(filename: str, limit: int):
    with open(filename, "w+", 1024) as fd:
        for i in range(0, limit):
            offset = random.randint(0, 20)
            person = "{0},{1}\n".format(
                CHARSET[offset:offset + 5], random.randint(0, 120))
            fd.write(person)


make_file("./person.txt", 1000000)
