package cdek

const jsonContentType = "application/json"
const xmlContentType = "application/xml"
const urlFormEncoded = "application/x-www-form-urlencoded"

func NewClient(clientConfig ClientConf) Client {
	return &client{
		clientConfig: clientConfig,
	}
}

type client struct {
	clientConfig ClientConf
}

type Client interface {
	RegisterOrder(req RegisterOrderReq) (*RegisterOrderResp, error)
	UpdateOrder(req UpdateOrderReq) (*UpdateOrderResp, error)
	DeleteOrder(req DeleteOrderReq) (*DeleteOrderResp, error)
	GetPvzList(filter map[PvzListFilter]string) (*PvzList, error)
	GetRegions(filter map[RegionFilter]string) (*GetRegionsResp, error)
	GetCities(filter map[CityFilter]string) (*GetCitiesResp, error)
	CalculateDelivery(getCostReq GetCostReq) (*GetCostResp, error)
	GetStatusReport(statusReportReq StatusReportReq) (*StatusReportResp, error)
}

func (cl client) GetPvzList(filter map[PvzListFilter]string) (*PvzList, error) {
	return getPvzList(cl.clientConfig, filter)
}

func (cl client) CalculateDelivery(getCostReq GetCostReq) (*GetCostResp, error) {
	return calculateDelivery(cl.clientConfig, getCostReq)
}

func (cl client) GetCities(filter map[CityFilter]string) (*GetCitiesResp, error) {
	return getCities(cl.clientConfig, filter)
}

func (cl client) GetRegions(filter map[RegionFilter]string) (*GetRegionsResp, error) {
	return getRegions(cl.clientConfig, filter)
}

func (cl client) GetStatusReport(statusReportReq StatusReportReq) (*StatusReportResp, error) {
	return getStatusReport(cl.clientConfig, statusReportReq)
}

func (cl client) RegisterOrder(req RegisterOrderReq) (*RegisterOrderResp, error) {
	return registerOrder(cl.clientConfig, req)
}

func (cl client) UpdateOrder(req UpdateOrderReq) (*UpdateOrderResp, error) {
	return updateOrder(cl.clientConfig, req)
}

func (cl client) DeleteOrder(req DeleteOrderReq) (*DeleteOrderResp, error) {
	return deleteOrder(cl.clientConfig, req)
}
