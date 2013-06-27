package symbolhash

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	hash := New("12345678901234567890123456789012", 8)
	fmt.Println(hash)
	fmt.Println(New("Alexander", 8))
	fmt.Println(New("var", 8))
	fmt.Println(New("her", 8))
	fmt.Println(New("abc123HalloHalloDerJajajajaja", 8))

	fmt.Println(New("Alexander", 16))
	fmt.Println(New("var", 16))
	fmt.Println(New("her", 16))
	fmt.Println(New("abc123HalloHalloDerJajajajaja", 16))
}
