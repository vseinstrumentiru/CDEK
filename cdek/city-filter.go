package cdek

type CityFilter string

const (
	FilterRegionCodeExt  CityFilter = "regionCodeExt"  // Код региона
	FilterRegionCode     CityFilter = "regionCode"     // Код региона в ИС СДЭК
	FilterRegionFiasGuid CityFilter = "regionFiasGuid" // Код региона из ФИАС
	FilterPage           CityFilter = "page"           // Номер страницы выборки результата.По умолчанию 0
	FilterSize           CityFilter = "size"           // Ограничение выборки результата.По умолчанию 1000
	FilterCountryCode    CityFilter = "countryCode"    // Код страны в формате ISO 3166-1 alpha-2
	FilterCityName       CityFilter = "cityName"       // Название города
	FilterPostcode       CityFilter = "postcode"       // Почтовый индекс
)

type CityFilterBuilder struct {
	filter map[CityFilter]string
}

func (filterBuilder *CityFilterBuilder) AddFilter(filter CityFilter, value string) *CityFilterBuilder {
	if filterBuilder.filter == nil {
		filterBuilder.filter = make(map[CityFilter]string)
	}

	filterBuilder.filter[filter] = value

	return filterBuilder
}

func (filterBuilder CityFilterBuilder) Filter() map[CityFilter]string {
	return filterBuilder.filter
}
