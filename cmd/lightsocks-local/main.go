package main

import (
	"fmt"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
	"github.com/gwuhaolin/lightsocks/local"
	"log"
	"net"
	"time"
)

const (
	DefaultListenAddr = ":10086"
	DeadLine          = "2017010104"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 默认配置
	config := &cmd.Config{
		ListenAddr: DefaultListenAddr,
		VerfyCode:  DeadLine,
	}
	config.ReadConfig()
	config.SaveConfig()

	// 校验码处理
	t, err := core.CheckDeadLine(config.VerfyCode)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 日期处理
	if t.Before(time.Now()) {
		log.Fatalln("ERROR: 使用已到期,请联系相关人员处理!!!")
		return
	}

	// 解析配置
	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	listenAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	remoteAddr, err := net.ResolveTCPAddr("tcp", config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// 启动 local 端并监听
	lsLocal := local.New(password, listenAddr, remoteAddr)
	log.Fatalln(lsLocal.Listen(func(listenAddr net.Addr) {
		log.Println("使用配置：", fmt.Sprintf(`
本地监听地址 listen：
%s
远程服务地址 remote：
%s
密码 password：
%s
	`, listenAddr, remoteAddr, password))
		log.Printf("lightsocks-local:%s 启动成功 监听在 %s\n", version, listenAddr.String())
	}))
}
