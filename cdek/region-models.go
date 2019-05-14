package cdek

type GetRegionsResponse []struct {
	RegionUUID     string `json:"regionUuid"`
	RegionName     string `json:"regionName"`
	Prefix         string `json:"prefix,omitempty"`
	RegionCodeExt  string `json:"regionCodeExt"`
	RegionCode     string `json:"regionCode"`
	RegionFiasGUID string `json:"regionFiasGuid,omitempty"`
	CountryName    string `json:"countryName"`
	CountryCode    string `json:"countryCode,omitempty"`
	CountryCodeExt int    `json:"countryCodeExt,omitempty,string"`
}
