package search

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	tdApiKey     = flag.String("tdApiKey", "", "worldtradingdata api key")
	tdApiVersion = flag.String("tdApiVersion", "", "worldtradingdata api version")
	tdEndPoint   = flag.String("tdEndPoint", "", "worldtradingdata endpoint")
	tdApi        = flag.String("tdApi", "", "worldtradingdata api")
	svcPort      = flag.String("svcPort", "", "service port number")
	timeOut      = flag.Int("timeOut", 0, "http get timeout")
)

func InitEnv() {
	var envStr string
	var envInt int
	var err error
	if envStr = os.Getenv("TD_API_KEY"); envStr != "" {
		tradingDataApiKey = envStr
	}
	if envStr = os.Getenv("TD_API_VERSION"); envStr != "" {
		tradingDataApiVersion = envStr
	}
	if envStr = os.Getenv("TD_ENDPOINT"); envStr != "" {
		tradingDataEndPoint = envStr
	}
	if envStr = os.Getenv("TD_API"); envStr != "" {
		tradingDataApi = envStr
	}
	if envStr = os.Getenv("SVC_PORT"); envStr != "" {
		SVCPort = envStr
	}
	if envStr = os.Getenv("HTTP_TIMEOUT"); envStr != "" {
		if envInt, err = strconv.Atoi(envStr); err == nil {
			httpTimeOut = envInt
		}
	}
}

func InitCmdLine() {
	if *tdApiKey != "" {
		tradingDataApiKey = *tdApiKey
	}
	if *tdApiVersion != "" {
		tradingDataApiVersion = *tdApiVersion
	}
	if *tdEndPoint != "" {
		tradingDataEndPoint = *tdEndPoint
	}
	if *tdApi != "" {
		tradingDataApi = *tdApi
	}
	if *svcPort != "" {
		SVCPort = *svcPort
	}
	if *timeOut > 0 && *timeOut < 60 {
		httpTimeOut = *timeOut
	}
}

func InitTradingDataUrl() {
	tradingDataUrl = fmt.Sprintf("%s/%s/%s", tradingDataApi, tradingDataApiVersion, tradingDataEndPoint)
}
