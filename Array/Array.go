package Array

func Index(v []int, val int) int {
	for pos, n := range v {
		if val == n {
			return pos
		}
	}
	return -1
}

func Contains(v []int, n int) bool {
	for _, val := range v {
		if val == n {
			return true
		}
	}
	return false
}

func Pop(v []int, pos int) []int {
	var nV []int
	if len(v) > 0 {
		if pos == 0 {
			nV = AppendRecursive(nV, v[pos+1:])
		} else if pos == len(v)-1 {
			nV = AppendRecursive(nV, v[0:pos])
		} else {
			nV = AppendRecursive(nV, v[:pos])
			nV = AppendRecursive(nV, v[pos+1:])
		}
		return nV
	}
	return nV
}

func AppendRecursive(Vo []int, Vc []int) []int {
	var nV []int
	for _, val := range Vo {
		nV = append(nV, val)
	}
	for _, val := range Vc {
		nV = append(nV, val)
	}
	return nV
}

func Remove(Vo []int, val int) []int {
	indice := Index(Vo, val)
	return Pop(Vo, indice)
}

func GetRange(V []int, posIni int, posFin int) []int {
	nV := make([]int, posFin-posIni)
	copy(nV, V[posIni:posFin])
	return nV
}
