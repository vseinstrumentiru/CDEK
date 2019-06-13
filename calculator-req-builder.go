package cdek

const apiVersion = "1.0"

//NewGetCostReq create new instance of GetCostReq
func NewGetCostReq(senderCityId int, receiverCityId int, tariffId int) *GetCostReq {
	apiVersion := apiVersion

	return &GetCostReq{Version: &apiVersion, SenderCityId: &senderCityId, ReceiverCityId: &receiverCityId, TariffId: &tariffId}
}

func (getCostReq *GetCostReq) setAuth(auth Auth) *GetCostReq {
	getCostReq.AuthLogin = &auth.Account

	date, sec := auth.EncodedSecure()
	getCostReq.DateExecute = &date
	getCostReq.Secure = &sec

	return getCostReq
}

//AddService add service to request
func (getCostReq *GetCostReq) AddService(service ServiceReq) *GetCostReq {
	if getCostReq.Services == nil {
		getCostReq.Services = []*ServiceReq{}
	}

	getCostReq.Services = append(getCostReq.Services, &service)

	return getCostReq
}

//AddGood add good to request
func (getCostReq *GetCostReq) AddGood(good Good) *GetCostReq {
	if getCostReq.Goods == nil {
		getCostReq.Goods = []*Good{}
	}

	getCostReq.Goods = append(getCostReq.Goods, &good)

	return getCostReq
}
