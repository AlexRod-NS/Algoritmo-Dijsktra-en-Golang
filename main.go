package main

import (
	"fmt"
	"sort"

	"./Array"
)

type Nodo struct {
	ni, nf, p int
}

func (self *Nodo) nodoMin(G []Nodo) int {
	N := self.getNodos(G)
	min := N[0]
	for i := 1; i < len(N); i++ {
		if N[i] < min {
			min = N[i]
		}
	}
	return min
}
func (self *Nodo) getNodos(G []Nodo) []int {
	var N []int
	for _, i := range G {
		if self.not_in(i.ni, N) {
			N = append(N, i.ni)
		}
		if self.not_in(i.nf, N) {
			N = append(N, i.nf)
		}
	}
	return N
}
func (self Nodo) not_in(n int, N []int) bool {
	for _, val := range N {
		if val == n {
			return false
		}
	}
	return true
}

type Djisktra struct {
}

func (this Djisktra) Get(G []Nodo) ([]int, []int) {
	Q := new(Nodo).getNodos(G)
	sort.IntsAreSorted(Q)

	Q_aux := Array.GetRange(Q, 0, len(Q))

	d := []int{0}
	for i := 0; i < len(Q)-1; i++ {
		d = append(d, 999999)
	}

	var Prev []int
	for i := 0; i < len(Q); i++ {
		Prev = append(Prev, new(Nodo).nodoMin(G))
	}

	var min int
	for len(Q) != 0 {
		min = d[Array.Index(Q_aux, Q[0])]
		for i := 0; i < len(d); i++ {
			if Array.Contains(Q, Q_aux[i]) {
				if d[i] < min {
					min = d[i]
				}
			}
		}
		indx_min := Array.Index(d, min)

		Q = Array.Remove(Q, Q_aux[indx_min])

		for _, val := range G {
			if val.ni == Q_aux[indx_min] {
				if d[Array.Index(Q_aux, val.nf)] > d[indx_min]+val.p {
					d[Array.Index(Q_aux, val.nf)] = d[indx_min] + val.p
					Prev[Array.Index(Q_aux, val.nf)] = Q_aux[indx_min]
				}
			}
		}
	}
	return d, Prev
}

func main() {
	var G []Nodo
	G = append(G, Nodo{1, 2, 10}, Nodo{1, 3, 3}, Nodo{2, 3, 3},
		Nodo{2, 4, 2}, Nodo{3, 2, 4}, Nodo{3, 4, 8},
		Nodo{3, 5, 2}, Nodo{4, 5, 9})
	d, Prev := new(Djisktra).Get(G)
	fmt.Println("d: ", d)
	fmt.Println("Prev: ", Prev)
}
