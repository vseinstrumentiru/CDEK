package cdek

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
