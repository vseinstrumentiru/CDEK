package cdek

type securableXML struct {
	Account *string `xml:"Account,attr"`
	Date    *string `xml:"Date,attr"`
	Secure  *string `xml:"Secure,attr"`
}

type securableJSON struct {
	AuthLogin   *string `json:"authLogin,omitempty"`
	Secure      *string `json:"secure,omitempty"`
	DateExecute *string `json:"dateExecute,omitempty"`
}

//TODO: there are some methods that MUST HAVE auth, need to handle this case
func (s *securableXML) setAuth(auth *Auth) *securableXML {
	if auth == nil {
		return s
	}

	s.Account = &auth.Account

	date, sec := auth.EncodedSecure()
	s.Date = &date
	s.Secure = &sec

	return s
}

func (s *securableJSON) setAuth(auth *Auth) *securableJSON {
	if auth == nil {
		return s
	}

	s.AuthLogin = &auth.Account

	date, sec := auth.EncodedSecure()
	s.DateExecute = &date
	s.Secure = &sec

	return s
}
