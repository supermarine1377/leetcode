from typing import List

class Solution:
    def twoSumN2(self, nums: List[int], target: int) -> List[int]:
        l = len(nums)
        for i in range(l):
            for j in range(i + 1, l):
                if nums[i] + nums[j] == target:
                    return [i, j]

    def twoSumN1(self, nums: List[int], target: int) -> List[int]:
        h = {}
        for i, n in enumerate(nums):
            want = target - n
            if want in h:
                return i, h[want]
            h[n] = i
