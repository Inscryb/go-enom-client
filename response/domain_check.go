package response

type DomainCheck struct {
	Response
	Name       string            `xml:"Name" json:"name"`
	IsPremium  bool              `xml:"IsPremium" json:"is_premium"`
	IsPlatinum bool              `xml:"IsPlatinum" json:"is_platinum"`
	IsEAP      bool              `xml:"IsEAP" json:"is_eap"`
	Prices     DomainCheckPrices `xml:"Prices" json:"prices"`
}

type DomainCheckPrices struct {
	Registration string `xml:"Registration" json:"Registration"`
	Renewal      string `xml:"Renewal" json:"Renewal"`
	Restore      string `xml:"Restore" json:"Restore"`
	Transfer     string `xml:"Transfer" json:"Transfer"`
}

func (r DomainCheck) IsAvailable() bool {
	return r.ResponseCode == 210
}
