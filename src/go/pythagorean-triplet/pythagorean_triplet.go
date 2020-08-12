package pythagorean

import (
	// "fmt"
	"sort"
)

// Triplet struct containing three numbers
type Triplet struct {
	a, b, c int
}

// Range  returns triplets in the range
func Range(min, max int) []Triplet {
	result := []Triplet{}

	for i := min + 2; i <= max; i++ {
		n := float64(i) / 5
		if n-float64(int(n)) != 0 {
			continue
		}

		a := int(3 * n)
		if a < min {
			continue
		}
		b := int(4 * n)
		result = append(result, Triplet{a, b, i})
	}

	return result
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
func Sum(p int) (result []Triplet) {

	// firstKeys := map[int]struct{}{}

	// triplets := Range(1, p)
	// fmt.Println("triplets")
	// fmt.Println(triplets)
	// for i := 0; i < len(triplets); i++ {
	// 	s := triplets[i].a + triplets[i].b + triplets[i].c
	// 	k := float64(p) / float64(s)

	// 	if float64(int(k)) == k {
	// 		factor := int(k)
	// 		_, ok := firstKeys[factor*triplets[i].a]

	// 		if !ok {
	// 			fmt.Println("a")
	// 			fmt.Println(factor * triplets[i].a)
	// 			result = append(result, Triplet{factor * triplets[i].a, factor * triplets[i].b, factor * triplets[i].c})
	// 			firstKeys[factor*triplets[i].a] = struct{}{}
	// 		}
	// 	}
	// }

	m, n := 2, 1

	firstKeys := map[int]struct{}{}

	for {
		for {
			triplet := []int{m*m + n*n, 2 * m * n, m*m - n*n}
			sort.Ints(triplet)

			s := 2*m*m + 2*m*n
			if s > p {
				break
			}
			if s == p {
				_, ok := firstKeys[triplet[0]]

				if !ok {
					result = append(result, Triplet{triplet[0], triplet[1], triplet[2]})
					firstKeys[triplet[0]] = struct{}{}
				}

				break
			}

			k := float64(p) / float64(s)

			if float64(int(k)) == k && k > 1 {
				factor := int(k)
				_, ok := firstKeys[factor*triplet[0]]

				if !ok {
					result = append(result, Triplet{factor * triplet[0], factor * triplet[1], factor * triplet[2]})
					firstKeys[factor*triplet[0]] = struct{}{}
				}
			}

			m++
		}

		n++
		m = n + 1
		if m*m > p/4 {
			break
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].a < result[j].a
	})

	return result

}
