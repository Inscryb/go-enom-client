package response

type SpinnerDomain struct {
	Name     string `xml:"name" json:"name"`
	Com      string `xml:"com" json:"com"`
	ComScore string `xml:"comscore" json:"comscore"`
	Net      string `xml:"net" json:"net"`
	NetScore string `xml:"netscore" json:"netscore"`
	Tv       string `xml:"tv" json:"tv"`
	TvScore  string `xml:"tvscore" json:"tvscore"`
	Cc       string `xml:"cc" json:"cc"`
	CcScore  string `xml:"ccscore" json:"ccscore"`
}

type NameSpinner struct {
	SpinCount int             `xml:"spincount" json:"spin_count"`
	Domains   []SpinnerDomain `xml:"domains" json:"domains"`
}

type SpinnerResponse struct {
	Response

	NameSpinner NameSpinner `xml:"namespin" json:"namespin"`
}
