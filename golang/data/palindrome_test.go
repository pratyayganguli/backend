package data

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	// aaa -> mid = 3/2 = 1
	// aabb -> mid 4/2 = 2
	if isPalindrome("hlh") {

	} else {
		//
	}
	trashCode()
}

func TestTrashCode(t *testing.T) {
	trashCode()
}

func trashCode() {
	// for odd case
	// message := "olllo"
	// for even case
	// need to handle is different for odd and even lengths...
	message := "ollo"
	mid := len(message) / 2
	if len(message)%2 == 1 {
		prefix := message[0:mid]
		suffix := message[mid+1:]
		fmt.Println(prefix, " ", suffix)
	} else {
		prefix := message[0:mid]
		suffix := message[mid:]
		fmt.Println(prefix, " ", suffix)
	}
}

func isPalindrome(input string) bool {
	charArr := []byte(input)
	if len(charArr) == 0 {
		fmt.Println("This is an empty string")
	}

	len := len(charArr)
	if len/2 < 1 {
		// only one char is there in the string
		return true
	} else {
		if len/2%2 == 0 {
			return false
		}
	}
	return false

}
