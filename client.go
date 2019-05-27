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
	RegisterOrder(req RegisterOrderReq) (*RegisterOrderRes, error)
	UpdateOrder(req UpdateOrderReq) (*UpdateOrderRes, error)
	DeleteOrder(req DeleteOrderReq) (*DeleteOrderRes, error)
	GetPvzList(filter map[PvzListFilter]string) (*PvzList, error)
	GetRegions(filter map[RegionFilter]string) (*GetRegionsRes, error)
	GetCities(filter map[CityFilter]string) (*GetCitiesRes, error)
	CalculateDelivery(getCostReq GetCostReq) (*GetCostRes, error)
	GetStatusReport(statusReportReq StatusReportReq) (*StatusReportRes, error)
}

func (cl client) GetPvzList(filter map[PvzListFilter]string) (*PvzList, error) {
	return getPvzList(cl.clientConfig, filter)
}

func (cl client) CalculateDelivery(getCostReq GetCostReq) (*GetCostRes, error) {
	return calculateDelivery(cl.clientConfig, getCostReq)
}

func (cl client) GetCities(filter map[CityFilter]string) (*GetCitiesRes, error) {
	return getCities(cl.clientConfig, filter)
}

func (cl client) GetRegions(filter map[RegionFilter]string) (*GetRegionsRes, error) {
	return getRegions(cl.clientConfig, filter)
}

func (cl client) GetStatusReport(statusReportReq StatusReportReq) (*StatusReportRes, error) {
	return getStatusReport(cl.clientConfig, statusReportReq)
}

func (cl client) RegisterOrder(req RegisterOrderReq) (*RegisterOrderRes, error) {
	return registerOrder(cl.clientConfig, req)
}

func (cl client) UpdateOrder(req UpdateOrderReq) (*UpdateOrderRes, error) {
	return updateOrder(cl.clientConfig, req)
}

func (cl client) DeleteOrder(req DeleteOrderReq) (*DeleteOrderRes, error) {
	return deleteOrder(cl.clientConfig, req)
}
