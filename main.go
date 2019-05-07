package main

import (
	"github.com/joho/godotenv"
)

const version = "0.0.1"

func main() {
	_ = godotenv.Load()

	//
	//args := os.Args
	//
	//if len(args) < 2 {
	//	PrintUsage()
	//	return
	//}
	//
	//switch cmd := args[1]; cmd {
	//case "version":
	//	fmt.Println("version: ", version)
	//case "help":
	//	PrintUsage()
	//case "pvzlist":
	//	filter := map[string]string{
	//		pvzlist.FilterType:   pvzlist.TypePvz,
	//		pvzlist.FilterCityId: "417",
	//	}
	//	fmt.Println(pvzlist.GetPvzList(filter))
	//case "calc":
	//	fmt.Println(calculator.Calculate(calculator.GetCostRequest{}))
	//default:
	//	PrintUsage()
	//}
}
