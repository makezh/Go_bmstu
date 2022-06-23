package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var example string

const LEN = 5

func main() {
	cin := bufio.NewScanner(os.Stdin)
	cin.Scan()
	example = cin.Text()
	fmt.Printf("%d", Solve())
}

func Solve() (answer int) {
	answer = 0
	pos := strings.Index(example, ")")
	for pos >= 0 {
		answer++
		pos -= LEN - 1
		substr := example[pos : pos+LEN]
		replaceString := string(rune('A' + pos))
		example = strings.Replace(example, substr, replaceString, -1)
		pos = strings.Index(example, ")")
	}
	return
}
