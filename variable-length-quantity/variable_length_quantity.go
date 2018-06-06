package variablelengthquantity

import "errors"

func DecodeVarint(varint []byte) ([]uint32, error) {
	var tmp uint32 = 0
	var vint byte
	result := []uint32{}
	for _, vint = range varint {
		tmp = (tmp << 7) | uint32(vint&(1<<7-1))
		if vint&(1<<7) == 0 {
			result = append(result, tmp)
			tmp = 0
		}
	}
	if vint&(1<<7) == 0 {
		return result, nil
	}
	return []uint32(nil), errors.New("Not a valid varint")
}

func EncodeVarint(varints []uint32) []byte {
	bytes := []byte{}
	for v := len(varints) - 1; 0 <= v; v-- {
		varint := varints[v]
		bytes = append([]byte{byte(varint & (1<<7 - 1))}, bytes...)
		for varint >>= 7; varint > 0; varint >>= 7 {
			bytes = append([]byte{byte(varint&(1<<7-1) | 1<<7)}, bytes...)
		}
	}
	return bytes
}
