class Solution:

    """ def add_column_to_2d_list(self, zigzag_pattern):
        for row in zigzag_pattern:
            row.append(None)

    def print_2d_list(self, zigzag_pattern):
        print("############################################################")
        for row in zigzag_pattern:
            print(row)

        print("############################################################")

    def convert(self, s: str, numRows: int) -> str:

        # base case
        if numRows == 1:
            return s

        # Creating column 0 with numRows rows.
        zigzag_pattern = [[None] for _ in range(numRows)]

        row = 0
        col = 0

        numrows_minus_one = numRows - 1

        string_index = 0

        # Filling col 0
        while row < numRows:
            if string_index < len(s):
                zigzag_pattern[row][col] = s[string_index]
                row += 1
                string_index += 1
            else:
                break

        row = 0
        col += 1

        self.add_column_to_2d_list(zigzag_pattern)

        while string_index < len(s):
            self.print_2d_list(zigzag_pattern)
            if (col % numrows_minus_one == 0 and string_index < len(s)):
                zigzag_pattern[row][col] = s[string_index]
                row += 1
                string_index += 1
                if(row >= numRows):
                    col += 1
                    row = 0
                    self.add_column_to_2d_list(zigzag_pattern)

            if (col % numrows_minus_one != 0 and string_index < len(s)):
                zigzag_pattern[numrows_minus_one - (col % numrows_minus_one)][col] = s[string_index]
                string_index += 1
                col += 1
                row = 0
                self.add_column_to_2d_list(zigzag_pattern)

        final_string = ""

        for row in zigzag_pattern:
            for val in row:
                if val is not None:
                    final_string += val

        return final_string """

    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1 or numRows >= len(s):
            return s

        rows = [""] * numRows
        idx, step = 0, 1

        for ch in s:
            rows[idx] += ch
            if idx == 0:
                step = 1
            elif idx == numRows - 1:
                step = -1
            idx += step

        return "".join(rows)

sol = Solution()
sol.convert("PAYPALISHIRING", 4)
                
                        