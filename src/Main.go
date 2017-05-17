package main

import (
	"./contracts/SaleService.go"

	"github.com/fiorix/wsdl2go/soap"
)

func main() {
	cli := soap.Client{
		URL: "https://api.mindbodyonline.com",
		Namespace: hello.Namespace,
	}
	conn := hello.NewService(&cli)
	reply, err := conn.Echo(cli, &hello.EchoRequest{Data: "echo"})
	...
}