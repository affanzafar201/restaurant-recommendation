package model

import (
	"sort"

	"restaurant-recommendation/utils"
)

var filtersProcessor = new(FilterProcessor)

func init() {
	filtersProcessor.Init()
}

type User struct {
	Cuisines              []CuisineTracking
	CostBracket           []CostTracking
	PrimaryCuisines       map[int]CuisineTracking
	SecondaryCuisines     map[int]CuisineTracking
	PrimaryCostBrackets   map[int]CostTracking
	SecondaryCostBrackets map[int]CostTracking
}

func NewUser(cuisines []CuisineTracking, costBracket []CostTracking) *User {
	u := &User{
		Cuisines:    cuisines,
		CostBracket: costBracket,
	}

	sort.Slice(u.Cuisines, func(i, j int) bool {
		return u.Cuisines[i].NoOfOrders > u.Cuisines[j].NoOfOrders
	})

	sort.Slice(u.CostBracket, func(i, j int) bool {
		return u.CostBracket[i].NoOfOrders > u.CostBracket[j].NoOfOrders
	})

	u.PrimaryCuisines, u.SecondaryCuisines = make(map[int]CuisineTracking), make(map[int]CuisineTracking)
	if len(u.Cuisines) > 0 {
		u.PrimaryCuisines[u.Cuisines[0].Type] = u.Cuisines[0]
	}

	for i := 1; i < len(u.Cuisines) || i < 3; i++ {
		u.SecondaryCuisines[u.Cuisines[i].Type] = u.Cuisines[i]
	}

	u.PrimaryCostBrackets, u.SecondaryCostBrackets = make(map[int]CostTracking), make(map[int]CostTracking)
	if len(u.CostBracket) > 0 {
		u.PrimaryCostBrackets[u.CostBracket[0].Type] = u.CostBracket[0]
	}

	for i := 1; i < len(u.CostBracket) || i < 3; i++ {
		u.SecondaryCostBrackets[u.CostBracket[i].Type] = u.CostBracket[i]
	}

	return u
}

func (u *User) GetRestaurantRecommendation(restaurants []Restaurant) []Restaurant {

	var recommendedRestaurants []Restaurant
	remainingRestaurants := restaurants
	for i := range filtersProcessor.Priority {
		var filteredRestaurants = filtersProcessor.Priority[i].GetAvailableRestaurants(*u, remainingRestaurants)
		recommendedRestaurants = append(recommendedRestaurants, filteredRestaurants...)
		remainingRestaurants = RestaurantSlice{}.DifferenceBetweenSlices(remainingRestaurants, filteredRestaurants)
	}

	recommendedRestaurants = utils.FetchTopNElements(recommendedRestaurants, 10)
	return recommendedRestaurants
}
