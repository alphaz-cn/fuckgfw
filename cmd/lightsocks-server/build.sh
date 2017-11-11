# @Author: happyyi
# @Date:   2017-10-20 21:54:37
# @Last Modified by:   happyyi(易罗阳)
# @Last Modified time: 2017-11-11 13:55:20
# GOOS=linux GOARCH=amd64 go build -o lsocks .
GOOS=linux GOARCH=386 go build -o lsocks_linux_386 .
# GOOS=linux GOARCH=arm64 go build -o lsocks_linux_arm64 .
# GOOS=linux GOARCH=arm GOARM=7 go build -o lsocks_linux_arm7 .
# GOOS=linux GOARCH=arm GOARM=6 go build -o lsocks_linux_arm6 .
# GOOS=linux GOARCH=arm GOARM=5 go build -o lsocks_linux_arm5 .
 GOOS=darwin GOARCH=amd64 go build -o lsocks_macos_amd64 .
# GOOS=windows GOARCH=amd64 go build -o lsocks_windows_amd64.exe .
# GOOS=windows GOARCH=386 go build -o lsocks_windows_386.exe .
mkdir ../../bin/lightsocks-server/
mv lsocks_linux_386 ../../bin/lightsocks-server/
mv lsocks_macos_amd64 ../../bin/lightsocks-server/
