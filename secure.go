package cdek

type credentialsXML struct {
	Account *string `xml:"Account,attr"`
	Date    *string `xml:"Date,attr"`
	Secure  *string `xml:"Secure,attr"`
}

type credentialsJSON struct {
	AuthLogin   *string `json:"authLogin,omitempty"`
	Secure      *string `json:"secure,omitempty"`
	DateExecute *string `json:"dateExecute,omitempty"`
}

//TODO: there are some methods that MUST HAVE auth, need to handle this case
func (s *credentialsXML) setAuth(auth *auth) *credentialsXML {
	if auth == nil {
		return s
	}

	s.Account = &auth.account

	date, sec := auth.encodedSecure()
	s.Date = &date
	s.Secure = &sec

	return s
}

func (s *credentialsJSON) setAuth(auth *auth) *credentialsJSON {
	if auth == nil {
		return s
	}

	s.AuthLogin = &auth.account

	date, sec := auth.encodedSecure()
	s.DateExecute = &date
	s.Secure = &sec

	return s
}
