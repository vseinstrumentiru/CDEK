package main

import (
	"cdek_sdk/cdek"
	"fmt"
)

func main() {
	config := cdek.ClientConfig{
		Account: "f62dcb094cc91617def72d9c260b4483",
		Secure: "6bd3937dcebd15beb25278bc0657014c",
		XmlApiUrl: "https://integration.edu.cdek.ru",
	}

	fmt.Println(config.EncodedSecure())

	client := cdek.NewClient(config)
	pvzlist, _ := client.GetPvzList(make(map[cdek.Filter]string))
	fmt.Println(pvzlist)
}
