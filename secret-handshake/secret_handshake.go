package secret

// steps is an ordered list of action in the secret handshake.
var steps = []string{"wink", "double blink", "close your eyes", "jump"}

/*Handshake preforms a secret handshake from Mary Poppins based on a code.
the steps are
00001 =  1 = wink
00010 =  2 = double blink
00100 =  4 = close your eyes
01000 =  8 = jump
10000 = 16 =  reverse the order of the steps*/
func Handshake(code int) []string {
	var handshake = make([]string, 0)
	if code <= 0 {
		return handshake
	}

	for s, step := range steps {
		if (1<<uint(s))&code > 0 {
			handshake = append(handshake, step)
		}
	}

	if (1<<uint(len(steps)))&code > 0 {
		reverse(handshake)
	}
	return handshake
}

/*reverse reverses the order of an array in place.*/
func reverse(strings []string) {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
}
