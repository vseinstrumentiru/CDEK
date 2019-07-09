package cdek

const jsonContentType = "application/json"
const xmlContentType = "application/xml"
const urlFormEncoded = "application/x-www-form-urlencoded"

//NewClient creates new instance of Client
func NewClient(clientConf ClientConf) Client {
	return &client{
		clientConf: clientConf,
	}
}

type client struct {
	clientConf ClientConf
}

//Client CDEK SDK client with public methods
type Client interface {
	RegisterOrder(req RegisterOrderReq) (*RegisterOrderResp, error)
	UpdateOrder(req UpdateOrderReq) (*UpdateOrderResp, error)
	DeleteOrder(req DeleteOrderReq) (*DeleteOrderResp, error)
	GetPvzList(filter map[PvzListFilter]string) ([]*Pvz, error)
	GetRegions(filter map[RegionFilter]string) (*GetRegionsResp, error)
	GetCities(filter map[CityFilter]string) (*GetCitiesResp, error)
	CalculateDelivery(getCostReq GetCostReq) (*GetCostRespResult, error)
	GetStatusReport(statusReportReq StatusReportReq) (*StatusReportResp, error)
}
