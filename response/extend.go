package response

type EDomainInfo struct {
	RegistryExpDate string `xml:"RegistryExpDate" json:"registry_expiration_date"`
}

type Extend struct {
	Response

	Extension  string `xml:"Extension"  json:"extension"`
	DomainName string `xml:"DomainName" json:"domain_name"`
	OrderID    int    `xml:"OrderID"    json:"order_id"`

	DomainInfo EDomainInfo `xml:"DomainInfo" json:"domain_info"`
}
