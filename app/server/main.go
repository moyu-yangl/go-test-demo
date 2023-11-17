package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, _ := net.Listen("tcp", "localhost:8998")
	for {
		conn, _ := listen.Accept()
		go func(c net.Conn) {
			defer c.Close()
			for {
				reader := bufio.NewReader(c)
				var buf [1024 * 8]byte
				n, err := reader.Read(buf[:])
				if err != nil {
					fmt.Printf("err :%v", err)
					break
				}
				rec := string(buf[:n])
				fmt.Printf("data : %v", rec)
				conn.Write([]byte("result: " + rec))
			}
		}(conn)
	}
}
