package util

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// CryptoUtil 加密解密工具类
type CryptoUtil struct {
}

func NewCryptoUtil() *CryptoUtil {
	return &CryptoUtil{}
}

// Md5Encode md5编码只有加密
func (c CryptoUtil) Md5Encode(v interface{}) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v", v))))
}

// RsaEncode rsa使用公钥编码
func (r *Rsa) RsaEncode(v interface{}) string {
	block, _ := pem.Decode([]byte(r.PublicKey))

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	bytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), []byte(fmt.Sprintf("%v", v)))
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// RsaDecode rsa使用私钥解码
func (r *Rsa) RsaDecode(v string) string {
	block, _ := pem.Decode([]byte(r.PrivateKey))
	publicKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	bytes, err := rsa.DecryptPKCS1v15(rand.Reader, publicKey.(*rsa.PrivateKey), []byte(v))
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// Rsa 存放私钥和公钥的对象
type Rsa struct {
	PrivateKey string
	PublicKey  string
}

// RsaCreate 获取Rsa秘钥对象
func (c CryptoUtil) RsaCreate() *Rsa {
	rng := rand.Reader
	generateKey, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		panic(err)
	}
	v, err := x509.MarshalPKCS8PrivateKey(generateKey)
	if err != nil {
		panic(err)
	}
	var blockPrivate = pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: v,
	}
	v, err = x509.MarshalPKIXPublicKey(&generateKey.PublicKey)
	if err != nil {
		panic(err)
	}
	var blockPublic = pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: v,
	}
	return &Rsa{
		PrivateKey: string(pem.EncodeToMemory(&blockPrivate)),
		PublicKey:  string(pem.EncodeToMemory(&blockPublic)),
	}
}
