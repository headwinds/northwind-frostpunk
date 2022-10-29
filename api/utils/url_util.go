package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/headwinds/northwind-frostpunk/api/types"

	//"strconv"
	"io/ioutil"
)

type HttpResp struct{
    Status      int         `json:"status"`
    Description string      `json:"description"`
    Body        interface{} `json:"body"`
}

func GetUrl(url string) HttpResp {
    // Get the data
    response, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    response.Header.Add("Accept", "application/json")
    defer response.Body.Close()

    jsonDataFromHttp, err := ioutil.ReadAll(response.Body)

    if err != nil {
            panic(err)
    }

    var jsonData HttpResp

    err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData) // here!

    if err != nil {
            panic(err)
    }

    return jsonData

}

// this is where I really need generics as the response body is different for each API call
func GetUrlProductsResponse(url string) types.ProductsHttpResp {
    // Get the data
    response, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    response.Header.Add("Accept", "application/json")
    defer response.Body.Close()

    jsonDataFromHttp, err := ioutil.ReadAll(response.Body)

    if err != nil {
            panic(err)
    }

    var jsonData types.ProductsHttpResp

    err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData) // here!

    if err != nil {
            panic(err)
    }

    return jsonData

}