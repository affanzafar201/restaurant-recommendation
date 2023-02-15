package model

import "fmt"

type RestaurantSlice []Restaurant

func (r RestaurantSlice) PrintRestaurant() {
	for i := range r {
		fmt.Printf("%d. %s\n", i+1, r[i].Id)
	}
}

func (RestaurantSlice) DifferenceBetweenSlices(restaurants []Restaurant, filteredRestaurants []Restaurant) []Restaurant {
	var remainingRestaurants []Restaurant
	filteredRestaurantsMap := make(map[string]struct{})
	for i := range filteredRestaurants {
		filteredRestaurantsMap[filteredRestaurants[i].Id] = struct{}{}
	}

	for i := range restaurants {
		if _, found := filteredRestaurantsMap[restaurants[i].Id]; !found {
			remainingRestaurants = append(remainingRestaurants, restaurants[i])
		}
	}
	return remainingRestaurants
}
