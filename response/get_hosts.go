package response

import "encoding/xml"

type GetHosts struct {
	Response

	Host            []DomainHost             `xml:"host" json:"host"`
	DomainServices  DomainHostDomainServices `xml:"DomainServices" json:"domain_services"`
	DomainLimits    DomainHostDomainLimits   `xml:"DomainLimits" json:"domain_limits"`
	HostRecordCount int                      `xml:"HostRecordCount" json:"host_record_count"`
	MxRecordCount   int                      `xml:"MxRecordCount" json:"mx_record_count"`
}

type DomainHost struct {
	XMLName xml.Name `xml:"host"`
	Name    string   `xml:"name" json:"name"`
	Type    string   `xml:"type" json:"type"`
	Address string   `xml:"address" json:"address"`
	HostID  string   `xml:"hostid" json:"host_id"`
}

type DomainHostDomainServices struct {
	XMLName         xml.Name `xml:"DomainServices"`
	EmailForwarding bool     `xml:"EmailForwarding" json:"email_forwarding"`
	HostRecords     bool     `xml:"HostRecords" json:"host_records"`
}

type DomainHostDomainLimits struct {
	HostsRecordLimit int `xml:"HostsRecordLimit" json:"hosts_record_limit"`
	MailForwardLimit int `xml:"MailForwardLimit" json:"mail_forward_limit"`
}
