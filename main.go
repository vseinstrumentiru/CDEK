package main

import (
	"cdek_sdk/cdek"
	"fmt"
	"time"
)

var cdekClient cdek.Client

func main() {
	//tryPvzlist()
	//tryCities()
	//tryRegions()
	//tryCalculator()
	//tryStatusReport()
	tryRegisterOrder()
}

func getClientConfig() *cdek.ClientConf {
	return &cdek.ClientConf{
		Auth: cdek.Auth{
			Account: "f62dcb094cc91617def72d9c260b4483",
			Secure:  "6bd3937dcebd15beb25278bc0657014c",
		},
		XmlApiUrl: "https://integration.edu.cdek.ru",
	}
}

func client() cdek.Client {
	if cdekClient == nil {
		cdekClient = cdek.NewClient(*getClientConfig())
	}

	return cdekClient
}

func tryRegions() {
	filterBuilder := cdek.RegionFilterBuilder{}
	filterBuilder.AddFilter(cdek.RegionFilterSize, "")

	regions, err := client().GetRegions(filterBuilder.Filter())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(regions)
}

func tryCities() {
	filterBuilder := cdek.CityFilterBuilder{}
	filterBuilder.AddFilter(cdek.CityFilterCityName, "Нижний Новгород")

	cities, err := client().GetCities(filterBuilder.Filter())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, city := range *cities {
		fmt.Println("город:", city.CityName, "Код:", city.CityCode)
	}
}

func tryPvzlist() {
	filterBuilder := cdek.PvzListFilterBuilder{}
	filterBuilder.AddFilter(cdek.PvzListFilterCityId, "44")
	pvzlist, _ := client().GetPvzList(filterBuilder.Filter())

	for _, pvz := range pvzlist.Pvz {
		fmt.Println(pvz.City, pvz.Name, pvz.Site, pvz.Code)
	}
}

func tryCalculator() {
	fmt.Println(client().CalculateDelivery(cdek.GetCostReq{
		Version:        "1.0",
		SenderCityId:   44,
		ReceiverCityId: 414,
		TariffId:       1,
		Goods: []cdek.Good{
			{
				Weight: 1,
				Length: 1,
				Width:  1,
				Height: 1,
			},
		},
	}))
}

func tryStatusReport() {
	StatusReportReq := cdek.NewStatusReportReq()
	changePeriod := cdek.NewChangePeriod().
		SetDateFirst(time.Now().AddDate(0, 0, -7)).
		SetDateLast(time.Now())

	StatusReportReq.SetShowHistory(true).
		SetShowReturnOrder(false).
		SetChangePeriod(*changePeriod)

	order := cdek.NewStatusReportOrderReq().
		SetDispatchNumber(123)

	StatusReportReq.AddOrder(*order)

	fmt.Println(client().GetStatusReport(*StatusReportReq))
}

func tryRegisterOrder() {
	addr := cdek.NewAddress().
		SetPvzCode("MSK117")
	pack := cdek.NewOrderPackage().
		SetWeight(321).
		SetNumber("123-321-123").
		SetBarCode("123-321-123")
	order := cdek.NewOrderReq().
		SetNumber("123-321-123").
		AddPackage(*pack).
		SetSendCityCode(441).
		SetRecCityCode(44).
		SetTariffTypeCode(123).
		SetPhone("12457655435346").
		SetRecipientName("Yasya").
		SetAddress(*addr)
	req := cdek.NewDeliveryRequest().
		SetNumber("123-321-123").
		SetOrderCount("1").
		SetOrder(*order)

	res, e := client().RegisterOrder(*req)
	fmt.Printf("%+v%+v\n", res, e)
	fmt.Println("_________________________________________________________________")
	for _, o := range res.Order {
		fmt.Println(*o.Msg)
	}
}
