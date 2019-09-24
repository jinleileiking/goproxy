package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

var (
	tcpAddr = flag.String("tcpAddr", ":8888", "listen addr")
)

func main() {
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(*tcpAddr, proxy))
}
