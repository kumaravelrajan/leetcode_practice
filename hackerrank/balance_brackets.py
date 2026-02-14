def isBalanced(s):
    # Write your code here
    bracket_sequence_stack = []
    opening_brackets_hash_table = {"{": "}", "[": "]", "(": ")"}
    closing_brackets_hash_table = {"}": "{", "]":"[", ")":"("}
    
    for char in s:
        if char in opening_brackets_hash_table:
            bracket_sequence_stack.append(char)
        elif char in closing_brackets_hash_table:
            if not bracket_sequence_stack:
                return "NO"
            elif bracket_sequence_stack[-1] == closing_brackets_hash_table.get(char):
                bracket_sequence_stack.pop()
            else:
                return "NO"
    
    if len(bracket_sequence_stack) == 0:
        return "YES"
    else:
        return "NO"

print(isBalanced("[({})]"))