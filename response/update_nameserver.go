package response

type UpdateNameserver struct {
	Response

	NsSuccess int `xml:"RegisterNameserver>NsSuccess" json:"ns_success"`
}
