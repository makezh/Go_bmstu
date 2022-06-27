package main

import (
	"fmt"
	"math/rand"
)

func find(v int, arr []int) int {
	if arr[v] == v {
		return v
	}
	arr[v] = find(arr[v], arr)
	return arr[v]
}

func unite(x int, y int, arr []int) {
	xx := find(x, arr)
	yy := find(y, arr)
	if xx == yy {
		return
	}
	if rand.Intn(32767)%2 == 1 {
		xx, yy = yy, xx
	}
	arr[xx] = yy
}

var ind int
var z int
var helpi []int
var eq int
var count int

func dfs(start int, arr []int, m3 int, in3 [][]int) {
	arr[start] = ind
	ind = ind + 1
	for i := 0; i < m3; i++ {
		if arr[in3[start][i]] == -1 {
			dfs(in3[start][i], arr, m3, in3)
		}
	}
}

func split1(n int, m int, out [][]string) {
	z = n
	help := make([]int, n)
	for i := 0; i < n; i++ {
		help[i] = i
	}

	for i := 0; i < n; i++ {
		for k := i + 1; k < n; k++ {
			if find(i, help) != find(k, help) {
				eq = 1
				for j := 0; j < m; j++ {
					if out[i][j] != out[k][j] {
						eq = 0
						break
					}
				}
				if eq == 1 {
					unite(i, k, help)
					z--
				}
			}
		}

	}

	for i := 0; i < n; i++ {
		helpi[i] = find(i, help)
	}
}

func split(n int, m int, in [][]int) {
	z = n
	help1 := make([]int, n)
	for i := 0; i < n; i++ {
		help1[i] = i
	}
	for i := 0; i < n; i++ {
		for k := i + 1; k < n; k++ {
			if (helpi[i] == helpi[k]) && (find(i, help1) != find(k, help1)) {
				eq := 1
				for j := 0; j < m; j++ {
					w1 := in[i][j]
					w2 := in[k][j]
					if helpi[w1] != helpi[w2] {
						eq = 0
						break
					}
				}
				if eq == 1 {
					unite(i, k, help1)
					z--
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		helpi[i] = find(i, help1)
	}
}

func main() {

	///////// создаем изначальный

	n := 0
	fmt.Scanf("%d", &n)
	m := 0
	fmt.Scanf("%d", &m)
	q := 0
	fmt.Scanf("%d", &q)

	in := make([][]int, n, n)
	for i := 0; i < n; i++ {
		in[i] = make([]int, m, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			num := 0
			fmt.Scanf("%d", &num)
			in[i][j] = num
		}
	}
	out := make([][]string, n, n)
	for i := 0; i < n; i++ {
		out[i] = make([]string, m, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var chr string
			fmt.Scanf("%s", &chr)
			out[i][j] = chr
		}
	}
	/////////////////// Минимизируем

	helpi = make([]int, n)
	z = -1
	zz := -1
	split1(n, m, out)
	for z != zz {
		zz = z
		split(n, m, in)
	}

	help3 := make([]int, n)
	help4 := make([]int, n)
	count = 0
	for i := 0; i < n; i++ {
		if helpi[i] == i {
			help4[count] = i
			help3[i] = count
			count = count + 1
		}
	}

	n3 := z
	m3 := m
	q3 := help3[helpi[q]]

	in3 := make([][]int, n3, n3)
	for i := 0; i < n3; i++ {
		in3[i] = make([]int, m3, m3)
	}
	out3 := make([][]string, n3, n3)
	for i := 0; i < n3; i++ {
		out3[i] = make([]string, m3, m3)
	}

	for i := 0; i < n3; i++ {
		for k := 0; k < m3; k++ {
			in3[i][k] = help3[helpi[in[help4[i]][k]]]
		}
	}

	for i := 0; i < n3; i++ {
		for j := 0; j < m3; j++ {
			var chr string
			chr = out[help4[i]][j]
			out3[i][j] = chr
		}
	}

	///////////////// канон
	///////// создаем новый автомат

	n2 := n3
	m2 := m3

	in2 := make([][]int, n2, n2)
	for i := 0; i < n2; i++ {
		in2[i] = make([]int, m2, m2)
	}

	out2 := make([][]string, n2, n2)
	for i := 0; i < n2; i++ {
		out2[i] = make([]string, m2, m2)
	}

	arr := make([]int, n3)
	for i := 0; i < n3; i++ {
		arr[i] = -1
	}
	ind = 0
	dfs(q3, arr, m3, in3)

	for i := 0; i < n3; i++ {
		if arr[i] != -1 {
			out2[arr[i]] = out3[i]
			for k := 0; k < m3; k++ {
				in2[arr[i]][k] = arr[in3[i][k]]
			}
		}
	}
	n2 = ind

	/////////////////////////  печать

	fmt.Printf("digraph {\n")
	fmt.Printf("    rankdir = LR\n")

	for i := 0; i < n2; i++ {
		for j, x := range in2[i] {
			fmt.Printf("    %d -> %d [label = \"%c(%v)\"]\n", i, x, 97+j, out2[i][j])
		}
	}
	fmt.Printf("}\n")
}
