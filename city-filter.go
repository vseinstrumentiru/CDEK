package cdek

type CityFilter string

const (
	//CityFilterRegionCodeExt Код региона
	CityFilterRegionCodeExt CityFilter = "regionCodeExt"

	//CityFilterRegionCode Код региона в ИС СДЭК
	CityFilterRegionCode CityFilter = "regionCode"

	//CityFilterRegionFiasGuid Код региона из ФИАС
	CityFilterRegionFiasGuid CityFilter = "regionFiasGuid"

	//CityFilterPage Номер страницы выборки результата.По умолчанию 0
	CityFilterPage CityFilter = "page"

	//CityFilterSize Ограничение выборки результата.По умолчанию 1000
	CityFilterSize CityFilter = "size"

	//CityFilterCountryCode Код страны в формате ISO 3166-1 alpha-2
	CityFilterCountryCode CityFilter = "countryCode"

	//CityFilterCityName Название города
	CityFilterCityName CityFilter = "cityName"

	//CityFilterPostcode Почтовый индекс
	CityFilterPostcode CityFilter = "postcode"
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
