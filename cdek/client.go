package cdek

const jsonContentType = "application/json"

func NewClient(clientConfig ClientConfig) *client {
	return &client{
		clientConfig: clientConfig,
	}
}

type client struct {
	clientConfig ClientConfig
}

type Client interface {
	RegisterOrder()
	UpdateOrder()
	DeleteOrder()
	GetPvzList()
	GetRegions()
	GetCities()
	CalculateDelivery()
}

func (cl client) GetPvzList(filter map[PvzListFilter]string) (*PvzList, error) {
	return getPvzList(cl.clientConfig, filter)
}

func (cl client) CalculateDelivery(getCostRequest GetCostRequest) (*GetCostResponse, error) {
	return calculateDelivery(getCostRequest)
}

func (cl client) GetCities(filter map[CityFilter]string) (*GetCitiesResponse, error) {
	return getCities(cl.clientConfig, filter)
}

func (cl client) GetRegions(filter map[RegionFilter]string) (*GetRegionsResponse, error) {
	return getRegions(cl.clientConfig, filter)
}

// TODO

func (cl client) RegisterOrder() (*PvzList, error) {
	return nil, nil
}

func (cl client) UpdateOrder() (*PvzList, error) {
	return nil, nil
}

func (cl client) DeleteOrder() (*PvzList, error) {
	return nil, nil
}



