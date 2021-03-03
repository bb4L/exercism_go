package variablelengthquantity

import "fmt"

const (
	bitmask    = 127
	checkmask  = 128
	shiftvalue = 7
)

// EncodeVarint encode to []byte
func EncodeVarint(input []uint32) []byte {
	var output []byte

	for _, v := range input {
		length := len(fmt.Sprintf("%b", v))
		seq := make([]byte, length)
		seq[length-1] = byte(v) & bitmask
		v = v >> shiftvalue
		k := length - 2

		for ; v > 0; v = v >> shiftvalue {
			val := byte(v)&bitmask + checkmask
			seq[k] = val
			k--
		}

		res := []byte{}
		keepProceeding := false

		for _, z := range seq {
			if z != 0 || keepProceeding {
				res = append(res, z)
				keepProceeding = true
			}
		}

		if len(res) == 0 {
			res = []byte{byte(0)}
		}

		output = append(output, res...)
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
