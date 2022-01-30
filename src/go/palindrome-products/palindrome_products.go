package palindrome

import (
	"errors"
	"strconv"
)

// Product holds the product and the factorizations
type Product struct {
	product        int
	Factorizations [][2]int
}

// Products returns the products in a given range
func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	upperProduct := Product{}
	lowerProduct := Product{}

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			val := i * j

			if !isPalindrome(val) {
				continue
			}

			if len(upperProduct.Factorizations) == 0 || val > upperProduct.product {
				upperProduct = Product{product: val, Factorizations: [][2]int{{i, j}}}
			} else if val == upperProduct.product {
				upperProduct.Factorizations = append(upperProduct.Factorizations, [2]int{i, j})
			}

			if len(lowerProduct.Factorizations) == 0 || val < lowerProduct.product {
				lowerProduct = Product{product: val, Factorizations: [][2]int{{i, j}}}
			} else if val == lowerProduct.product {
				lowerProduct.Factorizations = append(lowerProduct.Factorizations, [2]int{i, j})
			}
		}
	}

	if len(upperProduct.Factorizations) == 0 {
		return Product{}, upperProduct, errors.New("no palindromes")
	}

	if upperProduct.product == lowerProduct.product {
		return Product{}, upperProduct, nil
	}

	return lowerProduct, upperProduct, nil
}

func isPalindrome(a int) bool {
	strVal := strconv.Itoa(a)

	for i := 0; i < len(strVal)/2+1; i++ {
		if strVal[i] != strVal[len(strVal)-1-i] {
			return false
		}
	}
	return true
}
