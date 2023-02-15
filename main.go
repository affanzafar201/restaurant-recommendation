package main

import (
	"restaurant-recommendation/cmd"
	"restaurant-recommendation/model"
)

func main() {
	cli := cmd.CommandLineInterface{Restaurants: []model.Restaurant{}}
	cli.Init()
	cli.ProcessUsers()
}
