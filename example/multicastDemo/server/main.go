/*
 * @Author: linin00
 * @Date: 2022-06-14 08:35:16
 * @LastEditTime: 2022-06-14 10:21:01
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /doge/example/multicastDemo/server/main.go
 * 天道酬勤
 */
package main

import (
	"encoding/hex"
	"net"

	"github.com/sirupsen/logrus"
)

func main() {
	add, err := net.ResolveUDPAddr("udp", "224.0.0.0:9090")
	if err != nil {
		logrus.Fatalln(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, add)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer listener.Close()
	buf := make([]byte, 1024)
	for {
		n, srcAddr, err := listener.ReadFromUDP(buf)
		if err != nil {
			logrus.Fatalln(err)
		}
		logrus.Println(n, "bytes read from", srcAddr)
		logrus.Println(hex.Dump(buf[:n]))
		_, err = listener.WriteToUDP([]byte("hey, I'm here"), srcAddr)
		if err != nil {
			logrus.Fatalln(err)
		}
	}
}
