package caesar

import (
	"errors"
)

var (
	// DefaultCharacterSet contains a default table of
	// supported characters.
	DefaultCharacterSet = []rune{
		' ', '!', '"', '#',
		'$', '%', '&', '\'',
		'(', ')', '*', '+',
		',', '-', '.', '/',
		'[', '\\', ']', '^',
		'_', '`', '{', '|',
		'}', '~',
		'0', '1', '2', '3',
		'4', '5', '6', '7',
		'8', '9', ':', ';',
		'<', '=', '>', '?',
		'@', 'A', 'B', 'C',
		'D', 'E', 'F', 'G',
		'H', 'I', 'J', 'K',
		'L', 'M', 'N', 'O',
		'P', 'Q', 'R', 'S',
		'T', 'U', 'V', 'W',
		'X', 'Y', 'Z', 'a',
		'b', 'c', 'd', 'e',
		'f', 'g', 'h', 'i',
		'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q',
		'r', 's', 't', 'u',
		'v', 'w', 'x', 'y',
		'z',
	}

	// ErrUnsupportedRune indicates an unsupported rune was found
	// in the ciphered message.
	ErrUnsupportedRune = errors.New("unsupported rune")
)

// CaesarCipher is an implementation of the famous basic substitution cipher named after Julius Caesar.
type CaesarCipher struct {
	forward map[rune]rune
	reverse map[rune]rune
}

// NewCaesarCipher creates a CaesarCipher with the configured shift. If characterSet
// is nil the cipher will use the DefaultCharacterSet.
func NewCaesarCipher(shift uint, characterSet []rune) *CaesarCipher {
	if characterSet == nil {
		characterSet = DefaultCharacterSet
	}
	forward := map[rune]rune{}
	reverse := map[rune]rune{}
	for i := 0; i < len(characterSet); i++ {
		position := i + int(shift)
		if position >= len(characterSet) {
			position = (i + int(shift)) - len(characterSet)
		}
		forward[characterSet[i]] = characterSet[position]
		reverse[characterSet[position]] = characterSet[i]
	}
	return &CaesarCipher{
		forward: forward,
		reverse: reverse,
	}
}

// Encode translates the message into a cipher text. If the message contains
// a rune which was not specified in the CharacterSet it will return ErrUnsupportedRune.
func (c CaesarCipher) Encode(message string) (string, error) {
	var encoded string
	for i := 0; i < len(message); i++ {
		r := rune(message[i])
		char, ok := c.forward[r]
		if !ok {
			return "", ErrUnsupportedRune
		}
		encoded += string(char)
	}
	return encoded, nil
}

// MustEncode invokes the Encode method but causes a panic if an error is encountered.
func (c CaesarCipher) MustEncode(message string) string {
	encoded, err := c.Encode(message)
	if err != nil {
		panic(err)
	}
	return encoded
}

// Decode translates the cipher text into it's original state. If the ciphered
// message contains a rune which was not specified in the CharacterSet it will
// return ErrUnsupportedRune.
func (c CaesarCipher) Decode(cipherText string) (string, error) {
	var decoded string
	for i := 0; i < len(cipherText); i++ {
		r := rune(cipherText[i])
		char, ok := c.reverse[r]
		if !ok {
			return "", ErrUnsupportedRune
		}
		decoded += string(char)
	}
	return decoded, nil
}

// MustDecode invokes the Decode method but causes a panic if an error is encountered.
func (c CaesarCipher) MustDecode(message string) string {
	decoded, err := c.Decode(message)
	if err != nil {
		panic(err)
	}
	return decoded
}
