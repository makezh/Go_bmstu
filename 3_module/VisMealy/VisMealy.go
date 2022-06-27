package main

import "fmt"

func InitIntArray(n, m int, arr [][]int) {
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m, m)
	}
}

func FillIntArray(n, m int, arr [][]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			num := 0
			fmt.Scanf("%d", &num)
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
			fmt.Scanf("%s", &chr)
			arr[i][j] = chr
		}
	}
}

func main() {
	n := 0
	m := 0
	qq := 0
	fmt.Scanf("%d\n%d\n%d", &n, &m, &qq)

	in := make([][]int, n, n)
	InitIntArray(n, m, in)
	FillIntArray(n, m, in)

	out := make([][]string, n, n)
	InitStrArray(n, m, out)
	FillStrArray(n, m, out)

	PrintAuto(n, m, in, out)
}

func PrintAuto(n, m int, in [][]int, out [][]string) {
	fmt.Printf("digraph {\n")
	fmt.Printf("    rankdir = LR\n")

	for i := 0; i < n; i++ {
		for j, x := range in[i] {
			fmt.Printf("    %d -> %d [label = \"%c(%v)\"]\n", i, x, 'a'+j, out[i][j])
		}
	}
	fmt.Printf("}")
}
