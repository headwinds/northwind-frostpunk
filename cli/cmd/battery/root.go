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
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Input: %s\n", result)

    return result
}

func createPrompt() string {
	fmt.Printf("Welcome to Northwind Frostpunk!")

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


