package battery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// an alternative to json marshalling is to use a struct with a NewDecoder
// https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}


func decode() {
	gameDay := GameDay{}
	getJson("http://localhost:8080/game/start", &gameDay)
	fmt.Println("gameDay Description: ", gameDay.Description)
}



