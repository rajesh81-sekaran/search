# search

`To download and update the package`
```
go get -u github.com/rajesh81-sekaran/search
```

`To start the service`
```
go run $GOPATH/src/github.com/rajesh81-sekaran/search/service/service.go
```

To build the service and start the service using the binary:
cd $GOPATH/src/github.com/rajesh81-sekaran/search/service/
go build service.go
./service

Without any command line options or environment variables the service starts on the port 9292.

Available functionality:
search a given stock symbol, with the symbol being part of the path in the URL.
Yet to be developed:
(a) stock_exchange query parameter handling
(b) test cases

The service supprts the following command line flags:
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

The following environment variables are supported:
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

The following are the default values:
tradingDataApiKey     = "TFxDjEXdqrlzDdnxJMuBTUak9V7gH6YGSo802rQWTEi7IaE9H2zc5LVHvZ4t"
tradingDataApiVersion = "v1"
tradingDataEndPoint   = "stock"
tradingDataApi        = "https://www.worldtradingdata.com/api"
SVCPort               = "9292"
tradingDataUrl        = "https://www.worldtradingdata.com/api/v1/stock"
httpTimeOut           = 20

The following is the order of priority
(1) Command line options
(2) Environment variables
(3) Default values
