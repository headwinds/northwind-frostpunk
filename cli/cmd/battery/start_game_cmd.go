/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package battery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	//"io/ioutil"
	"net/http"
	//"net/url"

	"github.com/headwinds/northwind-frostpunk/cli/types"
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
	GameDayNumber     int      `json:"GameDayNumber,omitempty"`
	Description       string   `json:"Description,omitempty"`
	TemparatureCelius int      `json:"TemparatureCelius,omitempty"`
	MinutesToComplete int      `json:"MinutesToComplete,omitempty"`
	Choices           []string `json:"Choices,omitempty"`
}

type HttpResp struct {
	Status      int     `json:"status"`
	Description string  `json:"description"`
	Body        GameDay `json:"body"`
}

// not DRY - I think I need common structs file to import from

func displayGameDay(gameDay GameDay) {
	fmt.Println("\n")
	fmt.Println("Day ", gameDay.GameDayNumber)
	fmt.Println("Temparature ", gameDay.TemparatureCelius)
	fmt.Printf("\n%s", gameDay.Description)
	fmt.Println("\n")
	fmt.Println("What would you like to do?")
	fmt.Println("\n")

	decision := PromptGameChoice(gameDay.Choices, "Choose an option")

	fmt.Println("decision: ", decision)

	/*
	   switch(decision) {
	       case "Buy Supplies":
	           nextTurn("1")
	       case "Skip Supplies":
	           nextTurn("2")
	   }*/
	nextTurn(decision)
}

func getUrl(url string) HttpResp {
	// Get the data

	fmt.Println("getUrl called ")

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	response.Header.Add("Accept", "application/json")
	defer response.Body.Close()

	// print the response.Body
	fmt.Println("response.Body: ", response.Body)

	jsonDataFromHttp, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("getUrl resulted in an error BEFORE unmarshall")
		panic(err)
	}

	var jsonData HttpResp

	// print the jsonData
	fmt.Println("jsonData: ", jsonData)

	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData) // here!

	if err != nil {
		fmt.Println("getUrl resulted in an error AFTER unmarshall")
		panic(err)
	}

	return jsonData

}

// this is where I really need generics as the response body is different for each API call
func getUrlProductsResponse(url string) types.ProductsHttpResp {
	// Get the data
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	response.Header.Add("Accept", "application/json")
	defer response.Body.Close()

	fmt.Println("getUrlProductsResponse response.Body: ", response.Body)

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

func startGame() {
	API_URL := "https://northwind-frostpunk.headwinds.repl.co" //os.Getenv("REPL_API_URL")
	fmt.Println("API URL: ", API_URL)

	DBURL := os.Getenv("DBURL")
	REPL_API_URL := os.Getenv("REPL_API_URL")

	fmt.Println("DBURL: ", DBURL)
	fmt.Println("REPL_API_URL: ", REPL_API_URL)

	URL := API_URL + "/game/start" //"http://localhost:8080/game/start"

	fmt.Println("GAME ROUTE URL: ", URL)

	jsonData := getUrl(URL)

	if jsonData.Status == 200 {
		gameDay := jsonData.Body

    displayGameDay(gameDay)
	}
}

func nextTurn(decision string) {
	fmt.Println("CLI - nextTurn")
	//URL := "/game/turn/next?decision=" + url.QueryEscape(decision)   //"http://localhost:8080/game/turn/next?decision=" + url.QueryEscape(decision)
	//URL := "http://localhost:8080/products/view?page=1&limit=10"
	//fmt.Printf("CLI - nextTurn - decision ", decision)

	// use Printf to display the decision
	fmt.Printf("CLI - nextTurn - decision %s", decision)

	/*
		  jsonData := getUrlProductsResponse(URL)


			if jsonData.Status == 200 {
				var options []string
				for _, product := range jsonData.Body {
					option := fmt.Sprintf("%s - $%v", product.ProductName, product.UnitPrice)
					options = append(options, option)
				}
				// can add search https://medium.com/manifoldco/improve-your-command-line-go-application-with-promptui-6c4e6fb5a1bc
				decision := PromptGameChoice(options, "What would you like to buy?")

				fmt.Println("decision: ", decision)
			}*/
}

func init() {
	rootCmd.AddCommand(startGameCmd)
}
