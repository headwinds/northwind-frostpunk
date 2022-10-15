/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package battery

import (
	"fmt"
	"net/http"
    "strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// advanceDayCmd represents the advanceDay command
var startGameCmd = &cobra.Command{
    Use:   "startGame",
    Short: "start game",
    Long:  `This command will start the game! Can you survive the winter?`,
    Run: func(cmd *cobra.Command, args []string) {
        promptStartGameChoice()
    },
}

func callStartGame(){
    URL := "http://localhost:8080/game/start"

    // Get the data
    response, err := http.Get(URL)
    if err != nil {
        fmt.Println(err)
    }
    response.Header.Add("Accept", "application/json")
    defer response.Body.Close()

    if response.StatusCode == 200 {

         //var generic map[string]interface{}
         //err = json.NewDecoder(response.Body).Decode(&generic)            
        fmt.Println("Here comes winter!")
        promptCustom()
    } else {
        fmt.Println("Error: ",  response)
    }

}

func promptStartGameChoice() {
    prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

    if result == "Monday" {
        callStartGame()
    }
}

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

func promptCustom(){
    peppers := []pepper{
		{Name: "Bell Pepper", HeatUnit: 0, Peppers: 0},
		{Name: "Banana Pepper", HeatUnit: 100, Peppers: 1},
		{Name: "Poblano", HeatUnit: 1000, Peppers: 2},
		{Name: "Jalapeño", HeatUnit: 3500, Peppers: 3},
		{Name: "Aleppo", HeatUnit: 10000, Peppers: 4},
		{Name: "Tabasco", HeatUnit: 30000, Peppers: 5},
		{Name: "Malagueta", HeatUnit: 50000, Peppers: 6},
		{Name: "Habanero", HeatUnit: 100000, Peppers: 7},
		{Name: "Red Savina Habanero", HeatUnit: 350000, Peppers: 8},
		{Name: "Dragon’s Breath", HeatUnit: 855000, Peppers: 9},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .HeatUnit | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .HeatUnit | red }})",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details: `
--------- Pepper ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Heat Unit:" | faint }}	{{ .HeatUnit }}
{{ "Peppers:" | faint }}	{{ .Peppers }}`,
	}

	searcher := func(input string, index int) bool {
		pepper := peppers[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Spicy Level",
		Items:     peppers,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose number %d: %s\n", i+1, peppers[i].Name)
}

func init() {
	rootCmd.AddCommand(startGameCmd)
}