package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var s []rune // строка с выражением
var pos int  // позиция в строке

func main() {
	cin := bufio.NewScanner(os.Stdin)
	cin.Scan()
	s = []rune(cin.Text())
	fmt.Printf("%d", Solve())
}

func Solve() (answer int) {
	pos++
	if pos >= len(s) { // обрабатываем позицию
		answer = 0
		return
	}

	if unicode.IsDigit(s[pos]) { // обрабатываем цифры
		answer = int(s[pos] - '0')
		return
	}

	switch s[pos] { // обрабатываем операции
	case '+':
		answer = Solve() + Solve()
		return
	case '-':
		answer = Solve() - Solve()
		return
	case '*':
		answer = Solve() * Solve()
		return
	case '(':
		answer = Solve()
		return
	case ')':
		answer = Solve()
		return
	case ' ':
		answer = Solve()
		return
	default:
		answer = Solve() // если встретится лишний символ - просто пропустим его
		return
	}
}
