/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// advanceDayCmd represents the advanceDay command
var advanceDayCmd = &cobra.Command{
    Use:   "advanceDay",
    Short: "This command will test our API and advance the day",
    Long:  `This get command will call the API running localhost:8080 and return the next day eventually but lets use I a route I have like orders or products.`,
    Run: func(cmd *cobra.Command, args []string) {
        var argName string

        if len(args) >= 1 && args[0] != "" {
            argName = args[0]
        }

        //URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"
		URL := "http://localhost:8080/products"

        fmt.Println("This is the argument that I passed: " + argName)

        // Get the data
        response, err := http.Get(URL)
        if err != nil {
            fmt.Println(err)
        }
		response.Header.Add("Accept", "application/json")
        defer response.Body.Close()

        if response.StatusCode == 200 {

			 var generic map[string]interface{}
			 err = json.NewDecoder(response.Body).Decode(&generic)            
            fmt.Println("Perfect! I should see products: ", generic)
        } else {
            fmt.Println("Error: ",  response)
        }
    },
}

func init() {
	rootCmd.AddCommand(advanceDayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// advanceDayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// advanceDayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
