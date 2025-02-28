package main

import (
	"errors"
	"log"
)

type CarType string

const (
	Small  CarType = "small"
	Medium CarType = "medium"
	Big    CarType = "big"
)

func (ct CarType) isValid() bool {
	switch ct {
	case Small, Medium, Big:
		return true
	}
	return false
}

func calculateTollCost(carType CarType, balance int) (map[string]int, error) {
	var cost int

	if !carType.isValid() {
		return nil, errors.New("invalid car type")
	}

	switch carType {
	case Small:
		cost = 1000
	case Medium:
		cost = 2000
	case Big:
		cost = 3000
	default:
		return nil, nil
	}

	if balance < cost {
		return nil, errors.New("insufficient balance")
	}

	remainingBalance := balance - cost

	return map[string]int{
		"remainingBalance": remainingBalance,
		"cost":             cost,
	}, nil // why can this value be nil?
}

func main() {
	initialBalance := 20000

	result, err := calculateTollCost(Small, initialBalance)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Result:", result)
}
