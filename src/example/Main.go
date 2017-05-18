package main

import (
	"../contracts/siteservice"
	"github.com/fiorix/wsdl2go/soap"
	// "../contracts/SiteService"
	// "github.com/fiorix/wsdl2go/soap"
)

func main() {
	cli := soap.Client{
		URL:       "https://api.mindbodyonline.com",
		Namespace: siteservice.Namespace,
	}
	conn := siteservice.NewService(&cli)
	reply, err := conn.Echo(cli, &siteservice.EchoRequest{Data: "echo"})
}
