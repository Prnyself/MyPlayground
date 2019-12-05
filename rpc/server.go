package main

import (
	"net"
	"net/rpc"
)

func main() {
	if err := rpc.RegisterName("test", new(HelloService)); err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go rpc.ServeConn(conn)
	}

}
