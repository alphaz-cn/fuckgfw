package core

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"
)

const (
	BufSize = 1024
	TIMEOUT = 10 * time.Second
)

// 加密传输的 TCP Socket
type SecureSocket struct {
	Cipher     *Cipher
	ListenAddr *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

// 从输入流里读取加密过的数据，解密后把原数据放到bs里
func (secureSocket *SecureSocket) DecodeRead(conn *net.TCPConn, bs []byte, m byte) (n int, err error) {
	// 设置读超时
	conn.SetReadDeadline(time.Now().Add(TIMEOUT))
	n, err = conn.Read(bs)
	if err != nil {
		return
	}
	secureSocket.Cipher.decode(bs[:n], m)
	return
}

// 把放在bs里的数据加密后立即全部写入输出流
func (secureSocket *SecureSocket) EncodeWrite(conn *net.TCPConn, bs []byte, m byte) (int, error) {
	secureSocket.Cipher.encode(bs, m)
	// 设置写超时
	conn.SetWriteDeadline(time.Now().Add(TIMEOUT))
	return conn.Write(bs)
}

// 从src中源源不断的读取原数据加密后写入到dst，直到src中没有数据可以再读取
func (secureSocket *SecureSocket) EncodeCopy(dst *net.TCPConn, src *net.TCPConn, m byte) error {
	//src 远程服务器比如baidu.com
	//dst local端
	buf := make([]byte, BufSize)
	for {
		src.SetReadDeadline(time.Now().Add(TIMEOUT))
		readCount, errRead := src.Read(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			} else {
				return nil
			}
		}
		if readCount > 0 {
			writeCount, errWrite := secureSocket.EncodeWrite(dst, buf[0:readCount], m)
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

func (secureSocket *SecureSocket) LocalEncodeCopy(dst *net.TCPConn, src *net.TCPConn, pm *byte) error {
	//src 远程服务器比如baidu.com
	//dst local端
	// log.Printf("INFO: LocalEncodeCopy %d", *m)
	buf := make([]byte, BufSize)
	for {
		src.SetReadDeadline(time.Now().Add(TIMEOUT))
		readCount, errRead := src.Read(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			} else {
				return nil
			}
		}
		if readCount > 0 {
			if readCount == 3 && buf[0] == 5 && buf[1] == 1 && buf[2] == 0 {
				buf[2] = byte(rand.Intn(5))
				*pm = buf[2]
				writeCount, errWrite := secureSocket.EncodeWrite(dst, buf[0:readCount], 0X00)
				if errWrite != nil {
					return errWrite
				}
				if readCount != writeCount {
					return io.ErrShortWrite
				}
			} else {
				writeCount, errWrite := secureSocket.EncodeWrite(dst, buf[0:readCount], *pm)
				if errWrite != nil {
					return errWrite
				}
				if readCount != writeCount {
					return io.ErrShortWrite
				}
			}
		}
	}
}

// 从src中源源不断的读取加密后的数据解密后写入到dst，直到src中没有数据可以再读取
func (secureSocket *SecureSocket) DecodeCopy(dst *net.TCPConn, src *net.TCPConn, m byte) error {
	//src local端
	//dst 远程服务器比如baidu.com
	buf := make([]byte, BufSize)
	for {
		readCount, errRead := secureSocket.DecodeRead(src, buf, m)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			} else {
				return nil
			}
		}
		if readCount > 0 {
			dst.SetWriteDeadline(time.Now().Add(TIMEOUT))
			writeCount, errWrite := dst.Write(buf[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

// 和远程的socket建立连接，他们之间的数据传输会加密
func (secureSocket *SecureSocket) DialRemote() (*net.TCPConn, error) {
	remoteConn, err := net.DialTCP("tcp", nil, secureSocket.RemoteAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接到远程服务器 %s 失败:%s", secureSocket.RemoteAddr, err))
	}
	return remoteConn, nil
}
