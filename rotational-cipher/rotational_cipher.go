package rotationalcipher

func RotationalCipher(plainText string, distance int) string {
	result := []rune{}
	for _, char := range plainText {
		if 'a' <= char && char <= 'z' {
			char = rune((byte(char)-byte('a')+byte(distance))%26 + byte('a'))
		} else if 'A' <= char && char <= 'Z' {
			char = rune((byte(char)-byte('A')+byte(distance))%26 + byte('A'))
		}
		result = append(result, char)
	}
	return string(result)
}
