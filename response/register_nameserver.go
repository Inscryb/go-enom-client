package response

type RegisterNameserver struct {
	Response

	NsSuccess int `xml:"RegisterNameserver>NsSuccess" json:"ns_success"`
}
