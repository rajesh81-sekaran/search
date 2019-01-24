package search

import "fmt"

var (
	// https://www.worldtradingdata.com/api/v1/stock?symbol=AAPL,MSFT,HSBA.L&api_token=TFxDjEXdqrlzDdnxJMuBTUak9V7gH6YGSo802rQWTEi7IaE9H2zc5LVHvZ4t
	tradingDataApiKey     = "TFxDjEXdqrlzDdnxJMuBTUak9V7gH6YGSo802rQWTEi7IaE9H2zc5LVHvZ4t"
	tradingDataApiVersion = "v1"
	tradingDataEndPoint   = "stock"
	tradingDataApi        = "https://www.worldtradingdata.com/api"
	SVCPort               = "9292"
	tradingDataUrl        = fmt.Sprintf("%s/%s/%s", tradingDataApi, tradingDataApiVersion, tradingDataEndPoint)
	httpTimeOut           = 20
)
