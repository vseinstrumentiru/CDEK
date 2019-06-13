package cdek

type RegionFilter string

const (
	RegionFilterRegionCodeExt  RegionFilter = "regionCodeExt"  // Код региона
	RegionFilterRegionCode     RegionFilter = "regionCode"     // Код региона в ИС СДЭК
	RegionFilterRegionFiasGUID RegionFilter = "regionFiasGuid" // Код региона из ФИАС
	RegionFilterCountryCode    RegionFilter = "countryCode"    // Код страны в формате ISO 3166-1 alpha-2
	RegionFilterCountryCodeExt RegionFilter = "countryCodeExt" // Код ОКСМ
	RegionFilterPage           RegionFilter = "page"           // Номер страницы выборки результата.По умолчанию 0
	RegionFilterSize           RegionFilter = "size"           // Ограничение выборки результата.По умолчанию 1000
)

type RegionFilterBuilder struct {
	filter map[RegionFilter]string
}

func (filterBuilder *RegionFilterBuilder) AddFilter(filter RegionFilter, value string) *RegionFilterBuilder {
	if filterBuilder.filter == nil {
		filterBuilder.filter = make(map[RegionFilter]string)
	}

	filterBuilder.filter[filter] = value

	return filterBuilder
}

func (filterBuilder RegionFilterBuilder) Filter() map[RegionFilter]string {
	return filterBuilder.filter
}
