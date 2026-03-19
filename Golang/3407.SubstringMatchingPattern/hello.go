package main

import (
	"fmt"
	"strings"
)

// Solved using Claude
func hasMatch(s string, p string) bool {
    starIdx := strings.Index(p, "*")
    prefix := p[:starIdx]
    suffix := p[starIdx+1:]

    // Find where prefix starts in s
    start := strings.Index(s, prefix)
    if start == -1 {
        return false
    }

    // Search for suffix in the remaining string after prefix match
    rest := s[start+len(prefix):]
    if strings.Contains(rest, suffix) {
        return true
    }
    return false
}

func main() {
	fmt.Println(hasMatch("leetcode", "ee*e"))
	fmt.Println(hasMatch("car", "c*v"))
	fmt.Println(hasMatch("luck", "u*"))
	fmt.Println(hasMatch("carrarararararar", "c*"))
}
