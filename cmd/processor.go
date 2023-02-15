package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"restaurant-recommendation/model"
)

type CommandLineInterface struct {
	Restaurants []model.Restaurant
}

func (cli *CommandLineInterface) Init() {

	file, err := os.Open("restaurants.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		restaurant := new(model.Restaurant)
		_ = json.Unmarshal([]byte(scanner.Text()), restaurant)
		cli.Restaurants = append(cli.Restaurants, *restaurant)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("============================================")
	fmt.Println("==========Available Restaurants =========")
	fmt.Println("============================================")

	r := model.RestaurantSlice(cli.Restaurants)
	r.PrintRestaurant()
}

func (cli *CommandLineInterface) ProcessUsers() {
	users := []*model.User{
		model.NewUser([]model.CuisineTracking{{3, 200}, {4, 25}, {5, 10}},
			[]model.CostTracking{{2, 105}, {3, 100}, {4, 10}}),

		model.NewUser([]model.CuisineTracking{{5, 200}, {4, 25}, {1, 10}},
			[]model.CostTracking{{3, 105}, {2, 100}, {1, 10}}),

		model.NewUser([]model.CuisineTracking{{1, 200}, {3, 25}, {2, 10}},
			[]model.CostTracking{{2, 105}, {4, 100}, {3, 10}}),
	}

	var r model.RestaurantSlice
	fmt.Println()
	fmt.Println("============================================")
	fmt.Println("========= User Recommendation=========")
	fmt.Println("============================================")

	for i := range users {
		fmt.Printf("User: %d\n", i+1)
		recommendedRestaurants := users[i].GetRestaurantRecommendation(cli.Restaurants)
		r = recommendedRestaurants
		r.PrintRestaurant()
		fmt.Println()
	}
}
