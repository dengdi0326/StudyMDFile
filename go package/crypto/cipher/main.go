package main

import (
	"crypto/aes"
	"encoding/hex"
	"crypto/cipher"
	"fmt"
	"io"
	"crypto/rand"
)

func StreamDecode() {
	key := []byte("example key 1234")

	//返回 解密后的 字符串
	ciphertext, _ := hex.DecodeString("f363f3ccdcb12bb883abf484ba77d9cd7d32b5baecb3d4b1b3e0e4beffdb3ded")
	fmt.Println("first", ciphertext)

	//新建一个 Block 接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("second", block)

	//blocksize 字节块大小--工作在块模式--规定具体大小，一块一块处理
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	fmt.Println()
	fmt.Println("third", iv)

	ciphertext = ciphertext[aes.BlockSize:]
	fmt.Println("/n", ciphertext)

	mode := cipher.NewCBCDecrypter(block, iv)
	fmt.Println(mode)

	mode.CryptBlocks(ciphertext, ciphertext)
	fmt.Println(ciphertext)
	fmt.Printf("%s", ciphertext)
}

func main(){
	//加密
	key := []byte("example key 1234")
	plaintext := []byte("exampleplaintext")

	if  len(plaintext) % aes.BlockSize != 0 {
		panic("plaintext len is wrong")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize + len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[:aes.BlockSize], plaintext)
	fmt.Printf("%x\n", ciphertext)



	// 流数加密，没有长度限制，对比上述代码少了 长度比较 步骤
	key2 := []byte("keyexample123456")
	p := []byte("example")
	b, err := aes.NewCipher(key2)
	if err != nil {
		panic(err)
	}

	c := make([]byte, aes.BlockSize + len(p))
	iv = c[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	m := cipher.NewCFBEncrypter(b, iv)
	m.XORKeyStream(c[:aes.BlockSize], p)
	fmt.Printf("%x\n", c)

	//解密
	a, _:= hex.DecodeString("09a3fef9e0a1fa089f9a3566be20a43500000000000000")
	iv = a[:aes.BlockSize]
	a = a[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(b, iv)
	stream.XORKeyStream(a, a)
	fmt.Println(a)
}