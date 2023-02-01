package main

import "fmt"

func printState(machine map[string]int) {
	fmt.Printf("\nThe coffee machine has:\n"+
		"%d ml of water\n"+
		"%d ml of milk\n"+
		"%d g of coffee beans\n"+
		"%d disposable cups\n"+
		"$%d of money\n",
		machine["water"], machine["milk"], machine["beans"], machine["cups"], machine["money"])
}

func checkAndRunAction(action string, machine map[string]int) bool {
	var choose string

	switch action {
	case "buy":
		fmt.Println("\nWhat do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, 4 - raf:")
		fmt.Scan(&choose)

		buyCoffee(machine, choose)
	case "fill":
		fillMachine(machine)
	case "take":
		takeMoney(machine)
	case "remaining":
		printState(machine)
	case "exit":
		return false
	default:
		fmt.Println("Error: inappropriate action.\nPlease, try again.")
	}

	return true
}

func buyCoffee(machine map[string]int, choose string) {
	var espresso = map[string]int{
		"water": 250,
		"milk":  0,
		"beans": 16,
		"money": -4,
		"cups":  1,
	}
	var latte = map[string]int{
		"water": 350,
		"milk":  75,
		"beans": 20,
		"money": -7,
		"cups":  1,
	}
	var cappuccino = map[string]int{
		"water": 200,
		"milk":  100,
		"beans": 12,
		"money": -6,
		"cups":  1,
	}
	var raf = map[string]int{
		"water": 325,
		"milk":  125,
		"beans": 15,
		"money": -8,
		"cups":  1,
	}

	switch choose {
	case "1":
		if !checkIngredients(machine, espresso) {
			break
		}
		useIngredients(machine, espresso)
	case "2":
		if !checkIngredients(machine, latte) {
			break
		}
		useIngredients(machine, latte)
	case "3":
		if !checkIngredients(machine, cappuccino) {
			break
		}
		useIngredients(machine, cappuccino)
	case "4":
		if !checkIngredients(machine, raf) {
			break
		}
		useIngredients(machine, raf)
	case "back":
		break
	default:
		fmt.Println("Error: inappropriate action. \nPlease, try again.")
	}
}

func fillMachine(machine map[string]int) {
	var ingredients = make(map[string]int, 5)
	var input int

	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&input)
	ingredients["water"] = input

	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&input)
	ingredients["milk"] = input

	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&input)
	ingredients["beans"] = input

	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scan(&input)
	ingredients["cups"] = input

	addIngredients(machine, ingredients)
}

func takeMoney(machine map[string]int) {
	if machine["money"] > 0 {
		fmt.Printf("I gave you $%d\n", machine["money"])
		machine["money"] = 0
	} else {
		fmt.Println("Sorry, nothing to give you.")
	}

}

func checkIngredients(machine, drink map[string]int) bool {
	for i, _ := range machine {
		if machine[i]-drink[i] < 0 {
			fmt.Printf("Sorry, not enough %s\n", i)
			return false
		}
	}

	fmt.Println("I have enough resources, making you a coffee!")
	return true
}

func useIngredients(machine, drink map[string]int) {
	for i, _ := range machine {
		machine[i] -= drink[i]
	}
}

func addIngredients(machine, ingredients map[string]int) {
	for i, _ := range ingredients {
		machine[i] += ingredients[i]
	}
}

func main() {
	var machine = map[string]int{
		"water": 400,
		"milk":  540,
		"beans": 120,
		"cups":  9,
		"money": 550,
	}
	var action string

	for {
		fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)

		if !checkAndRunAction(action, machine) {
			break
		}
	}
}
