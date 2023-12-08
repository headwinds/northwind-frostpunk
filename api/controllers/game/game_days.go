package game

import (
	"encoding/json"
	"fmt"
	"github.com/headwinds/northwind-frostpunk/api/types"
	"net/http"
	//"github.com/headwinds/northwind-frostpunk/api/utils"
)

type GameDay struct {
	GameDayNumber     int
	Description       string
	TemparatureCelius int
	MinutesToComplete int
	Choices           []string
}

type GameDays struct {
	Days []GameDay
}

type GameState struct {
	CurrentGameDayIndex int
}

/*
I need to store game state in a global variable so that it can be accessed by all controllers.

	func Incr() {
	    gameDays.increaseGameDayIndex()
	}

	func GetDescription() string {
	    return gameDays.Description()
	}

	type foo struct {
	    sync.RWMutex
	    gameDayIndex int
	}

	func (gameDays *GameDays) increaseGameDayIndex() {
	    gameDays.Lock()
	    gameDays.currentIndex++
	    gameDays.Unlock()
	}

	func (gameDays *GameDays) count() int {
	    f.RLock()
	    defer f.RUnlock()
	    return f.count
	}

var gameState = &GameState{}
*/
var gameState GameState = GameState{0} // index must start at 0

/*
Each day, a event will take place that our players will need to encounter and deal with.
*/
func GameDayManager() GameDay {

	return GetGameDay(gameState.CurrentGameDayIndex)
}

func GetGameDay(gameDayIndex int) GameDay {

	gameDay1Description := "You are preparing for the launch of your expidition to Akelton, the new lithium mine in northern Ontario. The dropship is ready and waitng for your arrival. You have already secured your ticket, and now need to collect supplies from the online store. You have 10 minutes to get what you need before the dropship leaves."
	gameDay1Choices := []string{"View Supplies", "Skip Supplies"}
	gameDay1 := GameDay{1, gameDay1Description, 0, 10, gameDay1Choices}

	gameDay2Description := "You have arrived at the dropship and are ready to board. You have 10 minutes to get on board before the dropship leaves."
	gameDay2Choices := []string{"View Supplies", "Skip Supplies"}
	gameDay2 := GameDay{2, gameDay2Description, -5, 10, gameDay2Choices}

	gameDays := []GameDay{gameDay1, gameDay2}

	// should have some validation here and panic if incorrect index
	return gameDays[gameDayIndex]
}

func StartGame(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Hit /game/start route")
	gameDay := GameDayManager()

	json.NewEncoder(w).Encode(types.HttpResp{Status: 200, Body: gameDay})
}

func NextTurn(w http.ResponseWriter, r *http.Request) {

	decision := r.URL.Query().Get("decision")
	fmt.Println("decision =>", decision) // I should consider encoding each decision into an id
	jsonData := types.JsonMessageResponse{Type: "Success", Message: "You have chosen to " + decision}
	json.NewEncoder(w).Encode(jsonData)

	/*
	  //products
	  url := "/products/view?page=1&limit=10"
	  jsonData := utils.GetUrlProductsResponse(url)

	  if jsonData.Status == 200 {
	    // we don't want to double nest the response body so we discard the status and description
	    // so we could do
	    // //json.NewEncoder(w).Encode(types.HttpResp{Status: 200, Body: jsonData.Body})
	    // or simply do json.NewEncoder(w).Encode(jsonData)

	    w.Header().Set("Content-Type", "application/json")

	    json.NewEncoder(w).Encode(jsonData)
	    }*/

}
