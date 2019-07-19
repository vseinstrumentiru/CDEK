# GO SDK for CDEK API
[![GoDoc reference](https://godoc.org/github.com/vseinstrumentiru/CDEK?status.svg)](https://godoc.org/github.com/vseinstrumentiru/CDEK) 
[![Go Report Card](https://goreportcard.com/badge/github.com/vseinstrumentiru/CDEK)](https://goreportcard.com/report/github.com/vseinstrumentiru/CDEK)
[![GitHub release](https://img.shields.io/github/release/vseinstrumentiru/cdek.svg)](https://github.com/vseinstrumentiru/CDEK/releases)

The Go language implementation of SDK for [integration with CDEK](https://www.cdek.ru/clients/integrator.html)

----
Installation
------------

To install this package, you need to install Go and setup your Go workspace on
your computer. The simplest way to install the library is to run:

```
$ go get github.com/vseinstrumentiru/cdek
```
With Go module support (Go 1.11+), simply `import "github.com/vseinstrumentiru/cdek"` in
your source code and `go [build|run|test]` will automatically download the
necessary dependencies ([Go modules
ref](https://github.com/golang/go/wiki/Modules)).

Documentation
-------------
- See [godoc](https://godoc.org/github.com/vseinstrumentiru/CDEK) for package and API
  descriptions and examples.

Example
-------------
You cat get test `clientAccount` and `clientSecurePassword` from [the official CDEK documentation](https://confluence.cdek.ru/pages/viewpage.action?pageId=20264477#DataExchangeProtocol(v1.5)-TestAccount)
```
import "github.com/vseinstrumentiru/cdek"
...

client := cdek.NewClient("https://integration.edu.cdek.ru/").
    SetAuth(clientAccount, clientSecurePassword)

cities, err := client.GetCities(map[cdek.CityFilter]string{
    cdek.CityFilterPage: "1",
})
```
