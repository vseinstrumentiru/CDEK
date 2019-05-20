package cdek

type PvzListFilter string

const (
	// Почтовый индекс города, для которого необходим список ПВЗ
	PvzListFilterCityPostCode PvzListFilter = "citypostcode"

	// Код города по базе СДЭК
	PvzListFilterCityId PvzListFilter = "cityid"

	// Тип пункта выдачи, по умолчанию «PVZ».
	PvzListFilterType PvzListFilter = "type"

	// Код страны по базе СДЭК
	PvzListFilterCountryId PvzListFilter = "countryid"

	// Код страны в формате ISO_3166-1_alpha-2
	PvzListFilterCountryIso PvzListFilter = "countryiso"

	// Код региона по базе СДЭК
	PvzListFilterRegionId PvzListFilter = "regionid"

	// Наличие терминала оплаты («1», «true» - есть; «0», «false» - нет.)
	PvzListFilterHaveCashless PvzListFilter = "havecashless"

	// Разрешен наложенный платеж («1», «true» - да; «0», «false» - нет.)
	PvzListFilterAllowedCod PvzListFilter = "allowedcod"

	// Наличие примерочной («1», «true» - есть; «0», «false» - нет.)
	PvzListFilterIsDressingRoom PvzListFilter = "isdressingroom"

	// Максимальный вес, который может принять ПВЗ
	PvzListFilterWeightMax PvzListFilter = "weightmax"

	// Локализация ПВЗ. По-умолчанию "rus"
	PvzListFilterLang PvzListFilter = "lang"

	// Является ли ПВЗ только пунктом выдачи («1», «true» - да; «0», «false» - нет.)
	PvzListFilterTakeOnly PvzListFilter = "takeonly"
)

const (
	// только склады СДЭК
	TypePvz string = "PVZ"

	// постоматы партнёра
	TypePostomat string = "POSTOMAT"

	// все ПВЗ не зависимо от их типа
	TypeAll string = "ALL"
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
