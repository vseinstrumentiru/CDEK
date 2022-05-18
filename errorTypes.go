package cdek

import (
	"fmt"
	"strings"
)

//Error error values in responses
type Error struct {
	ErrorCode *string `xml:"ErrorCode,attr,omitempty" json:"code"`
	Msg       *string `xml:"Msg,attr,omitempty" json:"text"`
}

//IsErroneous check if struct has error
func (e *Error) IsErroneous() bool {
	return e.ErrorCode != nil
}

func (e Error) Error() string {
	errorMsgParts := []string{
		*e.Msg,
		fmt.Sprintf("ErrorCode: %s", *e.ErrorCode),
	}

	return strings.Join(errorMsgParts, "; ")
}

//AlertResponse CDEK Alerts model
type AlertResponse struct {
	Alerts []*Alert
}

//Alert CDEK Alert model
type Alert struct {
	Type      string
	Msg       string
	ErrorCode string
}

func (a *Alert) Error() string {
	return fmt.Sprintf("Type: %s; Msg: %s; ErrorCode: %s", a.Type, a.Msg, a.ErrorCode)
}
