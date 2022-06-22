package main

import (
	"fmt"
	"math"
)

type Fraction struct {
	Num, Denom int
}

func nod(a int, b int) int {
	mod := 1
	if a == 0 && b == 0 {
		return 0
	}

	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))

	if a == b {
		return a
	}
	if b > a {
		mod = a
		a = b
		b = mod
	}
	if b == 0 {
		return a
	}

	for mod > 0 {
		mod = a % b
		if mod > 0 {
			a = b
			b = mod
		}
	}
	return b
}

func (frac Fraction) reduce() Fraction {
	Nod := nod(frac.Num, frac.Denom)
	if Nod != 0 {
		frac.Num /= Nod
		frac.Denom /= Nod
		if frac.Denom < 0 {
			frac.Num = -frac.Num
			frac.Denom = -frac.Denom
		}
		return frac
	}
	return Fraction{0, 1}
}

func NewFraction(Num, Denom int) (answer Fraction) {
	answer = Fraction{Num: Num, Denom: Denom}
	answer = answer.reduce()
	return
}

func SumFractions(f1 Fraction, f2 Fraction) (answer Fraction) {
	answer = NewFraction(f1.Num*f2.Denom+f2.Num*f1.Denom, f1.Denom*f2.Denom)
	return
}

func DiffFractions(f1 Fraction, f2 Fraction) (answer Fraction) {
	answer = NewFraction(f1.Num*f2.Denom-f2.Num*f1.Denom, f1.Denom*f2.Denom)
	return
}

func MultFractions(f1 Fraction, f2 Fraction) (answer Fraction) {
	answer = NewFraction(f1.Num*f2.Num, f1.Denom*f2.Denom)
	return
}

func DivisFractions(f1 Fraction, f2 Fraction) (answer Fraction) {
	answer = NewFraction(f1.Num*f2.Denom, f1.Denom*f2.Num)
	return
}

func Gauss(matrix [][]Fraction, answers []Fraction, n int) (res []Fraction) {
	res = make([]Fraction, n, n)
	for i := 0; i < n; i++ {
		tmp := matrix[i][i]
		for j := n - 1; j >= i; j-- {
			matrix[i][j] = DivisFractions(matrix[i][j], tmp)
		}
		if tmp.Num == 0 {
			return res
		}
		answers[i] = DivisFractions(answers[i], tmp)
		for j := i + 1; j < n; j++ {
			tmp = matrix[j][i]
			for k := n - 1; k >= i; k-- {
				matrix[j][k] = DiffFractions(matrix[j][k], MultFractions(tmp, matrix[i][k]))
			}
			answers[j] = DiffFractions(answers[j], MultFractions(answers[i], tmp))
		}
	}
	res[n-1] = answers[n-1]
	for i := n - 2; i >= 0; i-- {
		res[i] = answers[i]
		for j := i + 1; j < n; j++ {
			res[i] = DiffFractions(res[i], MultFractions(matrix[i][j], res[j]))
		}
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)
	matrix := make([][]Fraction, n)
	for i := range matrix {
		matrix[i] = make([]Fraction, n)
	}
	answers := make([]Fraction, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var x int
			fmt.Scan(&x)
			matrix[i][j] = NewFraction(x, 1)
		}
		var res int
		fmt.Scan(&res)
		answers[i] = NewFraction(res, 1)
	}

	result := Gauss(matrix, answers, n)

	if result[0].Num == 0 && result[0].Denom == 0 {
		fmt.Println("No solution")
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("%d/%d\n", result[i].Num, result[i].Denom)
		}
	}
}
