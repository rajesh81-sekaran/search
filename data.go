package search

/*
{
"NASDAQ":{
"symbol":"AAPL",
"name":"Apple Inc.",
"price":"154.94",
"close_yesterday":"154.94",
"currency":"USD",
"market_cap":"732835688367",
"volume":"142022",
"timezone":"EST",
"timezone_name":"America/New_York",
"gmt_offset":"-18000",
"last_trade_time":"2019-01-16 16:00:01"
}
}
Please
*/

type StockResponse struct {
	SymbolsRequested int    `json:"symbols_requested,omitempty"`
	SymbolsReturned  int    `json:"symbols_returned,omitempty"`
	Message          string `json:"message,omitempty"`
	Data             []struct {
		Symbol             string `json:"symbol,omitempty"`
		Name               string `json:"name,omitempty"`
		Price              string `json:"price,omitempty"`
		CloseYesterday     string `json:"close_yesterday,omitempty"`
		Currency           string `json:"currency,omitempty"`
		MarketCap          string `json:"market_cap,omitempty"`
		Volume             string `json:"volume,omitempty"`
		Timezone           string `json:"timezone,omitempty"`
		TimezoneName       string `json:"timezone_name,omitempty"`
		GmtOffset          string `json:"gmt_offset,omitempty"`
		LastTradeTime      string `json:"last_trade_time,omitempty"`
		StockExchangeLong  string `json:"stock_exchange_long,omitempty"`
		StockExchangeShort string `json:"stock_exchange_short,omitempty"`
	} `json:"data"`
}
