package response

import "encoding/xml"

type TldList struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	TldList []string `xml:"tldlist>tld>tld" json:"tld_list"`
}
