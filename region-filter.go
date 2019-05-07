package cdek

//RegionFilter filter key for "List of Regions" request
type RegionFilter string

const (
	//RegionFilterRegionCodeExt  Region code
	RegionFilterRegionCodeExt RegionFilter = "regionCodeExt"

	//RegionFilterRegionCode Region code in the CDEK IS
	RegionFilterRegionCode RegionFilter = "regionCode"

	//RegionFilterRegionFiasGUID Region code according to the Federal Information Address System
	RegionFilterRegionFiasGUID RegionFilter = "regionFiasGuid"

	//RegionFilterCountryCode Country code in the CDEK IS
	RegionFilterCountryCode RegionFilter = "countryCode"

	//RegionFilterCountryCodeExt Code according to the Russian Classifier of Countries of the World
	RegionFilterCountryCodeExt RegionFilter = "countryCodeExt"

	//RegionFilterPage Number of the results page. Default value: 0
	RegionFilterPage RegionFilter = "page"

	//RegionFilterSize Limitation on the number of results displayed. Default value: 1,000
	RegionFilterSize RegionFilter = "size"
)

//RegionFilterBuilder builder for filer for "List of Regions" request
type RegionFilterBuilder struct {
	filter map[RegionFilter]string
}

//AddFilter add filter to set of filters for "List of Regions" request
func (filterBuilder *RegionFilterBuilder) AddFilter(filter RegionFilter, value string) *RegionFilterBuilder {
	if filterBuilder.filter == nil {
		filterBuilder.filter = make(map[RegionFilter]string)
	}

	filterBuilder.filter[filter] = value

	return filterBuilder
}

//Filter compile RegionFilterBuilder for "List of Regions" request
func (filterBuilder *RegionFilterBuilder) Filter() map[RegionFilter]string {
	return filterBuilder.filter
}
