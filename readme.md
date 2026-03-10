# Strategies for leetcode

## Linked Lists
1. Finding middle of Linked list

    We use the "tortoise and hare" (slow and fast pointers) technique. Both pointers start at the head. The slow pointer moves one step at a time, while the fast pointer moves two steps. By the time the fast pointer reaches the end of the list, the slow pointer will be at the middle node.

## Strings
1. Reverse a string

    Swap front and back, swap front and back

    ```go
    func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
        return string(runes)
    }
    ```