from collections import defaultdict
from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        lookup_table = defaultdict(list)

        for index, num in enumerate(nums):             
            if (target - num) in lookup_table:
                return [index, lookup_table.get(target-num)[0]]
            
            lookup_table[num].append(index)

            


sol = Solution()
print(sol.twoSum([2,7,11,15], 9))
print(sol.twoSum([3,2,4], 6))
print(sol.twoSum([3,3], 6))
