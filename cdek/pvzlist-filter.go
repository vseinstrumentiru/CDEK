package cdek

type PvzListFilter string

const (
	PvzListFilterCityPostCode   PvzListFilter = "citypostcode"   // Почтовый индекс города, для которого необходим список ПВЗ
	PvzListFilterCityId         PvzListFilter = "cityid"         // Код города по базе СДЭК
	PvzListFilterType           PvzListFilter = "type"           // Тип пункта выдачи, по умолчанию «PVZ».
	PvzListFilterCountryId      PvzListFilter = "countryid"      // Код страны по базе СДЭК
	PvzListFilterCountryIso     PvzListFilter = "countryiso"     // Код страны в формате ISO_3166-1_alpha-2
	PvzListFilterRegionId       PvzListFilter = "regionid"       // Код региона по базе СДЭК
	PvzListFilterHaveCashless   PvzListFilter = "havecashless"   // Наличие терминала оплаты («1», «true» - есть; «0», «false» - нет.)
	PvzListFilterAllowedCod     PvzListFilter = "allowedcod"     // Разрешен наложенный платеж («1», «true» - да; «0», «false» - нет.)
	PvzListFilterIsDressingRoom PvzListFilter = "isdressingroom" // Наличие примерочной («1», «true» - есть; «0», «false» - нет.)
	PvzListFilterWeightMax      PvzListFilter = "weightmax"      // Максимальный вес, который может принять ПВЗ
	PvzListFilterLang           PvzListFilter = "lang"           // Локализация ПВЗ. По-умолчанию "rus"
	PvzListFilterTakeOnly       PvzListFilter = "takeonly"       // Является ли ПВЗ только пунктом выдачи («1», «true» - да; «0», «false» - нет.)
)

const (
	TypePvz      string = "PVZ"      // только склады СДЭК
	TypePostomat string = "POSTOMAT" // постоматы партнёра
	TypeAll      string = "ALL"      // все ПВЗ не зависимо от их типа
)

type PvzListFilterBuilder struct {
	filter map[PvzListFilter]string
}

func (filterBuilder *PvzListFilterBuilder) AddFilter(filter PvzListFilter, value string) *PvzListFilterBuilder {
	if filterBuilder.filter == nil {
		filterBuilder.filter = make(map[PvzListFilter]string)
	}

	filterBuilder.filter[filter] = value

	return filterBuilder
}

func (filterBuilder PvzListFilterBuilder) Filter() map[PvzListFilter]string {
	return filterBuilder.filter
}
