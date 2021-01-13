#!/usr/bin/env python
# -*- coding: utf-8 -*-

def lengthOfLIS(nums):
    length = len(nums)
    if length < 2:
        return length

    dp = [1] * len(nums)
    res = dp[0]

    for i in range(length):
        for j in range(0, i):
            if nums[j] < nums[i]:
                dp[i] = dp[j] + 1
        res = max(res, dp[i])
    return res


if __name__ == '__main__':
    #  nums = [10, 9, 2, 5, 3, 7, 101, 18]
    nums = [0, 1, 0, 3, 2, 3]
    print(lengthOfLIS(nums))
