package main

func main() {
	text := []byte("My Super Secret code stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")

	// encrypted data is stored in myfile.data
	encrypt(text, key)
	decrypt(key)
}
