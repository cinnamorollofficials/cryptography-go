package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// PKCS5 padding (works for block size 8, same as PKCS7 for this size)
func pkcs5Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func pkcs5Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("invalid padding size")
	}
	padding := int(data[len(data)-1])
	if padding == 0 || padding > len(data) {
		return nil, fmt.Errorf("invalid padding")
	}
	for i := 0; i < padding; i++ {
		if data[len(data)-1-i] != byte(padding) {
			return nil, fmt.Errorf("invalid padding byte")
		}
	}
	return data[:len(data)-padding], nil
}

// CBC mode encryption using DES
func desEncryptCBC(key, plaintext []byte) (iv, ciphertext []byte, err error) {
	if len(key) != 8 {
		return nil, nil, fmt.Errorf("DES key must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	blockSize := block.BlockSize() // 8
	padded := pkcs5Pad(plaintext, blockSize)

	iv = make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	ciphertext = make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, padded)

	return iv, ciphertext, nil
}

// CBC mode decryption using DES
func desDecryptCBC(key, iv, ciphertext []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, fmt.Errorf("DES key must be 8 bytes")
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	return pkcs5Unpad(plaintext)
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

type ecbEncrypter ecb
type ecbDecrypter ecb

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("ecb encrypter: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("ecb encrypter: dst too short")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst[:x.blockSize], src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("ecb decrypter: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("ecb decrypter: dst too short")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst[:x.blockSize], src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// ECB encrypt (helper)
func desEncryptECB(key, plaintext []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, fmt.Errorf("DES key must be 8 bytes")
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	padded := pkcs5Pad(plaintext, blockSize)

	ciphertext := make([]byte, len(padded))
	enc := (*ecbEncrypter)(newECB(block))
	enc.CryptBlocks(ciphertext, padded)
	return ciphertext, nil
}

// ECB decrypt (helper)
func desDecryptECB(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, fmt.Errorf("DES key must be 8 bytes")
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("ciphertext not full blocks")
	}
	plaintext := make([]byte, len(ciphertext))
	dec := (*ecbDecrypter)(newECB(block))
	dec.CryptBlocks(plaintext, ciphertext)
	return pkcs5Unpad(plaintext)
}

func main() {
	// Example key: 8 bytes (56-bit effective). For real use, ensure proper key management.
	key := []byte("s3cr3tK!") // exactly 8 bytes

	plain := []byte("hadi")

	fmt.Println("Plaintext:", string(plain))

	// --- CBC example ---
	iv, ct, err := desEncryptCBC(key, plain)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nDES-CBC ciphertext (hex):", hex.EncodeToString(ct))
	fmt.Println("DES-CBC iv (hex):", hex.EncodeToString(iv))

	decrypted, err := desDecryptCBC(key, iv, ct)
	if err != nil {
		panic(err)
	}
	fmt.Println("DES-CBC decrypted:", string(decrypted))

	// --- ECB example (educational only) ---
	ct2, err := desEncryptECB(key, plain)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nDES-ECB ciphertext (hex):", hex.EncodeToString(ct2))

	decrypted2, err := desDecryptECB(key, ct2)
	if err != nil {
		panic(err)
	}
	fmt.Println("DES-ECB decrypted:", string(decrypted2))
}
