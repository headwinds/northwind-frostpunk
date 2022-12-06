package products

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/headwinds/northwind-frostpunk/api/types"
	//"golang.org/x/tools/go/expect"
)

func TestGetUrl(t *testing.T) {
	// TODO

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		expectedPath := "http://localhost:8080/products/view?page=1&limit=10";

		if r.URL.Path != expectedPath {
			t.Errorf("Expected to request " + expectedPath +  ", got: %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}

		/*

		productsHttpResp := types.ProductsHttpResp{
		*/
		productA := types.Product{"Chai", 18.00, 39}
		productB := types.Product{"Chang", 19.00, 17}
		productC := types.Product{"Aniseed Syrup", 10.00, 13}
		productsHttpResp := types.ProductsHttpResp{
			Status: 200,
			Description: "OK",
			Body: types.ProductsBody{productA, productB, productC},
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(productsHttpResp)
	}))
	defer server.Close()
  
	value, _ := GetProducts(server.URL)
	if value != "fixed" {
		t.Errorf("Expected 'fixed', got %s", value)
	}

}
