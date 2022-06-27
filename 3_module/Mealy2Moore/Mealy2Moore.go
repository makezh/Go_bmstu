package main

import (
	"fmt"
	"strconv"
)

type pair struct {
	fir int
	sec string
}

func main() {
	// СЧИТЫВАЕМ
	var m, k, n int
	fmt.Scan(&m)
	ins := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&ins[i])
	}
	fmt.Scan(&k)
	t := make([]string, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&t[i])
	}
	fmt.Scan(&n)
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&d[i][j])
		}
	}
	f := make([][]string, n)
	for i := 0; i < n; i++ {
		f[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&f[i][j])
		}
	}

	// СОЗДАЕМ ВЕРШИНЫ
	ver := make([]pair, 0)
	numVer := make(map[pair]int)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var was bool
			alf, _ := strconv.Atoi(f[i][j])
			_, was = numVer[pair{d[i][j], t[alf]}]
			if !was {
				ver = append(ver, pair{d[i][j], t[alf]})
				s := len(numVer)
				numVer[ver[s]] = s
			}
		}
	}

	// ВЫВОДИМ ОТВЕТ (DOT-ФОРМАТ)
	fmt.Print("digraph {\n\trankdir = LR\n")
	for i, v := range ver {
		fmt.Print("\t", i, " [label = \"(", v.fir, ",", v.sec, ")\"]\n")
	}
	for i, v := range ver {
		for j := 0; j < m; j++ {
			alf, _ := strconv.Atoi(f[v.fir][j])
			fmt.Print("\t", i, " -> ", numVer[pair{d[v.fir][j], t[alf]}], " [label = \"", ins[j], "\"]\n")
		}
	}
	fmt.Print("}")
}
