package core

// import (
// 	"log"
// )

type Cipher struct {
	// 编码用的密码
	encodePasswords [5]*Password
	// 解码用的密码
	decodePasswords [5]*Password
}

// 编码原数据
func (cipher *Cipher) encode(bs []byte, m byte) {
	// log.Println("encode m=", m)
	for i, v := range bs {
		bs[i] = cipher.encodePasswords[m][v]
	}
}

// 解码加密后的数据到原数据
func (cipher *Cipher) decode(bs []byte, m byte) {
	// log.Println("decode m=", m)
	for i, v := range bs {
		bs[i] = cipher.decodePasswords[m][v]
	}
}

// 新建一个编码解码器
func NewCipher(encodePasswords [5]*Password) *Cipher {
	chiper := &Cipher{}
	for i, v := range encodePasswords {
		encodePassword := v
		decodePassword := &Password{}
		for i2, v2 := range encodePassword {
			encodePassword[i2] = v2
			decodePassword[v2] = byte(i2)
		}
		chiper.encodePasswords[i] = encodePassword
		chiper.decodePasswords[i] = decodePassword
	}
	return chiper
}
