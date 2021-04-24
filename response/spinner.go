package response

type SpinnerDomain struct {
	//XMLName  xml.Name `xml:"domain"`
	Name     string `xml:"name,attr" json:"name"`
	Com      string `xml:"com,attr" json:"com"`
	ComScore string `xml:"comscore,attr" json:"comscore"`
	Net      string `xml:"net,attr" json:"net"`
	NetScore string `xml:"netscore,attr" json:"netscore"`
	Tv       string `xml:"tv,attr" json:"tv"`
	TvScore  string `xml:"tvscore,attr" json:"tvscore"`
	Cc       string `xml:"cc,attr" json:"cc"`
	CcScore  string `xml:"ccscore,attr" json:"ccscore"`
}

type NameSpinner struct {
	SpinCount int    `xml:"spincount" json:"spin_count"`
	TLDList   string `xml:"TLDList" json:"tld_list"`
	//XMLName   xml.Name        `xml:"domain"`
	Domains []SpinnerDomain `xml:"domains>domain" json:"domains"`
}

type SpinnerResponse struct {
	Response

	NameSpinner NameSpinner `xml:"namespin" json:"namespin"`
}
