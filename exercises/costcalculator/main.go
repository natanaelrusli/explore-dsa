package main

import (
	"errors"
	"log"
)

type MenuIngredient struct {
	name   string
	amount int
}

type CookingIngredient struct {
	name  string
	price int
}

func calculateMenuCost(menuName string, ingredients []MenuIngredient, availableIngredients []CookingIngredient) (int, error) {
	totalCost := 0
	found := false

	for _, ingredient := range ingredients {
		for _, availableIngredient := range availableIngredients {
			if ingredient.name == availableIngredient.name {
				totalCost += ingredient.amount * availableIngredient.price
				found = true
			}
		}

		if !found {
			return 0, errors.New("Ingredient not found")
		} else {
			found = false
		}
	}

	log.Println("Menu Name:", menuName)
	log.Println("Total Cost:", totalCost)
	return totalCost, nil
}

func main() {
	availableCookingIngredients := []CookingIngredient{
		{
			name:  "Pokcoi",
			price: 2000,
		},
		{
			name:  "Jahe",
			price: 10000,
		},
		{
			name:  "Jengkol",
			price: 5000,
		},
	}

	requiredIngredients := []MenuIngredient{
		{
			name:   "Pokcoi",
			amount: 20,
		},
		{
			name:   "Jengkol",
			amount: 200,
		},
	}

	_, err := calculateMenuCost("Sayur Jengkol", requiredIngredients, availableCookingIngredients)
	if err != nil {
		log.Fatal(err)
	}
}
