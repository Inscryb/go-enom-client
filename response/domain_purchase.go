package response

type DomainPurchase struct {
	Response
	OrderID           string             `xml:"OrderID" json:"order_id"`
	OrderDelayed      bool               `xml:"OrderDelayed" json:"order_delayed"`
	OrderStatus       string             `xml:"OrderStatus" json:"order_status"`
	DomainInfo        DomainPurchaseInfo `xml:"Info" json:"domain_info"`
	TotalCharged      float32            `xml:"TotalCharged" json:"total_charged"`
	RegistrantPartyID string             `xml:"RegistrantPartyID"  json:"registrant_party_id"`
}

type DomainPurchaseInfo struct {
	RegistryCreateDate string `xml:"RegistryCreateDate" json:"registry_create_date"`
	RegistryExpDate    string `xml:"RegistryExpDate" json:"registry_exp_date"`
}
