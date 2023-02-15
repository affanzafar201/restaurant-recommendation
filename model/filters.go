package model

import (
	"restaurant-recommendation/utils"
	"sort"
)

type Filter interface {
	GetAvailableRestaurants(User, []Restaurant) []Restaurant
}

type FeaturedRestaurantFilter struct{}
type TopPrimaryCuisinePrimaryCostBracketFilter struct{}
type TopPrimaryCuisineSecondaryCostBracketFilter struct{}
type TopSecondaryCuisinePrimaryCostBracketFilter struct{}
type TopNewlyCreatedFilter struct{}
type RestPrimaryCuisinePrimaryCostBracketFilter struct{}
type RestPrimaryCuisineSecondaryCostBracketFilter struct{}
type RestSecondaryCuisinePrimaryCostBracketFilter struct{}
type RemainingFilter struct{}

var _ Filter = (*FeaturedRestaurantFilter)(nil)
var _ Filter = (*TopPrimaryCuisinePrimaryCostBracketFilter)(nil)
var _ Filter = (*TopPrimaryCuisineSecondaryCostBracketFilter)(nil)
var _ Filter = (*TopSecondaryCuisinePrimaryCostBracketFilter)(nil)
var _ Filter = (*TopNewlyCreatedFilter)(nil)
var _ Filter = (*RestPrimaryCuisinePrimaryCostBracketFilter)(nil)
var _ Filter = (*RestPrimaryCuisineSecondaryCostBracketFilter)(nil)
var _ Filter = (*RestSecondaryCuisinePrimaryCostBracketFilter)(nil)
var _ Filter = (*RemainingFilter)(nil)

func (FeaturedRestaurantFilter) GetAvailableRestaurants(u User, restaurants []Restaurant) []Restaurant {
	var isPrimaryCuisinePrimaryCostFeaturedRestaurantPresent = false
	for i := range restaurants {
		_, primaryCuisineFound := u.PrimaryCuisines[int(restaurants[i].Cuisine)]
		_, primaryCostBracketFound := u.PrimaryCostBrackets[restaurants[i].CostBracket]
		if restaurants[i].IsRecommended && primaryCostBracketFound && primaryCuisineFound {
			isPrimaryCuisinePrimaryCostFeaturedRestaurantPresent = true
			break
		}
	}

	var filteredRestaurants []Restaurant
	for i := range restaurants {
		_, primaryCuisineFound := u.PrimaryCuisines[int(restaurants[i].Cuisine)]
		_, primaryCostBracketFound := u.PrimaryCostBrackets[restaurants[i].CostBracket]
		_, secondaryCuisineFound := u.SecondaryCuisines[int(restaurants[i].Cuisine)]
		_, secondaryCostBracketFound := u.SecondaryCostBrackets[restaurants[i].CostBracket]
		if isPrimaryCuisinePrimaryCostFeaturedRestaurantPresent {
			if primaryCuisineFound && primaryCostBracketFound {
				filteredRestaurants = append(filteredRestaurants, restaurants[i])
			}
		} else {
			if (primaryCostBracketFound && secondaryCuisineFound) || (primaryCuisineFound && secondaryCostBracketFound) {
				filteredRestaurants = append(filteredRestaurants, restaurants[i])
			}
		}
	}

	return filteredRestaurants
}

func (TopPrimaryCuisinePrimaryCostBracketFilter) GetAvailableRestaurants(u User, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for i := range restaurants {
		_, primaryCuisineFound := u.PrimaryCuisines[int(restaurants[i].Cuisine)]
		_, primaryCostBracketFound := u.PrimaryCostBrackets[restaurants[i].CostBracket]
		if primaryCuisineFound && primaryCostBracketFound && restaurants[i].Rating >= 4.0 {
			filteredRestaurants = append(filteredRestaurants, restaurants[i])
		}
	}

	return filteredRestaurants
}

func (TopPrimaryCuisineSecondaryCostBracketFilter) GetAvailableRestaurants(u User, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for i := range restaurants {
		_, primaryCuisineFound := u.PrimaryCuisines[int(restaurants[i].Cuisine)]
		_, secondaryCostBracketFound := u.SecondaryCostBrackets[restaurants[i].CostBracket]
		if primaryCuisineFound && secondaryCostBracketFound && restaurants[i].Rating >= 4.5 {
			filteredRestaurants = append(filteredRestaurants, restaurants[i])
		}
	}

	return filteredRestaurants
}

func (TopSecondaryCuisinePrimaryCostBracketFilter) GetAvailableRestaurants(u User, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for i := range restaurants {
		_, secondaryCuisineFound := u.PrimaryCuisines[int(restaurants[i].Cuisine)]
		_, primaryCostBracketFound := u.SecondaryCostBrackets[restaurants[i].CostBracket]
		if secondaryCuisineFound && primaryCostBracketFound && restaurants[i].Rating >= 4.5 {
			filteredRestaurants = append(filteredRestaurants, restaurants[i])
		}
	}

	return filteredRestaurants
}

func (TopNewlyCreatedFilter) GetAvailableRestaurants(u User, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for i := range restaurants {
		if restaurants[i].IsNewlyCreated() {
			filteredRestaurants = append(filteredRestaurants, restaurants[i])
		}
	}

	sort.Slice(filteredRestaurants, func(i, j int) bool {
		return filteredRestaurants[i].Rating < filteredRestaurants[j].Rating
	})

	return utils.FetchTopNElements(filteredRestaurants, 4)
}

func (RestPrimaryCuisinePrimaryCostBracketFilter) GetAvailableRestaurants(u User, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for i := range restaurants {
		_, primaryCuisineFound := u.PrimaryCuisines[int(restaurants[i].Cuisine)]
		_, primaryCostBracketFound := u.PrimaryCostBrackets[restaurants[i].CostBracket]
		if primaryCuisineFound && primaryCostBracketFound && restaurants[i].Rating < 4.0 {
			filteredRestaurants = append(filteredRestaurants, restaurants[i])
		}
	}

	return filteredRestaurants
}

func (RestPrimaryCuisineSecondaryCostBracketFilter) GetAvailableRestaurants(u User, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for i := range restaurants {
		_, primaryCuisineFound := u.PrimaryCuisines[int(restaurants[i].Cuisine)]
		_, secondaryCostBracketFound := u.SecondaryCostBrackets[restaurants[i].CostBracket]
		if primaryCuisineFound && secondaryCostBracketFound && restaurants[i].Rating < 4.5 {
			filteredRestaurants = append(filteredRestaurants, restaurants[i])
		}
	}

	return filteredRestaurants
}

func (RestSecondaryCuisinePrimaryCostBracketFilter) GetAvailableRestaurants(u User, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for i := range restaurants {
		_, secondaryCuisineFound := u.PrimaryCuisines[int(restaurants[i].Cuisine)]
		_, primaryCostBracketFound := u.SecondaryCostBrackets[restaurants[i].CostBracket]
		if secondaryCuisineFound && primaryCostBracketFound && restaurants[i].Rating < 4.5 {
			filteredRestaurants = append(filteredRestaurants, restaurants[i])
		}
	}

	return filteredRestaurants
}

func (RemainingFilter) GetAvailableRestaurants(u User, restaurants []Restaurant) []Restaurant {
	return restaurants
}
