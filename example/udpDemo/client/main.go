/*
 * @Author: linin00
 * @Date: 2022-06-10 13:29:34
 * @LastEditTime: 2022-06-10 14:50:06
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /doge/example/udpDemo/client/main.go
 * 天道酬勤
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0} // 如果是0，内核会自动分配一个端口
	dstAddr := &net.UDPAddr{IP: ip, Port: 9080}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr) // 创建一个connection，相当于一个socket。第一个参数是协议类型
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("hello")) // 向connection中写入数据
	fmt.Printf("<%s>\n", conn.RemoteAddr())
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf) // 读取connection中的数据
	fmt.Printf("<%s> %s\n", dstAddr, buf[:n])
}
