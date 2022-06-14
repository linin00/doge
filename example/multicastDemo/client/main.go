/*
 * @Author: linin00
 * @Date: 2022-06-14 08:27:40
 * @LastEditTime: 2022-06-14 10:21:16
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /doge/example/multicastDemo/client/main.go
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
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	conn, err := net.ListenUDP("udp", srcAddr)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer conn.Close()
	logrus.Println(conn.LocalAddr())
	conn.WriteToUDP([]byte("are you there?"), add)

	buf := make([]byte, 1024)
	length, tarAdd, err := conn.ReadFromUDP(buf)
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Println(length, "bytes read from", tarAdd)
	logrus.Println(hex.Dump(buf[:length]))
}
