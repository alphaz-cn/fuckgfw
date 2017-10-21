/*
* @Author: happyyi(易罗阳)
* @Date:   2017-10-22 01:59:40
* @Last Modified by:   happyyi(易罗阳)
* @Last Modified time: 2017-10-22 03:33:57
 */
//准备工作:
//ncat -l 65535  -k -c "cat" //开启回显服务
//local :10087               //开启本地监听
//      |
// chiper 加密
//      |
//server:xxxxx               //服务器
//      |
// chiper 解密
//      |
// 转发请求到 localhost:65535
//      |
// 获取回包,然后chiper加密
//      |
// 发送给local:10087
//      |
// chiper解密
//      |
// 转发给请求方,main()

package main

import (
	"log"
	"net"
)

func main() {
	//以下模拟一个Socks5客户端,chrome的插件就干了这个事情
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":10087")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")
	buf := make([]byte, 256)
	data := []byte{5, 1, 0}
	conn.Write(data)
	n, err := conn.Read(buf)
	log.Println(buf[:n])

	conn.Write([]byte{5, 1, 0, 3, 9, 49, 50, 55, 46, 48, 46, 48, 46, 49, 0XFF, 0XFF})
	n, err = conn.Read(buf)
	log.Println(buf[:n])

	conn.Write([]byte{1, 2, 3, 4})
	n, err = conn.Read(buf)
	log.Println(buf[:n])
}

/*
基本原理参考了:https://uname.github.io/2016/04/15/socks5-tcp-connect/
2017/10/22 03:25:09 begin dial...
2017/10/22 03:25:09 dial ok
2017/10/22 03:25:09 [5 0]
2017/10/22 03:25:09 [5 0 0 1 0 0 0 0 0 0]
2017/10/22 03:25:09 [1 2 3 4]
*/
