package cipher

// Cipher reprecenting a cipher
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
