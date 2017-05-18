package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/vacoj/Mindbody-API-Golang/siteservice"
	"github.com/vacoj/wsdl2go/soap"
)

// var (
// 	sourceName = "somecreds"
// 	sourcePass = "somepass"
// 	site       = -99
// )

func main() {
	cli := soap.Client{
		URL:       "https://api.mindbodyonline.com/0_5/SiteService.asmx",
		Namespace: siteservice.Namespace,
	}
	conn := siteservice.NewSite_x0020_ServiceSoap(&cli)
	sourceCreds := &siteservice.SourceCredentials{
		SourceName: sourceName,
		Password:   sourcePass,
		SiteIDs: &siteservice.ArrayOfInt{
			Int: []int{site},
		},
	}

	req := &siteservice.GetSitesRequest{
		SourceCredentials: sourceCreds,
	}

	reply, err := conn.GetSites(&siteservice.GetSites{Request: req})
	if err != nil {
		fmt.Println(err)
	}
	spew.Dump(reply.GetSitesResult.Sites.Site[0])
}
