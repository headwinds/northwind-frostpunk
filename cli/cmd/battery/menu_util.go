package battery

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

/*
[]string{"North", "East", "South", "West"}
*/

func PromptGameChoice( items []string) {

    prompt := promptui.Select{
		Label: "Select Day",
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	fmt.Printf("You choose %q\n", result)

    //return result
}

type Product struct {
	Name    string
	Price 	int
	Quantity  	int
}

func PromptSelectProducts() {
    products := []Product{
		{Name: "Bell Pepper", Price: 0, Quantity: 3},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .Price | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Price | red }})",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details: `
--------- Pepper ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Price:" | faint }}	{{ .Price }}
{{ "Quantity:" | faint }}	{{ .Quantity }}`,
	}

	searcher := func(input string, index int) bool {
		product := products[index]
		name := strings.Replace(strings.ToLower(product.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Quantity",
		Items:     products,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose number %d: %s\n", i+1, products[i].Name)

	//return products[i]
}
