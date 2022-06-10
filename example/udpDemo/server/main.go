/*
 * @Author: linin00
 * @Date: 2022-06-10 13:50:33
 * @LastEditTime: 2022-06-10 14:49:57
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /doge/example/udpDemo/server/main.go
 * 天道酬勤
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9080}) // 第一个参数的udp 是协议名，可以是udp、udp4或udp6
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data) // 将数据复制到用户态，并返回发送者地址。Read方法不会返回remoteAddr
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr) // 向发送者写入数据
		if err != nil {
			fmt.Printf(err.Error())
		}
	}
}
