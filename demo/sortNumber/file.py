# -*- coding: utf-8 -*-
#!/usr/bin/env python3

import random


def random_number(min, max):
    return random.randint(min, max)


def generate_file(filename, min, max):
    with open(filename, "w+", 1024) as f:
        for i in range(min, max):
            n = random_number(min, max)
            f.write(str(n) + '\n')


generate_file("./numbers.txt", 0, 100000000)
