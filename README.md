# Mindbody-API-Golang

This is a wrapper around the SOAP API XML contracts available to Mindbody API partners, written in Go.


## Getting Started


### Install packages for each service you wish to use
``` bash
go get -u "github.com/vacoj/Mindbody-API-Golang"
go get -u "github.com/vacoj/Mindbody-API-Golang/siteservice"
go get -u "github.com/vacoj/Mindbody-API-Golang/staffservice"
go get -u "github.com/vacoj/Mindbody-API-Golang/clientservice"
go get -u "github.com/vacoj/Mindbody-API-Golang/classservice"
go get -u "github.com/vacoj/Mindbody-API-Golang/saleservice"
go get -u "github.com/vacoj/Mindbody-API-Golang/appointmentservice"
go get -u "github.com/fiorix/wsdl2go"
go get -u "github.com/fiorix/wsdl2go/soap"
```


### Follow this pattern in order to make requests and parse calls
``` go
func makeGetSitesCall() {
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
```
the above example is also found <a hred="github.com/vacoj/Mindbody-API-Golang/example/Main.go">example/Main.go</a>

