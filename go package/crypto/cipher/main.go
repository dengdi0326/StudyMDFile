package main

import (
	"os"
	"fmt"
	"crypto/aes"
	"crypto/cipher"

	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha1"
)

var commonIIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func Aes(){
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

func main() {
	msg := []byte("second example")

	publicKeyData := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----`

	pubBlock, _ := pem.Decode([]byte(publicKeyData))

	pubKeyValue, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		panic(err)
	}
	pub := pubKeyValue.(*rsa.PublicKey)

	en, err := rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", en)

	en2, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, msg, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", en2)

	privateKeyData := `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`
	priBlock, _ := pem.Decode([]byte(privateKeyData))
	priKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	if err != nil {
		panic(err)
	}

	de, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, en)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(de))
	fmt.Println()

	de2, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, priKey, en2, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(de2))
}