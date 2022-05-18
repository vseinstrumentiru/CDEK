package cdek

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

const (
	pvzListURL = "pvzlist/v1/xml"
)

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

//GetPvzList The method is used to load the list of active pickup points, from which the client can pick up its order.
func (c clientImpl) GetPvzList(ctx context.Context, filter map[PvzListFilter]string) ([]*Pvz, error) {
	serverURL, err := url.Parse(c.apiURL)
	if err != nil {
		return nil, err
	}
	serverURL.Path = path.Join(serverURL.Path, pvzListURL)

	queryString := serverURL.Query()
	for filterKey, value := range filter {
		queryString.Set(string(filterKey), value)
	}
	serverURL.RawQuery = queryString.Encode()

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL.String(), nil)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", urlFormEncoded)

	resp, err := xmlReq[pvzList](r)
	if err != nil {
		return nil, err
	}

	return resp.Pvz, nil
}

type pvzList struct {
	Pvz []*Pvz `xml:"Pvz"`
}

//Pvz List of Pickup Points
type Pvz struct {
	Code           *string        `xml:"Code,attr"`
	PostalCode     *string        `xml:"PostalCode,attr"`
	Name           *string        `xml:"Name,attr"`
	CountryCode    *string        `xml:"CountryCode,attr"`
	CountryCodeIso *string        `xml:"countryCodeIso,attr"`
	CountryName    *string        `xml:"CountryName,attr"`
	RegionCode     *string        `xml:"RegionCode,attr"`
	RegionName     *string        `xml:"RegionName,attr"`
	CityCode       *int           `xml:"CityCode,attr"`
	City           *string        `xml:"City,attr"`
	WorkTime       *string        `xml:"WorkTime,attr"`
	Address        *string        `xml:"Address,attr"`
	FullAddress    *string        `xml:"FullAddress,attr"`
	AddressComment *string        `xml:"AddressComment,attr"`
	Phone          *string        `xml:"Phone,attr"`
	Email          *string        `xml:"Email,attr"`
	QqID           *string        `xml:"qqId,attr"`
	Note           *string        `xml:"Note,attr"`
	CoordX         *float64       `xml:"coordX,attr"`
	CoordY         *float64       `xml:"coordY,attr"`
	Type           *string        `xml:"Type,attr"`
	OwnerCode      *string        `xml:"ownerCode,attr"`
	IsDressingRoom *bool          `xml:"IsDressingRoom,attr"`
	HaveCashless   *bool          `xml:"HaveCashless,attr"`
	AllowedCod     *bool          `xml:"AllowedCod,attr"`
	NearestStation *string        `xml:"NearestStation,attr"`
	MetroStation   *string        `xml:"MetroStation,attr"`
	Site           *string        `xml:"Site,attr"`
	OfficeImage    []*OfficeImage `xml:"OfficeImage"`
	WorkTimeY      []*WorkTimeY   `xml:"WorkTimeY"`
	WeightLimit    *WeightLimit   `xml:"WeightLimit"`
}

//OfficeImage All photos of the office (except for a photo showing how to get to it)
type OfficeImage struct {
	URL *string `xml:"url,attr"`
}

//WorkTimeY Opening hours for every day
type WorkTimeY struct {
	Day     *int    `xml:"day,attr"`
	Periods *string `xml:"periods,attr"`
}

//WeightLimit Weight limits for a pickup point (the tag is used only if limits are set)
type WeightLimit struct {
	WeightMin *float64 `xml:"WeightMin,attr"`
	WeightMax *float64 `xml:"WeightMax,attr"`
}
