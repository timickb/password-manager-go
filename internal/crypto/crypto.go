package crypto

type Encrypter interface {
	Encrypt(plaintext string) string
	Decrypt(ciphertext string) string
}

type StubEncrypter struct{}

func (e *StubEncrypter) Encrypt(plaintext string) string {
	return plaintext
}

func (e *StubEncrypter) Decrypt(ciphertext string) string {
	return ciphertext
}
