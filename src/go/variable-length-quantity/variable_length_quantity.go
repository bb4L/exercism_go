package variablelengthquantity

const (
	bitmask   byte = 127
	checkmask byte = 128
)

// EncodeVarint encode to []byte
func EncodeVarint(input []uint32) []byte {
	var output []byte

	for _, v := range input {
		seq, v := []byte{byte(v) & bitmask}, v>>7
		for ; v > 0; v = v >> 7 {
			val := byte(v)&bitmask + checkmask
			seq = append([]byte{val}, seq...)
		}
		output = append(output, seq...)
	}

	return output
}

// DecodeVarint decode to []uint32
func DecodeVarint(input []byte) ([]uint32, error) {
	output, seq := []uint32{}, []uint32{}
	var val uint32 = 0

	for _, v := range input {
		seq = append(seq, uint32(v&bitmask))
		if (v & checkmask) != checkmask {
			for i, b := range seq {
				shift := (len(seq) - i - 1) * 7
				val += b << shift
			}
			output, seq, val = append(output, val), []uint32{}, 0
		}
	}

	return output, nil
}
