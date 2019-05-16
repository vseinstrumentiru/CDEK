package cdek

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type ClientConfig struct {
	Auth      Auth
	XmlApiUrl string
}

type Auth struct {
	Account string
	Secure  string
}

func (a Auth) EncodedSecure() (date string, encodedSecure string) {
	date = time.Now().Format("2006-01-02")
	encoder := md5.New()
	encoder.Write([]byte(date + "&" + a.Secure))

	return date, hex.EncodeToString(encoder.Sum(nil))
}
