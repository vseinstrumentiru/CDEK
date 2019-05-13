package cdek

type PvzListFilter string

const (
	FilterCityPostCode   PvzListFilter = "citypostcode"
	FilterCityId         PvzListFilter = "cityid"
	FilterType           PvzListFilter = "type"
	FilterCountryId      PvzListFilter = "countryid"
	FilterRegionId       PvzListFilter = "regionid"
	FilterHaveCashless   PvzListFilter = "havecashless"
	FilterAllowedCod     PvzListFilter = "allowedcod"
	FilterIsDressingRoom PvzListFilter = "isdressingroom"
	FilterWeightMax      PvzListFilter = "weightmax"
	FilterLang           PvzListFilter = "lang"
	FilterTakeOnly       PvzListFilter = "takeonly"
)

const (
	TypePvz      string = "PVZ"
	TypePostomat string = "POSTOMAT"
	TypeAll      string = "ALL"
)

type PvzListFilterBuilder struct {
	filter map[PvzListFilter]string
}

func (pvzListFilter PvzListFilterBuilder) AddFilter(filter PvzListFilter, value string) PvzListFilterBuilder {
	if pvzListFilter.filter == nil {
		pvzListFilter.filter = make(map[PvzListFilter]string)
	}

	pvzListFilter.filter[filter] = value

	return pvzListFilter
}

func (pvzListFilter PvzListFilterBuilder) Filter() map[PvzListFilter]string {
	return pvzListFilter.filter
}
