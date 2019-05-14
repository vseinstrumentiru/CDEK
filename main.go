package main

import (
	"cdek_sdk/cdek"
	"fmt"
)

func main() {
	tryPvzlist()
	tryCities()
	tryRegions()
}

func getClientConfig() *cdek.ClientConfig {
	return &cdek.ClientConfig{
		Account:   "f62dcb094cc91617def72d9c260b4483",
		Secure:    "6bd3937dcebd15beb25278bc0657014c",
		XmlApiUrl: "https://integration.edu.cdek.ru",
	}
}

func tryRegions()  {
	filterBuilder := cdek.RegionFilterBuilder{}
	filterBuilder.AddFilter(cdek.RegionFilterSize, "1")

	regions, err := cdek.NewClient(*getClientConfig()).GetRegions(filterBuilder.Filter())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(regions)
}

func tryCities()  {
	filterBuilder := cdek.CityFilterBuilder{}
	filterBuilder.AddFilter(cdek.CityFilterSize, "1")

	cities, err := cdek.NewClient(*getClientConfig()).GetCities(filterBuilder.Filter())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cities)
}

func tryPvzlist()  {
	filterBuilder := cdek.PvzListFilterBuilder{}
	filterBuilder.AddFilter(cdek.PvzListFilterCityId, "")
	pvzlist, _ := cdek.NewClient(*getClientConfig()).GetPvzList(filterBuilder.Filter())

	for i := 0; len(pvzlist.Pvz) > i; i++ {
		fmt.Println(pvzlist.Pvz[i].CityCode, " ", pvzlist.Pvz[i].City)
	}
}