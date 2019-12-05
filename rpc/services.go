package main

import "fmt"

type HelloService struct{}

func (p HelloService) Hello(request string, response *string) error {
	*response = "hello: " + request
	return nil
}

type NameReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type NameResp struct {
	Desc string `json:"desc"`
}

func (p HelloService) Name(req NameReq, resp *NameResp) error {
	*resp = NameResp{Desc: fmt.Sprintf("name: %s, age: %d", req.Name, req.Age)}
	return nil
}
