// server
package main

import (
	"fmt"
	"io"
	"net"
	//"reflect"
)

const RECV_BUF_LEN = 1024

func main() {
	//fmt.Println("Hello eoe!")
	listener, err := net.Listen("tcp", "0.0.0.0:6666")
	if err != nil {
		panic("error listening:" + err.Error())
	}

	fmt.Println("Starting the server")

	for {
		conn, err := listener.Accept() //accept link
		if err != nil {
			panic("Error accept:" + err.Error())
		}
		fmt.Println("Accepted the Connection :", conn.RemoteAddr())

		go EchoServer(conn)

	}
}

func EchoServer(conn net.Conn) {
	buf := make([]byte, 2)
	defer conn.Close()
	str := ""
	for {
		n, err := conn.Read(buf)
		//fmt.Println(string(buf), reflect.ValueOf(buf).Kind())
		switch err {
		case nil:
			conn.Write(buf[0:n])
			str += string(buf)
			fmt.Println(str)
		case io.EOF:
			fmt.Printf("Warning: End of data: %s \n", err)
			return
		default:
			fmt.Printf("Error: Reading data: %s \n", err)
			return
		}
	}

}
