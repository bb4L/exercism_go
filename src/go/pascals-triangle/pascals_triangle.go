package pascal

// Triangle return pascal's triangle
func Triangle(n int) (result [][]int) {
	result = [][]int{{1}}
	if n == 1 {
		return
	}

	for i := 2; i <= n; i++ {
		previousResult := result[len(result)-1]
		newResult := make([]int, i)
		newResult[0] = 1
		newResult[i-1] = 1

		for j, v := range previousResult[1 : i/2+i%2] {
			newResult[j+1] = v + previousResult[j]
			newResult[(i-1)-(j+1)] = v + previousResult[j]
		}
		result = append(result, newResult)
	}
	return result
}
