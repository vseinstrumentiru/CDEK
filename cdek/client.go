package cdek

const jsonContentType = "application/json"
const xmlContentType = "application/xml"

func NewClient(clientConfig ClientConf) Client {
	return &client{
		clientConfig: clientConfig,
	}
}

type client struct {
	clientConfig ClientConf
}

type Client interface {
	RegisterOrder() (*interface{}, error)
	UpdateOrder() (*interface{}, error)
	DeleteOrder() (*interface{}, error)
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


// TODO

func (cl client) RegisterOrder() (*interface{}, error) {
	return nil, nil
}

func (cl client) UpdateOrder() (*interface{}, error) {
	return nil, nil
}

func (cl client) DeleteOrder() (*interface{}, error) {
	return nil, nil
}
