package cdek

type PvzListFilter string

const (
	FilterCityPostCode   PvzListFilter = "citypostcode"   // Почтовый индекс города, для которого необходим список ПВЗ
	FilterCityId         PvzListFilter = "cityid"         // Код города по базе СДЭК
	FilterType           PvzListFilter = "type"           // Тип пункта выдачи, по умолчанию «PVZ».
	FilterCountryId      PvzListFilter = "countryid"      // Код страны по базе СДЭК
	FilterCountryIso     PvzListFilter = "countryiso"     // Код страны в формате ISO_3166-1_alpha-2
	FilterRegionId       PvzListFilter = "regionid"       // Код региона по базе СДЭК
	FilterHaveCashless   PvzListFilter = "havecashless"   // Наличие терминала оплаты («1», «true» - есть; «0», «false» - нет.)
	FilterAllowedCod     PvzListFilter = "allowedcod"     // Разрешен наложенный платеж («1», «true» - да; «0», «false» - нет.)
	FilterIsDressingRoom PvzListFilter = "isdressingroom" // Наличие примерочной («1», «true» - есть; «0», «false» - нет.)
	FilterWeightMax      PvzListFilter = "weightmax"      // Максимальный вес, который может принять ПВЗ
	FilterLang           PvzListFilter = "lang"           // Локализация ПВЗ. По-умолчанию "rus"
	FilterTakeOnly       PvzListFilter = "takeonly"       // Является ли ПВЗ только пунктом выдачи («1», «true» - да; «0», «false» - нет.)
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
