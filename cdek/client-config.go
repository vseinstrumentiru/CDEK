package cdek

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"time"
)

func GetServerUrl() string {
	return os.Getenv("SERVER_URL")
}

type ClientConfig struct {
	Account   string
	Secure    string
	XmlApiUrl string
}

func (cl ClientConfig) EncodedSecure() (date string, encodedSecure string) {
	date = time.Now().Format("2006-01-02")
	encoder := md5.New()
	encoder.Write([]byte(date + "&" + cl.Secure))
	return date, hex.EncodeToString(encoder.Sum(nil))
}
