package main

import (
	"flag"
	"log"
	"net/http"
	"search"
)

func init() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	search.InitEnv()
	search.InitCmdLine()
	search.InitTradingDataUrl()
	http.HandleFunc("/stock/", search.HandleStock)
	log.Printf("starting service on port %s\n", search.SVCPort)
	http.ListenAndServe(":"+search.SVCPort, nil)
}
