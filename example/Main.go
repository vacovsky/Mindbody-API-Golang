package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/fiorix/wsdl2go/soap"
)

// "../contracts/SiteService"
// "github.com/fiorix/wsdl2go/soap"

func main() {
	cli := soap.Client{
		URL:       "https://api.mindbodyonline.com/0_5/SiteService.asmx",
		Namespace: Namespace,
	}
	conn := NewSite_x0020_ServiceSoap(&cli)

	sourceCreds := &SourceCredentials{
		SourceName: sourceName,
		Password:   sourcePass,
		SiteIDs: &ArrayOfInt{
			Int: []int{site},
		},
	}

	req := &GetSitesRequest{
		SourceCredentials: sourceCreds,
	}

	reply, err := conn.GetSites(&GetSites{Request: req})
	if err != nil {
		fmt.Println(err)
	}
	spew.Dump(reply.GetSitesResult.Sites.Site[0])
}
