package response

type CheckNSStatus struct {
	Response

	Name         string   `xml:"CheckNsStatus>name"          json:"name"`
	AttribId     string   `xml:"CheckNsStatus>attrib-id"     json:"attrib_id"`
	IpAddress    string   `xml:"CheckNsStatus>ipaddress"     json:"ip_address"`
	AttribUpId   string   `xml:"CheckNsStatus>attrib-upid"   json:"attrib_up_id"`
	AttribClId   string   `xml:"CheckNsStatus>attrib-clid"   json:"attrib_cl_id"`
	AttribCrId   string   `xml:"CheckNsStatus>attrib-crid"   json:"attrib_cr_id"`
	Statuses     []string `xml:"CheckNsStatus>status>status" json:"statuses"`
	AttribUpdate string   `xml:"CheckNsStatus>attrib-update" json:"attrib_update"`
	AttribCrdate string   `xml:"CheckNsStatus>attrib-crdate" json:"attrib_crdate"`
}
