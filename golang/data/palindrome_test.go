package data

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	// aaa -> mid = 3/2 = 1
	// aabb -> mid 4/2 = 2
	if isPalindrome("xy") {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}

func comparison(prefix, suffix string) bool {
	prefixByteArr := []byte(prefix)
	suffixByteArr := []byte(suffix)
	j := 0
	for i := len(prefixByteArr) - 1; i >= 0; i-- {
		if prefixByteArr[i] != suffixByteArr[j] {
			return false
		}
		j++
	}
	return true
}

func isPalindrome(input string) bool {
	mid := len(input) / 2
	if len(input)%2 == 1 {
		prefix := input[0:mid]
		suffix := input[mid+1:]
		return comparison(prefix, suffix)
	} else {
		prefix := input[0:mid]
		suffix := input[mid:]
		return comparison(prefix, suffix)
	}
}
