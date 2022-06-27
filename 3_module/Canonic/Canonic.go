package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var cin = bufio.NewReader(os.Stdin)
var cout = bufio.NewWriter(os.Stdout)

var ind int

func dfs(start int, arr []int, positions []int, m int, in [][]int) { // поиск в глубину
	arr[start] = ind
	positions[ind] = start
	ind++
	for i := 0; i < m; i++ {
		if arr[in[start][i]] == -1 {
			dfs(in[start][i], arr, positions, m, in)
		}
	}
}

func InitIntArray(n, m int, arr [][]int) {
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m, m)
	}
}

func FillIntArray(n, m int, arr [][]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			num := 0
			fmt.Fscan(cin, &num)
			arr[i][j] = num
		}
	}
}

func InitStrArray(n, m int, arr [][]string) {
	for i := 0; i < n; i++ {
		arr[i] = make([]string, m, m)
	}
}

func FillStrArray(n, m int, arr [][]string) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var chr string
			fmt.Fscan(cin, &chr)
			arr[i][j] = chr
		}
	}
}

func main() {

	///////// Создаем первоначальный
	n := 0
	fmt.Fscan(cin, &n)
	m := 0
	fmt.Fscan(cin, &m)
	q := 0
	fmt.Fscan(cin, &q)

	in := make([][]int, n, n)
	InitIntArray(n, m, in)
	FillIntArray(n, m, in)

	out := make([][]string, n, n)
	InitStrArray(n, m, out)
	FillStrArray(n, m, out)

	///////////////// Канонический

	///////// Создаем новый автомат

	n2 := n
	m2 := m
	q2 := 0

	///////////////////////
	positions := make([]int, n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = -1
	}
	ind = 0
	dfs(q, arr, positions, m, in)
	cout.WriteString(strconv.Itoa(n2) + "\n")
	cout.WriteString(strconv.Itoa(m2) + "\n")
	cout.WriteString(strconv.Itoa(q2) + "\n")

	for i := 0; i < n2; i++ {
		for j := 0; j < m2; j++ {
			cout.WriteString(strconv.Itoa(arr[in[positions[i]][j]]) + " ")
		}
		cout.WriteString("\n")
	}

	for i := 0; i < n2; i++ {
		for j := 0; j < m2; j++ {
			cout.WriteString(out[positions[i]][j] + " ")
		}
		cout.WriteString("\n")
	}

	cout.Flush()
}
