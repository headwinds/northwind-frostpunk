/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package battery

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "battery",
	Short: "Explore the Northwind database with a Frostpunk twist",
	Long: `Explore the Northwind database with a Frostpunk twist:

1. New Game
2. Load Game (not implemented)
3. Exit`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
        createPrompt()
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

type promptContent struct {
    errorMsg string
    label    string
}

/* generic here - result could be a string, int, etc. */

type ValidOptions[T comparable] struct {
	vals []T
}

type CustomType string

// not exactly what I want here... v should be T not string?!
func validateOption[T comparable](v CustomType) bool {

    vals := []string{"1", "2", "3", "4"}
    validOptions := ValidOptions[string]{vals}

    for _, val := range validOptions.vals {
        if string(v) == val {
            return true
        }
    }

    return false;
}


func promptGetInput(pc promptContent) string {
    validate := func(input string) error {
        if len(input) <= 0 {
            return errors.New(pc.errorMsg)
        }
        return nil
    }

    templates := &promptui.PromptTemplates{
        Prompt:  "{{ . }} ",
        Valid:   "{{ . | green }} ",
        Invalid: "{{ . | red }} ",
        Success: "{{ . | bold }} ",
    }

    prompt := promptui.Prompt{
        Label:     pc.label,
        Templates: templates,
        Validate:  validate,
    }

    result, err := prompt.Run()

    //isValid := validateOption(string(result))

    // validate the input
    if (result == "1" || result == "2" || result == "3" || result == "4") {
        fmt.Printf("You selected: %s\n", result)    
    } else {
        fmt.Println("Please provide a valid option.")
        promptGetInput(pc)
    }

    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    } 

 

    return result
}

func createPrompt() string {
	fmt.Println("\n\nNorthwind Frostpunk v0.0.1\n\n")

    fmt.Println("1. New Game")
    fmt.Println("2. Load Game")
    fmt.Println("3. Credits")
    fmt.Println("4. Quit")
    fmt.Println("\n")

    wordPromptContent := promptContent{
        "Please provide a number.",
        "Please select an option number:",
    }
    option := promptGetInput(wordPromptContent)

	return option

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


