package main

import (
	"cdek_sdk/cdek"
	"fmt"
)

func main() {
	config := cdek.ClientConfig{
		Account:   "f62dcb094cc91617def72d9c260b4483",
		Secure:    "6bd3937dcebd15beb25278bc0657014c",
		XmlApiUrl: "https://integration.edu.cdek.ru",
	}

	client := cdek.NewClient(config)
	filterBuilder := cdek.PvzListFilterBuilder{}
	filterBuilder = filterBuilder.AddFilter(cdek.FilterCityId, "")
	pvzlist, _ := client.GetPvzList(filterBuilder.Filter())

	for i := 0; len(pvzlist.Pvz) > i; i++ {
		fmt.Println(pvzlist.Pvz[i].CityCode, " ", pvzlist.Pvz[i].City)
	}
}
