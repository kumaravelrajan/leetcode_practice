from typing import List

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        lookup_table = {}
        keys_to_delete = {}
        # ct_consec_elems = 0
        max_consec_elems = 0

        for num in nums:
            lookup_table[num] = True

        for key in lookup_table:

            ct_consec_elems = 0
            ct_key = key

            if (ct_key + 1) not in lookup_table and (ct_key - 1) not in lookup_table:
                keys_to_delete[ct_key] = True
                continue
            
            else:

                # This means there is either ct_key + 1 OR ct_key - 1 present in lookup_table. In any case ct_key forms a part in a consecutive element chain. Hence, we make sure to include this in the ct_consec_elems.

                ct_consec_elems += 1
                keys_to_delete[ct_key] = True

                # Check if ct_key + 1 present in lookup_table
                while (ct_key + 1) in lookup_table and (ct_key + 1) not in keys_to_delete:
                    ct_key += 1
                    ct_consec_elems += 1
                    keys_to_delete[ct_key] = True

                # Reset ct_key to key
                ct_key = key

                # Check if ct_key - 1 present in lookup table
                while (ct_key - 1) in lookup_table and (ct_key - 1) not in keys_to_delete:
                    ct_key -= 1
                    ct_consec_elems += 1
                    keys_to_delete[ct_key] = True

            if ct_consec_elems > max_consec_elems:
                max_consec_elems = ct_consec_elems

        return max_consec_elems


sol = Solution()
print(sol.longestConsecutive([100, 4, 200, 1, 3, 2]))
print(sol.longestConsecutive([0,3,7,2,5,8,4,6,0,1]))
