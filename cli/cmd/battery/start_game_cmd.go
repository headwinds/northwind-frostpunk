/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package battery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// advanceDayCmd represents the advanceDay command
var startGameCmd = &cobra.Command{
    Use:   "startGame",
    Short: "start game",
    Long:  `This command will start the game! Can you survive the winter?`,
    Run: func(cmd *cobra.Command, args []string) {
        startGame()
    },
}

// should considering importing this struct from the API to make it more DRY
// but not sure how to import across projects?! Perhaps its not a good idea?!
// since the CLI is meant to be a separate project from the API
// "github.com/headwinds/northwind-frostpunk/api/controllers/game" and use it game.GameDay

type GameDay struct {
	GameDayNumber 			int         `json:"GameDayNumber,omitempty"`
	Description   			string      `json:"Description,omitempty"`
	TemparatureCelius   	int         `json:"TemparatureCelius,omitempty"`
	MinutesToComplete 		int         `json:"MinutesToComplete,omitempty"`
	Choices					[]string    `json:"Choices,omitempty"`    
}

type HttpResp struct{
    Status      int         `json:"status"`
    Description string      `json:"description"`
    Body        GameDay     `json:"body"`
}



func displayGameDay(gameDay GameDay) {
    fmt.Println("\n")
    fmt.Println("Day ", gameDay.GameDayNumber)
    fmt.Println("Temparature ", gameDay.TemparatureCelius)
    fmt.Printf("\n%s", gameDay.Description)
    fmt.Println("\n")
    fmt.Println("What would you like to do?")
    fmt.Println("\n")
    //fmt.Println("Minutes to Complete: ", gameDay.MinutesToComplete)
    // iterate over choices
    for i, choice := range gameDay.Choices { 
        fmt.Println(i, choice)
    }

}

func startGame(){
    URL := "http://localhost:8080/game/start"

    // Get the data
    response, err := http.Get(URL)
    if err != nil {
        fmt.Println(err)
    }
    response.Header.Add("Accept", "application/json")
    defer response.Body.Close()
    /*
	if response.StatusCode == 200 {

        body, err := ioutil.ReadAll(response.Body)
        if(err != nil){
            fmt.Println(err)
        }
        var gameDay GameDay
        json.Unmarshal(body, &gameDay)

        displayGameDay(gameDay);

    } else {
        fmt.Println("Error: ",  response)
    }*/
    // read json http response
    jsonDataFromHttp, err := ioutil.ReadAll(response.Body)

    if err != nil {
            panic(err)
    }


    var jsonData HttpResp

    err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData) // here!

    if err != nil {
            panic(err)
    }

    if jsonData.Status == 200 {
        gameDay := jsonData.Body
        displayGameDay(gameDay)
    }
}

func init() {
	rootCmd.AddCommand(startGameCmd)
}
