package cdek

type PvzList struct {
	Pvz []*Pvz `xml:"Pvz"`
}

type Pvz struct {
	Code           *string        `xml:"Code,attr"`
	PostalCode     *string        `xml:"PostalCode,attr"`
	Name           *string        `xml:"Name,attr"`
	CountryCode    *string        `xml:"CountryCode,attr"`
	CountryCodeIso *string        `xml:"countryCodeIso,attr"`
	CountryName    *string        `xml:"CountryName,attr"`
	RegionCode     *string        `xml:"RegionCode,attr"`
	RegionName     *string        `xml:"RegionName,attr"`
	CityCode       *string        `xml:"CityCode,attr"`
	City           *string        `xml:"City,attr"`
	WorkTime       *string        `xml:"WorkTime,attr"`
	Address        *string        `xml:"Address,attr"`
	FullAddress    *string        `xml:"FullAddress,attr"`
	AddressComment *string        `xml:"AddressComment,attr"`
	Phone          *string        `xml:"Phone,attr"`
	Email          *string        `xml:"Email,attr"`
	QqId           *string        `xml:"qqId,attr"`
	Note           *string        `xml:"Note,attr"`
	CoordX         *string        `xml:"coordX,attr"`
	CoordY         *string        `xml:"coordY,attr"`
	Type           *string        `xml:"Type,attr"`
	OwnerCode      *string        `xml:"ownerCode,attr"`
	IsDressingRoom *string        `xml:"IsDressingRoom,attr"`
	HaveCashless   *string        `xml:"HaveCashless,attr"`
	AllowedCod     *string        `xml:"AllowedCod,attr"`
	NearestStation *string        `xml:"NearestStation,attr"`
	MetroStation   *string        `xml:"MetroStation,attr"`
	Site           *string        `xml:"Site,attr"`
	OfficeImage    []*OfficeImage `xml:"OfficeImage"`
	WorkTimeY      []*WorkTimeY   `xml:"WorkTimeY"`
	WeightLimit    *WeightLimit   `xml:"WeightLimit"`
}

type OfficeImage struct {
	URL *string `xml:"url,attr"`
}

type WorkTimeY struct {
	Day     *string `xml:"day,attr"`
	Periods *string `xml:"periods,attr"`
}

type WeightLimit struct {
	WeightMin *string `xml:"WeightMin,attr"`
	WeightMax *string `xml:"WeightMax,attr"`
}
