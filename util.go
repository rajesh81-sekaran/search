package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func HandleStock(resp http.ResponseWriter, req *http.Request) {
	symbol := getSymbol(req.URL.Path)
	if symbol == "" {
		return
	}
	url := fmt.Sprintf("%s?symbol=%s&api_token=%s", tradingDataUrl, symbol, tradingDataApiKey)
	sr, err := doHttpGet(url)
	if err != nil {
		fmt.Printf("Error %s in http get request\n", err)
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	if sr.Message != "" {
		if err = json.NewEncoder(resp).Encode(sr.Message); err != nil {
			msg := fmt.Sprintf("Error %s while encoding json\n", err)
			log.Printf("%s\n", msg)
			http.Error(resp, msg, http.StatusInternalServerError)
		}
	}
	for _, data := range sr.Data {
		if err = json.NewEncoder(resp).Encode(data); err != nil {
			msg := fmt.Sprintf("Error %s while encoding json\n", err)
			log.Printf("%s\n", msg)
			http.Error(resp, msg, http.StatusInternalServerError)
		}
	}
}

func getSymbol(path string) string {
	pathSlice := strings.Split(path, "/")
	if len(pathSlice) != 3 {
		return ""
	}
	return pathSlice[2]
}

func doHttpGet(url string) (StockResponse, error) {
	sr := StockResponse{}
	var netClient = &http.Client{
		Timeout: time.Second * time.Duration(httpTimeOut),
	}
	response, err := netClient.Get(url)
	if err != nil {
		log.Printf("Error <%s> in http get\n", err)
		return sr, err
	}
	httpRespBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read body: %+v", err)
		return sr, err
	}
	response.Body.Close()
	if err = json.Unmarshal(httpRespBytes, &sr); err != nil {
		log.Printf("Failed Unmarshalling endpoint response: %+v", err)
		return sr, err
	}
	return sr, nil

}
