package strand

// MAPPING to convert rna to dna
var MAPPING = map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}

// ToRNA convert dna to RNA
func ToRNA(dna string) string {
	result := ""
	for _, c := range dna {
		result += MAPPING[string(c)]
	}
	return result
}

//from input on exercism
// var mapping = map[rune]string{'G': "C", 'C': "G", 'T': "A", 'A': "U"}

// ToRNA convert dna to RNA
// func ToRNA(dna string) (result string) {
// 	for _, c := range dna {
// 		result += mapping[c]
// 	}
// 	return
// }
