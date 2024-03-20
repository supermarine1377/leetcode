from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> List[int]:
        if len(nums) == 1:
            return 1
        
        l, r = 0, 1
        while r < len(nums):
            if nums[l] == nums[r]:
                r = r + 1
            else:
                l = l + 1
                nums[r], nums[l] = nums[l], nums[r]
                r = r + 1

        return nums[:l+1]
            
