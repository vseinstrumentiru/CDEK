package cdek

type GetCitiesRes []City

type City struct {
	CityUUID       *string  `json:"cityUuid"`
	CityName       *string  `json:"cityName"`
	CityCode       *int     `json:"cityCode,string"`
	Region         *string  `json:"region"`
	RegionCodeExt  *int     `json:"regionCodeExt,string"`
	RegionCode     *int     `json:"regionCode,string"`
	SubRegion      *string  `json:"subRegion"`
	Country        *string  `json:"country"`
	CountryCode    *string  `json:"countryCode"`
	Latitude       *float64 `json:"latitude"`
	Longitude      *float64 `json:"longitude"`
	Kladr          *string  `json:"kladr"`
	FiasGUID       *string  `json:"fiasGuid"`
	RegionFiasGuid *string  `json:"regionFiasGuid"`
	PaymentLimit   *float64 `json:"paymentLimit"`
}
