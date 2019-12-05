package main

import (
	"fmt"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	// var reply string
	// if err = client.Call("test.Hello", "abc", &reply); err != nil {
	// 	panic(err)
	// }

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
