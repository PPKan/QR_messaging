package main

import (
    "crypto/aes"
    "fmt"
)

var (
    // We're using a 32 byte long secret key
    secretKey string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

func encrypt(plaintext string) (string) {
    aes, err := aes.NewCipher([]byte(secretKey))
    if err != nil {
        panic(err)
    }

    // Make a buffer the same length as plaintext
    ciphertext := make([]byte, len(plaintext))
    aes.Encrypt(ciphertext, []byte(plaintext))

    return string(ciphertext)
}

func main() {
    // This will successfully encrypt.
    ciphertext := encrypt("This is some sensitive information fdjklafjkldasfjkladsjfklfjdklafjkljewoiriowerio")
    fmt.Printf("Ciphertext: %x \n", ciphertext)

    // This will cause an error since the
    // plaintext is less than 16 bytes.
    ciphertext = encrypt("Hello")
}
