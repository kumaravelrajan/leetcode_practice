class Solution:
    def maxArea(self, height: list[int]) -> int:
        heights_and_indices = {}

        for index, h in enumerate(height):

            if h not in heights_and_indices:
                heights_and_indices[h] = []

            heights_and_indices[h].append(index)

        sorted_heights_and_indices = dict(sorted(heights_and_indices.items(), reverse=True))

        print(sorted_heights_and_indices)

        max_area = 0

        for key, value in sorted_heights_and_indices.items():
            
            

sol = Solution()
sol.maxArea([1,8,6,2,5,4,8,3,7])
        