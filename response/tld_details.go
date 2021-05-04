package response

import (
	"encoding/xml"
)

type TldDetails struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	Tlds TldDetail `xml:"tlds>tld" json:"tlds"`
}

type TldDetail struct {
	Tld                     string `xml:"TLD" json:"tld"`
	AbleToLock              bool   `xml:"AbleToLock" json:"able_to_lock"`
	AllowWPPS               bool   `xml:"AllowWPPS" json:"allow_wpps"`
	AutoRenewOnly           bool   `xml:"AutoRenewOnly" json:"auto_renew_only"`
	ExtAttributes           bool   `xml:"ExtAttributes" json:"ext_attributes"`
	AllowWBL                bool   `xml:"AllowWBL" json:"allow_wbl"`
	TransRequiresNewContact bool   `xml:"TransRequiresNewContact" json:"trans_requires_new_contact"`
	EncodingType            string `xml:"EncodingType" json:"encoding_type"`

	ClearPhoneNumber    bool                  `xml:"ClearPhoneNumber" json:"clear_phone_number"`
	ClearFax            bool                  `xml:"ClearFax" json:"clear_fax"`
	ValidateDNSHosting  bool                  `xml:"ValidateDNSHosting" json:"validate_dns_hosting"`
	HasPremiumNames     bool                  `xml:"HasPremiumNames" json:"has_premium_names"`
	SupportsDnsSec      bool                  `xml:"SupportsDnsSec" json:"supports_dns_sec"`
	AllowWhoisPublicity bool                  `xml:"AllowWhoisPublicity" json:"allow_whois_publicity"`
	Registration        TldDetailRegistration `xml:"Registration" json:"registration"`
	Renewal             TldDetailRenewal      `xml:"Renewal" json:"renewal"`
	Transfer            TldDetailTransfer     `xml:"Transfer" json:"transfer"`
	PremiumNames        TldDetailPremiumNames `xml:"PremiumNames" json:"premium_names"`
	NameServers         TldDetailNameServers  `xml:"NameServers" json:"name_servers"`
}

type TldDetailRegistration struct {
	Realtime                     bool   `xml:"Realtime" json:"realtime"`
	Unit                         string `xml:"Unit" json:"unit"`
	Minimum                      int    `xml:"Minimum" json:"minimum"`
	Maximum                      int    `xml:"Maximum" json:"maximum"`
	HasPremiumNames              bool   `xml:"HasPremiumNames" json:"has_premium_names"`
	ExtAttributes                bool   `xml:"ExtAttributes" json:"ext_attributes"`
	ExtAttributesDomainLevel     bool   `xml:"ExtAttributesDomainLevel" json:"ext_attributes_domain_level"`
	DNSRequired                  bool   `xml:"DNSRequired" json:"dns_required"`
	DNSMinimum                   int    `xml:"DNSMinimum" json:"dns_minimum"`
	DNSMaximum                   int    `xml:"DNSMaximum" json:"dns_maximum"`
	GeneralAvailabilityStartDate string `xml:"GeneralAvailabilityStartDate" json:"general_availability_start_date"`
}

type TldDetailRenewal struct {
	AutoRenewOnly        bool `xml:"AutoRenewOnly" json:"auto_renew_only"`
	AutoRenewed          bool `xml:"AutoRenewed" json:"auto_renewed"`
	RenewBeforeExpMonths int  `xml:"RenewBeforeExpMonths" json:"renew_before_exp_months"`
	DeleteType           int  `xml:"DeleteType" json:"delete_type"`
	DeleteDay            int  `xml:"DeleteDay" json:"delete_day"`
	GracePeriod          int  `xml:"GracePeriod" json:"grace_period"`
	Reactivate           bool `xml:"Reactivate" json:"reactivate"`
	Restorable           bool `xml:"Restorable" json:"restorable"`
	RGP                  bool `xml:"RGP" json:"rgp"`
	RGPDays              int  `xml:"RGPDays" json:"rgp_days"`
	ExtendedRGP          bool `xml:"ExtendedRGP" json:"extended_rgp"`
	TransferPeriod       int  `xml:"TransferPeriod" json:"transfer_period"`
}

type TldDetailTransfer struct {
	Transferable       bool `xml:"Transferable" json:"transferable"`
	AuthInfo           bool `xml:"AuthInfo" json:"auth_info"`
	Realtime           bool `xml:"Realtime" json:"realtime"`
	AutoVerification   bool `xml:"AutoVerification" json:"auto_verification"`
	AutoFax            bool `xml:"AutoFax" json:"auto_fax"`
	TransferByFOA      bool `xml:"TransferByFOA" json:"transfer_by_foa"`
	RequiresNewContact bool `xml:"RequiresNewContact" json:"requires_new_contact"`
}

type TldDetailPremiumNames struct {
	HasPremiumNames    bool   `xml:"HasPremiumNames" json:"has_premium_names"`
	MaxPremiumRegYears int    `xml:"MaxPremiumRegYears" json:"max_premium_reg_years"`
	TrademarkStartDate string `xml:"TrademarkStartDate" json:"trademark_start_date"`
	TrademarkEndDate   string `xml:"TrademarkEndDate" json:"trademark_end_date"`
}

type TldDetailNameServers struct {
	NameServers    bool `xml:"NameServers" json:"name_servers"`
	MinNameServers int  `xml:"MinNameServers" json:"min_name_servers"`
	MaxNameServers int  `xml:"MaxNameServers" json:"max_name_servers"`
}
