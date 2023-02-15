package model

type FilterProcessor struct {
	Priority []Filter
}

func (fp *FilterProcessor) Init() {
	var filtersProcessor []Filter
	filtersProcessor = append(filtersProcessor, FeaturedRestaurantFilter{}, TopPrimaryCuisinePrimaryCostBracketFilter{},
		TopPrimaryCuisineSecondaryCostBracketFilter{}, TopSecondaryCuisinePrimaryCostBracketFilter{},
		TopNewlyCreatedFilter{}, RestPrimaryCuisinePrimaryCostBracketFilter{}, RestPrimaryCuisineSecondaryCostBracketFilter{},
		RestSecondaryCuisinePrimaryCostBracketFilter{}, RemainingFilter{})

	fp.Priority = filtersProcessor
}
