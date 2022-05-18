package cdek

type securable struct {
	Account *string `xml:"Account,attr" json:"authLogin,omitempty"`
	Date    *string `xml:"Date,attr" json:"dateExecute,omitempty"`
	Secure  *string `xml:"Secure,attr" json:"secure,omitempty"`
}

//TODO: there are some methods that MUST HAVE auth, need to handle this case
func (s *securable) setAuth(auth *auth) *securable {
	if auth == nil {
		return s
	}

	s.Account = &auth.account

	date, sec := auth.encodedSecure()
	s.Date = &date
	s.Secure = &sec

	return s
}
