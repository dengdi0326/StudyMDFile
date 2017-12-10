package main

import (
	"os"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
)

var commonIIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main(){
	plintext := []byte("My name is ")

	if len(os.Args) > 1 {
		plintext = []byte(os.Args[1])
		fmt.Println("args:", os.Args[1])
	}

	key_test := "astaxie12798akljzmknm.ahkjkljl;k"
	if len(os.Args) > 2 {
		key_test = os.Args[2]
	}

	c, err := aes.NewCipher([]byte(key_test))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	cfb := cipher.NewCFBEncrypter(c, commonIIV)
	cipjertext := make([]byte, len(plintext))
	cfb.XORKeyStream(cipjertext, plintext)
	fmt.Printf("%s=>%x\n", plintext, cipjertext)

	cfb = cipher.NewCFBDecrypter(c, commonIIV)
	cipjertextcopy := make([]byte, len(plintext))
	cfb.XORKeyStream(cipjertextcopy, cipjertext)
	fmt.Printf("%s", cipjertextcopy)
}
