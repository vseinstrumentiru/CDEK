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

func (cl client) GetPvzList(filter map[PvzListFilter]string) ([]*Pvz, error) {
	return getPvzList(cl.clientConf, filter)
}

func (cl client) CalculateDelivery(getCostReq GetCostReq) (*GetCostRespResult, error) {
	return calculateDelivery(cl.clientConf, getCostReq)
}

func (cl client) GetCities(filter map[CityFilter]string) (*GetCitiesResp, error) {
	return getCities(cl.clientConf, filter)
}

func (cl client) GetRegions(filter map[RegionFilter]string) (*GetRegionsResp, error) {
	return getRegions(cl.clientConf, filter)
}

func (cl client) GetStatusReport(statusReportReq StatusReportReq) (*StatusReportResp, error) {
	return getStatusReport(cl.clientConf, statusReportReq)
}

func (cl client) RegisterOrder(req RegisterOrderReq) (*RegisterOrderResp, error) {
	return registerOrder(cl.clientConf, req)
}

func (cl client) UpdateOrder(req UpdateOrderReq) (*UpdateOrderResp, error) {
	return updateOrder(cl.clientConf, req)
}

func (cl client) DeleteOrder(req DeleteOrderReq) (*DeleteOrderResp, error) {
	return deleteOrder(cl.clientConf, req)
}
