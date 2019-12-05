package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	req := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "lance", Age: 28,
	}

	type Resp struct {
		Desc string `json:"desc"`
	}

	resp := &Resp{}
	if err = client.Call("test.Name", req, resp); err != nil {
		panic(err)
	}

	fmt.Println(time.Now().String())
	// fmt.Println(reply)
	fmt.Println(resp.Desc)
}
