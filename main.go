package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

const file = "data/ranges.csv"
const panLength = 16

/*
Sources:
https://github.com/durango/go-credit-card/blob/master/creditcard_test.go
https://gist.github.com/FreedomCoder/2981812#file-luhn-go
https://github.com/durango/go-credit-card/blob/master/creditcard.go
*/

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
	print("here")
	for i := rangeStart; i <= rangeEnd; i++ {
		c := strconv.Itoa(i)
		res := Luhn(c)
		if res {
			fmt.Println(i)

		}
	}
}

func endRange(input string) string {
	return "ok"

}
func main() {

	file, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	r := csv.NewReader(file)
	recods, err := r.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	start := make([]string, len(recods)-1)
	end := make([]string, len(recods)-1)

	for i, line := range recods {
		if i == 0 {
			// skip header line
			continue
		}
		start[i-1] = line[0]
		end[i-1] = line[1]
	}

	for i, _ := range start {
		//fmt.Println(start[i], end[i])

		startInt, err1 := strconv.Atoi(start[i])
		if err1 == nil {
			endInt, err := strconv.Atoi(end[i])
			if err == nil {

				digits := panLength - len(start[i])
				//TODO conditional padding fuction
				fmt.Println("going to iterate", startInt, endInt, digits)
				iterRange(startInt, endInt)

				//			fmt.Println(err, start[i])
			}
			//fmt.Println(end[i])
		}

	}

}
