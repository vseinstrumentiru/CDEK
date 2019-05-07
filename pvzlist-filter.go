package cdek

//PvzListFilter key for filtering pvzList
type PvzListFilter string

const (
	//PvzListFilterCityPostCode Почтовый индекс города, для которого необходим список ПВЗ
	PvzListFilterCityPostCode PvzListFilter = "citypostcode"

	//PvzListFilterCityID Код города по базе СДЭК
	PvzListFilterCityID PvzListFilter = "cityid"

	//PvzListFilterType Тип пункта выдачи, по умолчанию «PVZ».
	PvzListFilterType PvzListFilter = "type"

	//PvzListFilterCountryID Код страны по базе СДЭК
	PvzListFilterCountryID PvzListFilter = "countryid"

	//PvzListFilterCountryIso Код страны в формате ISO_3166-1_alpha-2
	PvzListFilterCountryIso PvzListFilter = "countryiso"

	//PvzListFilterRegionID Код региона по базе СДЭК
	PvzListFilterRegionID PvzListFilter = "regionid"

	//PvzListFilterHaveCashless Наличие терминала оплаты («1», «true» - есть; «0», «false» - нет.)
	PvzListFilterHaveCashless PvzListFilter = "havecashless"

	//PvzListFilterAllowedCod Разрешен наложенный платеж («1», «true» - да; «0», «false» - нет.)
	PvzListFilterAllowedCod PvzListFilter = "allowedcod"

	//PvzListFilterIsDressingRoom Наличие примерочной («1», «true» - есть; «0», «false» - нет.)
	PvzListFilterIsDressingRoom PvzListFilter = "isdressingroom"

	//PvzListFilterWeightMax Максимальный вес, который может принять ПВЗ
	PvzListFilterWeightMax PvzListFilter = "weightmax"

	//PvzListFilterLang Локализация ПВЗ. По-умолчанию "rus"
	PvzListFilterLang PvzListFilter = "lang"

	//PvzListFilterTakeOnly Является ли ПВЗ только пунктом выдачи («1», «true» - да; «0», «false» - нет.)
	PvzListFilterTakeOnly PvzListFilter = "takeonly"
)

const (
	//TypePvz только склады СДЭК
	TypePvz string = "PVZ"

	//TypePostomat постоматы партнёра
	TypePostomat string = "POSTOMAT"

	//TypeAll все ПВЗ не зависимо от их типа
	TypeAll string = "ALL"
)

//PvzListFilterBuilder builder for pvzList filter
type PvzListFilterBuilder struct {
	filter map[PvzListFilter]string
}

//AddFilter adds filter for pvzList filter
func (filterBuilder *PvzListFilterBuilder) AddFilter(filter PvzListFilter, value string) *PvzListFilterBuilder {
	if filterBuilder.filter == nil {
		filterBuilder.filter = make(map[PvzListFilter]string)
	}

	filterBuilder.filter[filter] = value

	return filterBuilder
}

//Filter returns complete filter
func (filterBuilder PvzListFilterBuilder) Filter() map[PvzListFilter]string {
	return filterBuilder.filter
}
