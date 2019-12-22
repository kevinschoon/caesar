package caesar_test

import (
	"testing"

	"github.com/kevinschoon/caesar"
)

const loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In blandit auctor quam et elementum. Nunc in volutpat urna, eget posuere ex. Praesent pretium non tellus eu aliquet. Nullam posuere et lacus nec pulvinar. Sed vestibulum blandit dapibus. Aenean nec dignissim diam, id faucibus nisl. Duis gravida pharetra leo sed euismod. Nullam quis mollis augue, non porttitor nunc. Donec elit ex, ultrices sed tempus ut, aliquam a ante. Morbi volutpat ante sed arcu pharetra semper. Donec dignissim ante in velit ornare fringilla. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Suspendisse ornare, enim non lobortis gravida, metus mauris imperdiet leo, pellentesque aliquam ligula nunc nec magna. Aliquam tincidunt felis nec faucibus elementum."

func TestCaesarCipher(t *testing.T) {
	cipher := caesar.NewCaesarCipher(10, nil)
	cipherText := cipher.MustEncode("hello world")
	if cipherText != "rovvy*&y!vn" {
		t.Fatalf("expected %s, got %s", "rovvy*&y!vn", cipherText)
	}
	decodedText := cipher.MustDecode(cipherText)
	if decodedText != "hello world" {
		t.Fatalf("expected %s, got %s", "hello world", decodedText)
	}
}

func BenchmarkCaesarCipher(b *testing.B) {
	cipher := caesar.NewCaesarCipher(6, nil)
	for n := 0; n < b.N; n++ {
		cipher.MustDecode(cipher.MustEncode(loremIpsum))
	}
}
