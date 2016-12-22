package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strconv"
)

///https://github.com/durango/go-credit-card/blob/master/creditcard_test.go

//https://gist.github.com/FreedomCoder/2981812#file-luhn-go
//https://github.com/durango/go-credit-card/blob/master/creditcard.go
func Luhn(card string) bool {
	var sum int
	var alternate bool

	numberLen := len(card)

	if numberLen < 13 || numberLen > 19 {
		return false
	}

	for i := numberLen - 1; i > -1; i-- {
		mod, _ := strconv.Atoi(string(card))
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate

		sum += mod
	}

	return sum%10 == 0
}

func hash(s string) string {

	h := sha1.New()
	h.Write([]byte(s))
	sum := h.Sum(nil)
	bs := base64.URLEncoding.EncodeToString(sum)
	//res := fmt.Sprintf("%s", bs)
	//	fmt.Println(res)
	return (bs)
}

func print(s string) {
	fmt.Println(s)
}
func iterRange(rangeStart int, rangeEnd int) {
	for i := rangeStart; i <= rangeEnd; i++ {
		c := strconv.Itoa(i)
		res := Luhn(c)
		if res {
			fmt.Println(i)

		}
	}
}

func main() {
	iterRange(5105105105105100, 5105105105105200)
}
