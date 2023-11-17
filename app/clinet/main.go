package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8998")
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, _ := reader.ReadString('\n')
		trim := strings.Trim(readString, "\r\n")
		if strings.ToUpper(trim) == "Q" {
			return
		}
		_, err := conn.Write([]byte(trim))
		if err != nil {
			return
		}
		var buf = [512]byte{}
		n, _ := conn.Read(buf[:])
		fmt.Println(string(buf[:n]))
	}

}
