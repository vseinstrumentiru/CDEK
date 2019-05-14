package cdek

type CityFilter string

const (
	CityFilterRegionCodeExt  CityFilter = "regionCodeExt"  // Код региона
	CityFilterRegionCode     CityFilter = "regionCode"     // Код региона в ИС СДЭК
	CityFilterRegionFiasGuid CityFilter = "regionFiasGuid" // Код региона из ФИАС
	CityFilterPage           CityFilter = "page"           // Номер страницы выборки результата.По умолчанию 0
	CityFilterSize           CityFilter = "size"           // Ограничение выборки результата.По умолчанию 1000
	CityFilterCountryCode    CityFilter = "countryCode"    // Код страны в формате ISO 3166-1 alpha-2
	CityFilterCityName       CityFilter = "cityName"       // Название города
	CityFilterPostcode       CityFilter = "postcode"       // Почтовый индекс
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
