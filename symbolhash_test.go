package symbolhash

import (
	"fmt"
	"testing"
)

func TestBasics(t *testing.T) {
	hash := New("12345678901234567890123456789012", 8)
	fmt.Println(hash)
	fmt.Println(New("Alexander", 8))
	fmt.Println(New("var", 8))
	fmt.Println(New("her", 8))
	fmt.Println(New("abc123 Testing Testing", 8))

	fmt.Println(New("Alexander", 16))
	fmt.Println(New("var", 16))
	fmt.Println(New("her", 16))
	fmt.Println(New("abc123 Testing Testing", 16))

	fmt.Println(New("Several different words lalalalala", 16))
	fmt.Println(New("hi", 16))
}
