# search

`To download and update the package`
```
go get -u github.com/rajesh81-sekaran/search
```

`To start the service`
```
go run $GOPATH/src/github.com/rajesh81-sekaran/search/service/service.go
```

`To build the service and start the service using the binary`
```
cd $GOPATH/src/github.com/rajesh81-sekaran/search/service/
go build service.go
./service
```

**Without any command line options or environment variables the service starts on the port 9292.**

`Available functionality`
```
Search a given stock symbol, with the symbol being part of the path in the URL.
```
`Yet to be developed`
```
(a) stock_exchange query parameter handling
(b) test cases
```

`Command line flags`
```
-tdApiKey string
    worldtradingdata api key
-tdApiVersion string
    worldtradingdata api version
-tdEndPoint string
    worldtradingdata endpoint
-tdApi string
    worldtradingdata api
-svcPort string
    service port number
-timeOut int
    http get timeout
```

`Environment variables`
```
TD_API_KEY
    worldtradingdata api key
TD_API_VERSION
    worldtradingdata api version
TD_ENDPOINT
    worldtradingdata endpoint
TD_API
    worldtradingdata api
SVC_PORT
    service port number
HTTP_TIMEOUT
    http get timeout
```

`Default values`
```
tradingDataApiKey     = "TFxDjEXdqrlzDdnxJMuBTUak9V7gH6YGSo802rQWTEi7IaE9H2zc5LVHvZ4t"
tradingDataApiVersion = "v1"
tradingDataEndPoint   = "stock"
tradingDataApi        = "https://www.worldtradingdata.com/api"
SVCPort               = "9292"
tradingDataUrl        = "https://www.worldtradingdata.com/api/v1/stock"
httpTimeOut           = 20
```

`Order of priority`
```
(1) Command line options
(2) Environment variables
(3) Default values
```

`Example requests and responses`
`Request1`
`curl -X GET http://localhost:9292/stock/MSFT 2>>/dev/null | json`
`Response1`
```
{
  "symbol": "MSFT",
  "name": "Microsoft Corporation",
  "price": "106.20",
  "close_yesterday": "106.71",
  "currency": "USD",
  "market_cap": "815351750970",
  "volume": "23158848",
  "timezone": "EST",
  "timezone_name": "America/New_York",
  "gmt_offset": "-18000",
  "last_trade_time": "2019-01-24 16:00:01",
  "stock_exchange_long": "NASDAQ Stock Exchange",
  "stock_exchange_short": "NASDAQ"
}
```

`Request2`
`curl -X GET http://localhost:9292/stock/AAPL 2>>/dev/null | json`
`Response2`
```
{
  "symbol": "AAPL",
  "name": "Apple Inc.",
  "price": "152.70",
  "close_yesterday": "153.92",
  "currency": "USD",
  "market_cap": "722240903665",
  "volume": "25434224",
  "timezone": "EST",
  "timezone_name": "America/New_York",
  "gmt_offset": "-18000",
  "last_trade_time": "2019-01-24 16:00:02",
  "stock_exchange_long": "NASDAQ Stock Exchange",
  "stock_exchange_short": "NASDAQ"
}
```

`Request3`
`curl -X GET http://localhost:9292/stock/NOTEXIST 2>>/dev/null | json`
`Response3`
```
"Error! The requested stock(s) could not be found."
```
