package cdek

//GetRegionsResp List of Regions response
type GetRegionsResp []Region

//Region Region response
type Region struct {
	RegionUUID     *string `json:"regionUuid"`
	RegionName     *string `json:"regionName"`
	Prefix         *string `json:"prefix,omitempty"`
	RegionCodeExt  *int    `json:"regionCodeExt,string,omitempty"`
	RegionCode     *int    `json:"regionCode,string,omitempty"`
	RegionFiasGUID *string `json:"regionFiasGuid,omitempty"`
	CountryName    *string `json:"countryName"`
	CountryCode    *string `json:"countryCode,omitempty"`
	CountryCodeExt *int    `json:"countryCodeExt,omitempty,string"`
}
