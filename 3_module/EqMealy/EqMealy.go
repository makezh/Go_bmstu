package main

import (
	"fmt"
)

// Кратко: минимализируем 2 автомата и сравниваем получившийся ответ

func prtD(delt [][]int) {
	fmt.Print("\tdelt:\n")
	for _, d := range delt {
		for _, dd := range d {
			fmt.Print(dd, " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
func prtO(outs [][]string) {
	fmt.Print("\touts:\n")
	for _, d := range outs {
		for _, dd := range d {
			fmt.Print(dd, " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func read(n int, m int, delt [][]int, outs [][]string) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&delt[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&outs[i][j])
		}
	}
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func unite(q1 int, q2 int, par []int, size []int) {
	if size[q1] < size[q2] {
		swap(&q1, &q2)
	}
	size[q1] += size[q2]
	par[q2] = q1
}

func find(q int, par []int) int {
	if q == par[q] {
		return q
	}
	par[q] = find(par[q], par)
	return par[q]
}

func split1(n int, m int, outs [][]string) (int, []int) {
	par := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = i
		size[i] = 0
	}

	n1 := n
	pi := make([]int, n)
	for q1 := 0; q1 < n-1; q1++ {
		for q2 := q1 + 1; q2 < n; q2++ {
			if find(q1, par) != find(q2, par) {
				eq := true
				for x := 0; x < m; x++ {
					if outs[q1][x] != outs[q2][x] {
						eq = false
						break
					}
				}
				if eq {
					unite(q1, q2, par, size)
					n1--
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		pi[i] = find(i, par)
	}

	return n1, pi
}

func split(n int, m int, delt [][]int, pi []int) int {
	par := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = i
		size[i] = 0
	}

	n1 := n
	for q1 := 0; q1 < n-1; q1++ {
		for q2 := q1 + 1; q2 < n; q2++ {
			if pi[q1] == pi[q2] && find(q1, par) != find(q2, par) {
				eq := true
				for x := 0; x < m; x++ {
					if pi[delt[q1][x]] != pi[delt[q2][x]] {
						eq = false
						break
					}
				}
				if eq {
					unite(q1, q2, par, size)
					n1--
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		pi[i] = find(i, par)
	}

	return n1
}

func AufHon(n int, m int, q0 int, delt [][]int, outs [][]string) (int, [][]int, [][]string) {
	n1, pi := split1(n, m, outs)
	k := n1

	for {
		n1 = split(n, m, delt, pi)
		if n1 == k {
			break
		}
		k = n1
	}

	id := make(map[int]int)
	t := 0
	for a, b := range pi {
		if a == b {
			id[a] = t
			t++
		}
	}

	delt1 := make([][]int, n1)
	outs1 := make([][]string, n1)
	t = 0
	for a, b := range pi {
		if a == b {
			delt1[t] = make([]int, m)
			outs1[t] = make([]string, m)
			for j := 0; j < m; j++ {
				delt1[t][j] = id[pi[delt[a][j]]]
				outs1[t][j] = outs[a][j]
			}
			t++
		}
	}

	return id[pi[q0]], delt1, outs1
}

func avt_dfs(v int, delt1 [][]int, rev []int, order []int, cnt *int) {
	if rev[v] < 0 {
		rev[v] = *cnt
		order[*cnt] = v
		*cnt++
		for _, s := range delt1[v] {
			avt_dfs(s, delt1, rev, order, cnt)
		}
	}
}

func main() {
	var n, m, q0 int
	fmt.Scan(&n, &m, &q0)
	delt := make([][]int, n)
	outs := make([][]string, n)
	for i := 0; i < n; i++ {
		delt[i] = make([]int, m)
		outs[i] = make([]string, m)
	}
	read(n, m, delt, outs)

	var _n, _m, _q0 int
	fmt.Scan(&_n, &_m, &_q0)
	_delt := make([][]int, _n)
	_outs := make([][]string, _n)
	for i := 0; i < _n; i++ {
		_delt[i] = make([]int, _m)
		_outs[i] = make([]string, _m)
	}
	read(_n, _m, _delt, _outs)

	// МИНИМАЛИЗАЦИЯ
	q01, delt1, outs1 := AufHon(n, m, q0, delt, outs)
	_q01, _delt1, _outs1 := AufHon(_n, _m, _q0, _delt, _outs)

	// КАНОНИЗАЦИЯ
	rev := make([]int, n)
	order := make([]int, n)
	for i := 0; i < n; i++ {
		rev[i] = -1
		order[i] = -1
	}
	cnt := 0
	avt_dfs(q01, delt1, rev, order, &cnt)

	_rev := make([]int, _n)
	_order := make([]int, _n)
	for i := 0; i < _n; i++ {
		_rev[i] = -1
		_order[i] = -1
	}
	_cnt := 0
	avt_dfs(_q01, _delt1, _rev, _order, &_cnt)

	// ИЩЕМ ОТВЕТ
	eq := true
	if cnt != _cnt {
		eq = false
	} else {
		for i := 0; i < cnt; i++ {
			for j := 0; j < m; j++ {
				if rev[delt1[order[i]][j]] != _rev[_delt1[_order[i]][j]] ||
					outs1[order[i]][j] != _outs1[_order[i]][j] {
					eq = false
				}
			}
		}
	}

	// ВЫВОДИМ ОТВЕТ
	if eq {
		fmt.Print("EQUAL")
	} else {
		fmt.Print("NOT EQUAL")
	}
}
