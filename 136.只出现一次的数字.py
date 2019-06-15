#
# @lc app=leetcode.cn id=136 lang=python
#
# [136] 只出现一次的数字
#
class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return 2*sum(set(nums)) - sum(nums)

