// Command diceware labels its input with numerals from
// 1 to 6 to produce a diceware-style formatted word list.
//
// Usage
//
//   diceware <words
//
// where words contains one word on each line.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	lines, err := readlines(os.Stdin)
	if err != nil {
		log.Fatalln("readlines:", err)
	}
	ndigit := b6digits(len(lines) - 1)
	for i, line := range lines {
		fmt.Println(dicenum(i, ndigit), line)
	}
}

// b6digits returns the number of digits necessary
// to represent n in base 6.
func b6digits(n int) int {
	return int(math.Floor(math.Log(float64(n))/math.Log(6))) + 1
}

// dicenum formats n as a diceware label: in base 6 using the
// numerals 1 through 6 (not 0), padded to b digits.
func dicenum(n, b int) string {
	dig := "123456"
	s := ""
	for i := 0; i < b; i++ {
		s = string(dig[n%6]) + s
		n /= 6
	}
	return s
}

// readlines reads r to EOF and returns a slice of lines
// (not including trailing CR and NL chars).
func readlines(r io.Reader) (lines []string, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
